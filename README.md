### [Github page link](https://tinghaolai.github.io/)

## Hugo develop

### Check draft
`hugo server --port 1414 -D`

### Check release
`hugo server --port 1414`

## Sync from github into hugo

* `cd cmd/sync-from-github`
* `go run main.go`
* `bash sync-content.sh`
* `git push origin master`