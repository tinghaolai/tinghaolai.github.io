---
title: "github-action-common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["cicd","cicd-github"]
---



---

### the input device is not a TTY

example: 
```bash
docker-compose exec app /app/migrate migrate
```

Solution: add `-T` to the command
```bash
docker-compose exec -T app /app/migrate migrate
```



---

