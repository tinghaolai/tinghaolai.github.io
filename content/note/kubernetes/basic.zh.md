---
title: "basic.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["kubernetes"]
---



# Translated by ChatGTP

## 基本

安裝完成後，執行 `kubectl get node`，

應該只有一個節點，這是具有角色 "控制平面" 的主節點。

## 錯誤

* 獲得 `The connection to the server 192.168.1.101:6443 was refused - did you specify the right host or port?`
  * 嘗試 `sudo kubeadm init --pod-network-cidr=10.244.0.0/16 -v=9`
  * 或者 `systemctl restart kubelet`
* 只是開始使用就獲得太多錯誤。
  * 嘗試其他方法，例如 `MicroK8s`

## 卸載

```bash
kubeadm reset
sudo apt-get purge kubeadm kubectl kubelet kubernetes-cni kube*   
sudo apt-get autoremove  
sudo rm -rf ~/.kube
```

## 結構

![img.png](https://raw.githubusercontent.com/tinghaolai/just-random-note/master/kubernetes/img.png)

## 安裝指南建議

https://skyao.io/learning-kubernetes/docs/installation/kubeadm/ubuntu.html