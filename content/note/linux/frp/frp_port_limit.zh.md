---
title: "frp_port_limit.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["linux","linux-frp"]
---



# Translated by ChatGTP

## 端口绑定限制

如果没有超级用户权限，只能绑定大于1024的端口。

## 解决方案

* `vim /etc/systemd/system/frps.service`
* `systemctl daemon-reload`
* 在仪表盘中检查端口绑定情况