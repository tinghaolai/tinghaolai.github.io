package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var fileContentList []map[string]string

func main() {
	fileContentList = make([]map[string]string, 0)
	fmt.Println("sync-content-to-navbar-content print")
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("fail get current folder:", err)
		return
	}

	currentDir = filepath.Dir(filepath.Dir(currentDir))
	currentDir = filepath.Join(currentDir, "content")
	currentDir = filepath.Join(currentDir, "note")
	runFolder(currentDir, "")

	writeJsonToNavbar()
}

func runFolder(currentDir string, contentPath string) {
	dir, err := os.Open(currentDir)
	if err != nil {
		fmt.Println("fail to open content folder", err)
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("can not read content folder", err)
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			runFolder(filepath.Join(currentDir, fileInfo.Name()), filepath.Join(contentPath, fileInfo.Name()))
		} else {
			handelMarkdownFile(filepath.Join(currentDir, fileInfo.Name()), filepath.Join(contentPath, fileInfo.Name()))
		}
	}
}

func handelMarkdownFile(path string, contentPath string) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("fail to read file:", err)
		return
	}

	contentString := string(content)
	re := regexp.MustCompile(`[^\p{L}\p{N}]+`)

	contentString = re.ReplaceAllString(contentString, " ")

	fileContentList = append(fileContentList, map[string]string{
		"filename": contentPath,
		"content":  contentString,
	})
}

func writeJsonToNavbar() {
	jsonData, err := json.Marshal(fileContentList)
	if err != nil {
		fmt.Println("fail to marshal json:", err)
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("fail get current folder:", err)
		return
	}

	currentDir = filepath.Dir(filepath.Dir(currentDir))
	currentDir = filepath.Join(currentDir, "layouts")
	currentDir = filepath.Join(currentDir, "partials")
	navPath := filepath.Join(currentDir, "navbar.html")
	content, err := os.ReadFile(navPath)
	if err != nil {
		fmt.Println("fail to read file:", err)
		return
	}

	contentString := string(content)
	pattern := `\/\*\s*dynamic\s*command\s*sync\s*\*\/([^\/]*)\/\*\s*dynamic\s*command\s*sync\s*\*\/`

	reg := regexp.MustCompile(pattern)

	match := reg.FindStringSubmatch(contentString)
	if match != nil {
		fullMatch := match[0]

		stringJsonData := string(jsonData)
		re := regexp.MustCompile(`\\`)
		stringJsonData = re.ReplaceAllString(stringJsonData, " ")

		replacement := "articles: JSON.parse('" + stringJsonData + "'),"

		replacedString := reg.ReplaceAllString(fullMatch, "/* dynamic command sync */"+replacement+"/* dynamic command sync */")

		contentString = reg.ReplaceAllString(contentString, replacedString)

		err := os.WriteFile(navPath, []byte(contentString), 0644)
		if err != nil {
			return
		}
	} else {
		fmt.Println("can not find match")
	}
}
