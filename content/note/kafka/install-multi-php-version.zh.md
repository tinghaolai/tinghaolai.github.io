---
title: "install-multi-php-version.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["kafka"]
---



# Translated by ChatGTP

## 在Debian上為kafka安裝多個php版本

如果您想要為php安裝kafka，您可以執行 `pecl install rdkafka`，並在php.ini中設置庫，但是當您嘗試安裝多個php版本使用kafka時，只有一個版本能夠工作。

例如：

`PHP啟動：無法加載動態庫 'rdkafka.so'（嘗試過：/usr/lib/php/20190902/rdkafka.so（/usr/lib/php/20190902/rdkafka.so：無法打開共享對象文件：沒有此類文件或目錄），/`

> 20190902是php7.4的php庫文件夾，2021/09/02是php8.1

同時，您可以看到在另一個版本的庫文件夾中安裝了 `rdkafka.so`，在我的情況下是 `/usr/lib/php/20210902/rdkafka.so`。

所以您可以做的是：將pecl設置為另一個版本的php

```shell
sudo pecl config-set php_ini /etc/php/7.4/cli/php.ini
sudo pecl config-set ext_dir /usr/lib/php/20190902/
sudo pecl config-set bin_dir /usr/bin/
sudo pecl config-set php_bin /usr/bin/php7.4
sudo pecl config-set php_suffix 7.4
sudo pecl install -f rdkafka
```

> 不確定是否必須切換php版本，但這裡是命令 `sudo update-alternatives --config php`。

安裝完成後，您可以看到 `rdkafka.so` 安裝在正確的文件夾中，但原始文件夾中的庫已經被刪除，這可能是當kafka被重新安裝時由 `pecl` 刪除的。

並不知道正確的做法是什麼，但這是我的解決方案：

1. 將lib和 `php.ini` 中的 `rdkafa.so` 名稱都改為 `{--anything--}.so`。
2. 執行上述命令。