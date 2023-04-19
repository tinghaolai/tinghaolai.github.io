---
title: "frp_setting_example.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["linux"]
---



# Translated by ChatGTP

## FRP

### frps.ini

```shell
[common]
bind_port = 20022
```

### frpc.ini

```shell
[common]
server_addr = 00.000.000.000
server_port = 20022
tls_enable = true 

[ssh]
type = tcp
local_ip = 127.0.0.1
local_port = 20022
remote_port = 6000

[web-dev]
type = tcp
local_port = 23306
remote_port = 6001

[web-dev-v2]
type = tcp
local_port = 20080
remote_port = 6002
```

> 將 `00.000.000.000` 替換為實際執行 frps 的伺服器 IP。


### 注意事項

如果使用雲端伺服器，請記得開啟埠口，不是在電腦內設定，而是由雲端平台設定。

### 使用方法

在瀏覽器中輸入 `00.000.000.000:6002` 以進入由 frpc 機器主機托管的網頁伺服器。