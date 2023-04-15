---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["operating_system","operating_system-macOS"]
---



# Translated by ChatGTP

### 檢查開啟的埠口
netstat -ap tcp | grep -i "listen"

### 重啟 MySQL
mysql.server restart

### 檢查 SSH 是否開啟
sudo systemsetup -getremotelogin

### 查找本地 IP
ifconfig | grep "inet " | grep -Fv 127.0.0.1 | awk '{print $2}'

### 切換為超級使用者
sudo su

### 重啟 Nginx
sudo pkill nginx && sudo nginx