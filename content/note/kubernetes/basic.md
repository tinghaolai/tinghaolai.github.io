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
* Try use `sudo -s` other than `sudo su -`

## Uninstall

```bash
kubeadm reset
sudo apt-get purge kubeadm kubectl kubelet kubernetes-cni kube*   
sudo apt-get autoremove  
sudo rm -rf ~/.kube
rm -rf /etc/kubernetes 
```

## Structure

![img.png](https://raw.githubusercontent.com/tinghaolai/just-random-note/master/kubernetes/img.png)

## Install guide recommend

* kubeadm
  https://skyao.io/learning-kubernetes/docs/installation/kubeadm/ubuntu.html
* MicroK8s
  * Just follow the official guide, 
  * but somehow it became extremely slow after re-installed, I did all what I can do to entirely remove, but still.
* For ubuntu, really helpful tutorial
  * https://www.youtube.com/watch?v=7k9Rdlx30OY

## About multiple machines cluster from frp

Maybe i do something wrong, 

but i can't make it work,

so don't try it,

just use regular way.





---

