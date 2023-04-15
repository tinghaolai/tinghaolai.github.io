---
title: "mysql-account-permission.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["google-cloud"]
---



# Translated by ChatGTP

## 帳戶權限

首先，永遠不要相信 ssh 工具的複製功能，它甚至可能什麼都不貼上，這就是我在這個問題上遇到更多問題的原因。

## 修改 mysql 標誌

不要嘗試透過 ssh 實例來修改權限，因為這是不允許的，完全浪費你的時間。
您應該執行 gcloud 命令，例如 `gcloud sql instances patch dave-test --database-flags sql_mode=TRADITIONAL`,
或者使用 GUI 來修改 DB 設置。 （位於 instance/edit 中）