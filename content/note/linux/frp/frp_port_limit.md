---
title: "frp_port_limit.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["linux","linux-frp"]
---



---

## Port binding limit

Without root, you can only bind ports above 1024. 

## Solution

* `vim /etc/systemd/system/frps.service`
* `systemctl daemon-reload`
* check port binding in dashboard


---

