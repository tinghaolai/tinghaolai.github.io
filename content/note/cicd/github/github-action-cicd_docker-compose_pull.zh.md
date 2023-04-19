---
title: "github-action-cicd_docker-compose_pull.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["cicd","cicd-github"]
---



# Translated by ChatGTP

## 權限問題

* sudo groupadd docker
* sudo gpasswd -a $USER docker
  * sudo gpasswd -a deploy docker
* newgrp docker
* sudo systemctl restart docker