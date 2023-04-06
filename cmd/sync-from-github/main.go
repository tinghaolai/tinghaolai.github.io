package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

type RepoItem struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Path        string `json:"path"`
	DownloadUrl string `json:"download_url"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	repoName := os.Getenv("GITHUB_REPO_NAME")
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	client := &http.Client{}

	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/contents/", repoName)
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
		panic(err)
	}

	for _, item := range items {
		if item.Type == "file" {
			resp, err := http.Get(item.DownloadUrl)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			content, _ := io.ReadAll(resp.Body)
			file, err := os.Create("./sync-test/" + item.Path)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			_, err = file.Write(content)
			if err != nil {
				panic(err)
			}
		}
	}
}
