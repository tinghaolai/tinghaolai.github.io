---
title: "gitlab-ci-common-note.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["cicd","cicd-gitlab"]
---



# Translated by ChatGTP

## Gitlab CI 常用筆記

* 從 Gitlab 拉取映像檔
  * docker pull registry.gitlab.com/dave.lai/laravel-react-crawler/dave-test-image
    * `registry.gitlab.com` - Gitlab 網站
    * `dave.lai` - 擁有者或群組
    * `laravel-react-crawler` - 專案名稱
* 在映像檔作業中使用其他映像檔: 使用 [services](https://docs.gitlab.com/ee/ci/services/)。
