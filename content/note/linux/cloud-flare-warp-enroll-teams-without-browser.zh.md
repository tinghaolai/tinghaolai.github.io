---
title: "cloud-flare-warp-enroll-teams-without-browser.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["linux"]
---



# Translated by ChatGTP

warp-cli teams-enroll <TeamName>

會開啟瀏覽器 `https://<Team>.cloudflareaccess.com/warp`

但如果您的伺服器不是桌面呢？

## 解決方案

在已經登入同一帳戶的桌面瀏覽器中打開此網址，

從控制台複製 token。

回到您的伺服器，

運行 `warp-cli teams-enroll-token com.cloudflare.warp://<Team>.cloudflareaccess.com/auth?token=XXXXXXXXXXXXXXXXX`。