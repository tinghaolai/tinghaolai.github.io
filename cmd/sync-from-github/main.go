package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type RepoItem struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Path        string `json:"path"`
	DownloadUrl string `json:"download_url"`
}

var client = &http.Client{}
var accessToken string
var repoPath string
var excludeFolders map[string]bool
var excludeFiles map[string]bool
var storeFolderRoot string

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	repoPath = os.Getenv("GITHUB_REPO_PATH")
	accessToken = os.Getenv("GITHUB_ACCESS_TOKEN")
	excludeFolders = make(map[string]bool)
	excludeFiles = make(map[string]bool)
	storeFolderRoot = "../../content"

	for _, key := range strings.Split(os.Getenv("EXCLUDE_FOLDERS"), ",") {
		excludeFolders[key] = true
	}

	for _, key := range strings.Split(os.Getenv("EXCLUDE_FILES"), ",") {
		excludeFiles[key] = true
	}

	clearFolder(storeFolderRoot)
	scanFolder("/")
}

func clearFolder(folderPath string) {
	os.RemoveAll(folderPath)
}

func scanFolder(subFolderPath string) {
	storeFolderPath := path.Join(storeFolderRoot, subFolderPath)
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repoPath, subFolderPath)
	req, _ := http.NewRequest("GET", apiUrl, nil)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var items []RepoItem
	err = json.Unmarshal(body, &items)
	if err != nil {
		bodyString := string(body)
		println(bodyString)
		println("panic: repoPath: " + repoPath + " storeFolderPath: " + storeFolderPath + ", apiUrl: " + apiUrl)
		panic(err)
	}

	os.MkdirAll(storeFolderPath, 0755)
	for _, item := range items {
		if item.Type == "file" {
			if _, ok := excludeFiles[item.Name]; ok {
				continue
			}

			resp, err := http.Get(item.DownloadUrl)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			content, _ := io.ReadAll(resp.Body)
			ext := filepath.Ext(item.Name)
			if ext != ".md" {
				file, err := os.Create(path.Join(storeFolderPath, item.Name))
				if err != nil {
					println("panic: repoPath: " + repoPath + " storeFolderPath: " + storeFolderPath + ", apiUrl: " + apiUrl)
					panic(err)
				}

				defer file.Close()
				_, err = file.Write(content)
				continue
			}

			writeMarkdown(subFolderPath, item.Name, string(content))
		} else if item.Type == "dir" {
			if _, ok := excludeFolders[item.Name]; ok {
				continue
			}

			scanFolder(path.Join(subFolderPath, item.Name))
		}
	}
}

func writeMarkdown(subFolderPath string, filename string, content string) {
	storeFolderPath := path.Join(storeFolderRoot, subFolderPath)
	re := regexp.MustCompile(`(?s)(<!--HugoNoteFlag-->)(.*)`)
	match := re.FindStringSubmatch(content)

	if len(match) > 0 {
		file, err := os.Create(path.Join(storeFolderPath, filename))
		if err != nil {
			panic(err)
		}

		zhFile, err := os.Create(path.Join(storeFolderPath, convertI18NFileName(filename)))
		if err != nil {
			panic(err)
		}

		trueContent := match[2]
		re = regexp.MustCompile(`(?s)(.*)(<!--HugoNoteZhFlag-->)(.*)`)
		match = re.FindStringSubmatch(trueContent)
		if len(match) > 0 {
			writeHugoFile(file, match[1], filename, subFolderPath)
			writeHugoFile(zhFile, match[3], filename, subFolderPath)
			return
		}

		writeHugoFile(file, trueContent, filename, subFolderPath)
		writeHugoFile(zhFile, "此文章沒有中文版本", filename, subFolderPath)

		return
	}

	draftFileName := "__draft_note__" + filename
	file, err := os.Create(path.Join(storeFolderPath, draftFileName))
	if err != nil {
		panic(err)
	}

	githubLink := "https://github.com/" + path.Join(repoPath, "blob/master/", subFolderPath, filename)

	writeHugoFile(file, "# "+filename+"\n\n"+"Draft note, watch origin note from github: [link]("+githubLink+")", draftFileName, subFolderPath)

	draftZhFileName := convertI18NFileName(draftFileName)
	file, err = os.Create(path.Join(storeFolderPath, draftZhFileName))
	if err != nil {
		panic(err)
	}

	writeHugoFile(file, "# "+filename+"\n\n"+"草稿筆記, 原版請前往 github 查看: [連結]("+githubLink+")", draftFileName, subFolderPath)
}

func convertI18NFileName(filename string) string {
	return filename[:len(filename)-3] + ".zh" + filename[len(filename)-3:]
}

func writeHugoFile(file *os.File, content string, fileName string, subFolderPath string) {
	if subFolderPath[0] == '/' {
		subFolderPath = subFolderPath[1:]
	}

	parts := strings.Split(subFolderPath, "/")
	var result []string
	for i := range parts {
		result = append(result, "\""+strings.Join(parts[:i+1], "-")+"\"")
	}

	categories := "[" + strings.Join(result, ",") + "]"
	content = "---\ntitle: \"" + fileName + "\"\ndate: 1919-08-10T11:45:14Z\ndraft: false\ncategories: " + categories + "\n---\n\n" + content
	file.WriteString(content)
}
