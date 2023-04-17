---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



# Translated by ChatGTP

## 關閉運行太久的進程

運行 `Show processlist;` 並使用 `kill {processID}` 關閉該進程。

## count(*) vs count(1) 速度比較

在 `innodb` 中速度相同。