---
title: "how-much-data-is-too-big-for-table.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



# Translated by ChatGTP

## 結論

對於建議，限制約為 2 億行，實際行數取決於數據大小。

## 原因

由於 B+ 樹的存在，

行數越多，B+ 樹的層數也會越多，

這會導致更多的 I/O 讀取。