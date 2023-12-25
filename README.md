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

> `go run main.go && bash sync-content.sh && git push origin master` 

> `cd cmd/sync-from-github && go run main.go && bash sync-content.sh && git push origin master`
 
## Update regex 

* `cd cmd/sync-content-to-navbar-content`
* `go run main.go`


## Skip Log Feature

* type `skipLog()` in blog website console to skip log
* type `startLog()`  in blog website console to start log