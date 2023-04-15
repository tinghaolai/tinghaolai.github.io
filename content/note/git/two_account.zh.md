---
title: "two_account.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["git"]
---



# Translated by ChatGTP

## 在一台電腦上使用兩個 git 帳號

當您需要在不同專案中使用不同的 git 帳號進行提交時，頻繁切換 git 帳號會變得非常痛苦。

### 我的解決方案

運行一個 Docker 實例來連接所有的 git 專案文件夾，每個 git 帳戶都有一個 Docker 實例，然後您可以使用相對應的 Docker 映像來使用您想要的帳戶進行提交。