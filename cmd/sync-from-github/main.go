package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"path"
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

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	repoPath = os.Getenv("GITHUB_REPO_PATH")
	accessToken = os.Getenv("GITHUB_ACCESS_TOKEN")

	clearFolder("./sync-test/")
	scanFolder("", "./sync-test/")
}

func clearFolder(folderPath string) {
	os.RemoveAll(folderPath)
}

func scanFolder(subFolderPath string, storeFolderPath string) {
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
			resp, err := http.Get(item.DownloadUrl)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			content, _ := io.ReadAll(resp.Body)
			file, err := os.Create(storeFolderPath + item.Name)
			if err != nil {
				println("panic: repoPath: " + repoPath + " storeFolderPath: " + storeFolderPath + ", apiUrl: " + apiUrl)
				panic(err)
			}
			defer file.Close()

			_, err = file.Write(content)
			if err != nil {
				panic(err)
			}
		} else if item.Type == "dir" {
			scanFolder(path.Join(subFolderPath, item.Name), path.Join(storeFolderPath, item.Name))
		}
	}
}
