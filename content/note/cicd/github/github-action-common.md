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

### Make register image public

I really thank Github's UI/UX team, 

they make me to write this note,

for just modify the visibility.

1. Go to the organization page
    * Not from the repository page
    * Note from the package page
2. Go to the settings page
    * If there's no `packages` tab, then you're in th wrong place
3. Then modify the visibility

Or just replace the url like this

https://github.com/organizations/streaming-data-architecture/settings/packages



---

