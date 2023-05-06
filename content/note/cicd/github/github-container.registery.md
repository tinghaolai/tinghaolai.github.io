---
title: "github-container.registery.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["cicd","cicd-github"]
---



---

## GitHub Container Registry

**Your github project name**

https://github.com/davetestingaccount/test-for-ci

* Build image
  * `docker build -t local_image_name/test-for-ci:v0.1 .`
* Put tag
  * `docker tag local_image_name/test-for-ci:v0.1 ghcr.io/davetestingaccount/test-for-ci:v0.2`
* Login
  * `docker login ghcr.io -u davetestingaccount`
  * Try `sudo` if can't login
* Push to github
  * `docker push ghcr.io/<GITHUB_USERNAME><REPO_NAME>:<TAG>`
  * `docker push ghcr.io/davetestingaccount/test-for-ci:v0.2`
* Check upload
  * enter `ghcr.io/davetestingaccount/test-for-ci` in browser
    



---

