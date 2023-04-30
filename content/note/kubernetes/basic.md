---
title: "basic.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["kubernetes"]
---



---

## Basic

After installed, run `kubectl get node`,

should be only one node, which is master node with role: `control-plane`.

## Errors

* Get `The connection to the server 192.168.1.101:6443 was refused - did you specify the right host or port?`
  * Try `sudo kubeadm init --pod-network-cidr=10.244.0.0/16 -v=9`
  * or `systemctl restart kubelet`
* Got too many errors event just start using it.
  * Tru another way, such as `MicroK8s`

## Uninstall

```bash
kubeadm reset
sudo apt-get purge kubeadm kubectl kubelet kubernetes-cni kube*   
sudo apt-get autoremove  
sudo rm -rf ~/.kube
```

## Structure

![img.png](https://raw.githubusercontent.com/tinghaolai/just-random-note/master/kubernetes/img.png)

## Install guide recommend

https://skyao.io/learning-kubernetes/docs/installation/kubeadm/ubuntu.html

---

