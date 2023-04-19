---
title: "github-action-cicd_docker-compose_pull.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["cicd","cicd-github"]
---



---

## Permission problem

* sudo groupadd docker
* sudo gpasswd -a $USER docker
  * sudo gpasswd -a deploy docker
* newgrp docker
* sudo systemctl restart docker


---

