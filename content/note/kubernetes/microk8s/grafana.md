---
title: "grafana.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["kubernetes","kubernetes-microk8s"]
---



---

## Install

` microk8s.enable dashboard prometheus`

## Open dashboard

* `microk8s kubectl get all --all-namespaces` to get IP
* open in browser
* enter password
  * getting from `microk8s config`
  * if it doesn't work, try `admin`, `admin`
  * or for me, user: `admin`, `prom-operator` works

---

