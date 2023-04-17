---
title: "only-select-columns-you-need.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



# Translated by ChatGTP

## 不要選取所有欄位

從資料表選擇所有欄位是一種不良的做法。最好只選擇您需要的欄位。

## 但為什麼呢？

在 `innoDB` 中，有兩種索引類型，即主索引和次要索引。如果我們請求次要索引中沒有的欄位，即使我們通過次要索引找到所需的行，我們仍然需要讀取主索引以獲取其他欄位，這就是為什麼如果您使用 `select *` 和 `select index_columns` 來解釋查詢，MySQL 優化器會選擇不同的查詢計劃的原因。