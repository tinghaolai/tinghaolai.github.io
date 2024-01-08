---
title: "example_flexible_aws_ec2_namecheap_frp_docker_laravel.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["devOps","devOps-cloudflare","devOps-cloudflare-ssl"]
---



# Translated by ChatGTP

## 範例

* VM：aws ec2
* 本地機器：實際用來提供網絡服務的機器
* 代理：frp
* 域名：namecheap
* 網絡：docker laravel
* SSL：cloudflare - 可彈性選擇

## 設定

* DNS：在cloudflare上設定子域名指向aws ec2的IP/CNAME
* 在namecheap註冊Nameserver
* 本地機器使用frp客戶端從aws ec2代理網絡服務（docker laravel使用任意端口），將80端口映射到aws ec2
* 在瀏覽器中輸入https://開頭的子域名

> 瀏覽器（443）-> Namecheap -> cloudflare -> aws ec2（80）-> frp（80映射到8305）-> 本地機器（8305）-> docker laravel（8305映射到80）