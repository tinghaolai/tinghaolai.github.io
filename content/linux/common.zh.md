---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
---



## Linux 常用記事 (Translated by chatGPT)

## CPU 負載過高

有時候 `htop` 無法找出 CPU 負載過高的問題。

嘗試使用以下指令：`ps -eo pcpu,pid,user,args | tail -n +2 | sort -k1 -r -n | head -10`

## 開啟 SSH 時關閉筆電螢幕

在某些筆電上，關閉筆電螢幕可能會導致伺服器關閉，但實際上在運行 Linux 伺服器時不需要螢幕。

### 解決辦法

* `sudo vim /etc/systemd/logind.conf`
* 新增以下內容：`HandleLidSwitch=ignore`
* 執行 `sudo systemctl restart systemd-logind`

### `Expect` 指令

用於腳本命令。

以下是一個 SSH 登入後的範例腳本：

```shell
expect -c '
spawn sudo su - root
expect "*"
send "password\n"
interact