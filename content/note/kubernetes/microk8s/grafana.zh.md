---
title: "grafana.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["kubernetes","kubernetes-microk8s"]
---



# Translated by ChatGTP

## 安裝

`microk8s.enable dashboard prometheus`

## 開啟儀表板

* `microk8s kubectl get all --all-namespaces` 取得 IP
* 在瀏覽器中開啟
* 輸入密碼
  * 從 `microk8s config` 取得
  * 如果不能使用，試試 `admin`，`admin`
  * 或者對我來說，用戶名：`admin`，`prom-operator` 也可以工作。