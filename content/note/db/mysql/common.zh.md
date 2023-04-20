---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



# Translated by ChatGTP

## 空值

結論是不建議在 `MySQL` 數據庫中使用 `NULL`。

使用 `NULL` 值存在著很多問題，而且在不同的版本/情況下會有各種不同的問題。

可能會出現以下一些問題：

* 導致性能問題
* 導致排序錯誤
* 導致計算結果錯誤，例如 `count()`、`max()` 和 `min()`

## 終止運行時間過長的進程

運行 `Show processlist;` 並執行 `kill {processID}`。

## count(*) vs count(1) 的速度

在 `innodb` 中相同。