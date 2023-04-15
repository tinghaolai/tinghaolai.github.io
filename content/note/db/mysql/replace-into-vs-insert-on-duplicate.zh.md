---
title: "replace-into-vs-insert-on-duplicate.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



# Translated by ChatGTP

## Replace into vs insert on duplicate

有幾種方式可以大量插入或更新資料。

* 刪除重複並批量插入
    * 如果插入是計算的且大小很大，無法放入 RAM，則無法知道要先刪除什麼。
* Replace into
* 批量插入和“批量重複”

使用 `replace into` 的一個問題是它實際上首先刪除重複的資料並插入新的資料，這將觸發刪除角落，可能不是原本想要執行的 SQL。

> 基於此，我沒有比較效能，除此之外，應該差不多。