---
title: "multiple-cmd.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["windows"]
---



# Translated by ChatGTP

**multi-worker.bat**

```shell
@echo off

設置 count 為 1
設置 max 為 40

:loop
如果 count 大於 max，則轉向結束

啟動一個新的命令提示符窗口 "cd c:/git/project && php artisan queue:work"
設置 count 等於 count 加 1
跳轉到 loop

:end

```