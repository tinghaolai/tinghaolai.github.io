---
title: "arch-linux-install-cloudflare-zero-trust.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["linux"]
---



# Translated by ChatGTP

## 在 Arch Linux 安裝 Cloudflare Zero Trust

安裝完 `warp-cli` 後，若在註冊客戶端時遇到 `Cloudflare warp daemon error` 錯誤，可以執行 `systemctl enable --now warp-svc.service` 命令。

### 客戶端註冊團隊

* 執行命令 `warp-cli teams-enroll {團隊名稱}`
* 用 Firefox 開啟 URL，從控制台複製 token，token 以 "com.cloudflare.warp://" 開頭。
* 執行命令 `warp-cli status`
* 執行命令 `warp-cli start`
* 執行命令 `warp-cli status`