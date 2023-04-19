---
title: "github-action-common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["cicd","cicd-github"]
---



# Translated by ChatGTP

### 輸入設備不是TTY

範例： 
```bash
docker-compose exec app /app/migrate migrate
```

解決方法：在指令中加入 `-T` 
```bash
docker-compose exec -T app /app/migrate migrate
```