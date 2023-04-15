---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["operating_system","operating_system-macOS"]
---



---


### Check open port
netstat -ap tcp | grep -i "listen"


### mysql restart
mysql.server restart

### check if ssh on
sudo systemsetup -getremotelogin

### Find local Ip
ifconfig | grep "inet " | grep -Fv 127.0.0.1 | awk '{print $2}'

### su
sudo su

### nginx restart

sudo pkill nginx && sudo nginx


---

