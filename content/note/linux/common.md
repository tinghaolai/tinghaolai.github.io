---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["linux"]
---


## Linux common note

## High cpu usage

Sometimes can not find high CPU usage by `htop`.

Try `ps -eo pcpu,pid,user,args | tail -n +2 | sort -k1 -r -n | head -10`

## Turn off laptop monitor when openSSH

On some laptops, closing the monitor may cause the server to shut down, even though you don't need the monitor when running a Linux server on the laptop.

### Solution

* `sudo vim /etc/systemd/logind.conf`
* Add line
    * HandleLidSwitch=ignore
* Run `sudo systemctl restart systemd-logind`

### `Expect` command

Used for script command.

Example script after ssh login
```shell
expect -c '
spawn sudo su - root
expect "*"
send "password\n"
interact
```

---

