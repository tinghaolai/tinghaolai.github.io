---
title: "split-database-and-table.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



# Translated by ChatGTP

## 為什麼

如果表格的大小太大，

性能會變差，

其中一種解決方法是將表格分割。

## 如何分割表格

* 基於不同的數據類型
* 按 id 分割

## 按 id 分割

但如何找到需要獲取行的數據庫/表格。

例如：如果我們有太多用戶，

我們想將其分割成多個表格。

並且有一個用戶 ID：`559342`，

我們如何知道 `user_id:559342` 在哪個表中？

### 解決方案

* 使用哈希來知道要存儲/讀取哪個表。
* 哈希方法
  * Snowflake