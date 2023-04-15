---
title: "docker-package-container.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["docker"]
---



# Translated by ChatGTP

## 容器套件

圖像名稱：我的專案

* 提交當前容器
  * docker commit myproject myproject-container:1.0
* 將容器匯出為檔案
  * docker save -o test.tar myproject:1.0
* 在另一個 Docker 環境中匯入
  * docker load < test.tar
* 執行容器
  * docker run myproject:1.0