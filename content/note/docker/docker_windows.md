---
title: "docker_windows.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["docker"]
---



---

## Volume to Absolute folder

```yaml
version: '3'
services:
  php-fpm:
    image: php:8-fpm
    ports:
      - 9004:9000
    volumes:
      - "C:/git:/var/www/html"
    container_name: personal-git
```


---

