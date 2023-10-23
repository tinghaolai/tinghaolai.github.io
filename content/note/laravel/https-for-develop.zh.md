---
title: "https-for-develop.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["laravel"]
---



# Translated by ChatGTP

## 開發時使用 HTTPS

需要使用 HTTPS 時的開發中 HTTPS 的執行方式。

### 非 Docker

* Windows: Apache (MAMP)
* 其他: Nginx (推薦)

### Docker

* 使用 php-fpm 容器 + Nginx 容器執行靜態 HTML 檔案
* 使用 php artisan serve + Nginx 容器代理到 Laravel Serve