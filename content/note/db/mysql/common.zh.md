---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



# Translated by ChatGTP

## IP 允許在 8.0 中

* `GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;`
* `FLUSH PRIVILEGES;`

## 空值

總之，不建議在 `MySQL` 數據庫中使用 `NULL`。

`NULL` 值存在很多問題，在不同版本/情況中也有所不同。

以下是可能發生的一些問題。

* 導致性能問題
* 導致排序錯誤
* 導致計算（例如 `count()`、`max()`、`min()`）結果錯誤

## 終止運行時間過長的進程

運行 `Show processlist;` 和 `kill {processID}`。

## count(\*) vs count(1) 速度

在 `innodb` 中相同。

## 分頁

通過限制和偏移量進行常見的分頁存在問題，如果偏移量太大，將導致性能問題，因為它會掃描偏移量之前的所有行。

其中一個解決方案是使用最後的 `id` 進行分頁。