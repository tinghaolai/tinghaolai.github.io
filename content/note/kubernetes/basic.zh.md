---
title: "basic.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["kubernetes"]
---



# Translated by ChatGTP

## 基本設置

安裝完成後，運行 `kubectl get node`，

這可能只是具有角色 `control-plane` 的主控節點。

## 錯誤

* 獲取 `The connection to the server 192.168.1.101:6443 was refused - did you specify the right host or port?`
  * 嘗試 `sudo kubeadm init --pod-network-cidr=10.244.0.0/16 -v=9`
  * 或 `systemctl restart kubelet`
* 初次使用時出現太多錯誤。
  * 嘗試其他方式，如 `MicroK8s`。

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

* kubeadm
  https://skyao.io/learning-kubernetes/docs/installation/kubeadm/ubuntu.html
* MicroK8s
  * 只需遵循官方指南，但重新安裝後它變得非常緩慢，我嘗試完全刪除它，但仍然有問題。

## 關於來自frp的多台機器叢集

也許我做錯了什麼，

但我無法讓它工作，

所以不要嘗試，

只使用常規方法。