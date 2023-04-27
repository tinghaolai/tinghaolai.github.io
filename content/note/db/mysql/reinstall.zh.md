---
title: "reinstall.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



# Translated by ChatGTP

## 重新安裝 MySQL

若你的 SQL 卡住了，可能需要運行多天，甚至回滾操作也需要很長時間，這時可以通過重新安裝 MySQL 來解決。

### 卸載卡住的 MySQL

當你在運行卡住的 SQL 時卸載 MySQL，`apt-get purge` 也會卡住。

**解決方案**

運行 `sudo killall -KILL mysql mysqld_safe mysqld`。

### 卸載腳本

* `apt-get purge --auto-remove mysql-common mysql-server mariadb-server`
* `apt-get autoremove`
* `apt-get autoclean`
* `rm -rm /var/lib/mysql`

### 安裝腳本

* `apt-get install --reinstall mysql-server`

### 設置

* 本地用戶權限
    * `ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';`
* 創建遠程用戶
    * `CREATE USER 'username'@'%' IDENTIFIED BY 'passwd';`
    * `GRANT ALL PRIVILEGES ON *.* TO 'username'@'%' IDENTIFIED BY 'passwd' WITH GRANT OPTION;`
    * `service mysql restart`
* 遠程設置
    * 位置 `vim /etc/mysql/mysql.conf.d/mysqld.cnf`
    * 端口 `port=23306`
    * 遠程 IP 地址 `bind-address=0.0.0.0`