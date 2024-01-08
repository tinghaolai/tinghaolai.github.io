---
title: "frp_setting_example.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["linux","linux-frp"]
---



---

## FRP

### frps.ini

```shell
[common]
bind_port = 20022
# need to set token in frpc if set token in frps
token = 123456
```

### frpc.ini

```shell
[common]
server_addr = 00.000.000.000
server_port = 20022
tls_enable = true 
token = 123456

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

> fill `00.000.000.000` with actual IP from server running frps. 


### Note

Remember open ports if using cloud server, not inside the machine but setting from the cloud platform.

### Usage

Enter `00.000.000.000:6002` in browser to enter the web server hosted by frpc machine.


### Reload config

`systemctl restart frpc`



---

