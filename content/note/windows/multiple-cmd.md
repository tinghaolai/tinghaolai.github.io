---
title: "multiple-cmd.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["windows"]
---



---

**multi-worker.bat**

```shell
@echo off

set count=1
set max=40

:loop
if %count% gtr %max% goto end

start cmd /k "cd c:/git/project && php artisan queue:work"
set /A count+=1
goto loop

:end

```

---

