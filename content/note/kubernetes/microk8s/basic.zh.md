---
title: "basic.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["kubernetes","kubernetes-microk8s"]
---



# Translated by ChatGTP

## 開啟儀表板

* 使用`microk8s kubectl get all --all-namespaces`命令來獲取 `service/kubernetes-dashboard` 的 IP。
* 在瀏覽器中輸入：`https://{ip}`。
* 登入使用令牌
  * `token=$(microk8s kubectl -n kube-system get secret | grep default-token | cut -d " " -f1)
microk8s kubectl -n kube-system describe secret $token`。