---
title: "when-will-index-lose-effect.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



# Translated by ChatGTP

## 當索引欄位失效時

### 索引欄位的操作

`explain select * from target_tag where CONCAT('t', name) = 'etname';`

#### 在開頭進行模糊搜尋

`explain select * from target_tag where name like '%etnam%';`

> 全部的

`explain select * from target_tag where name like 'tetnam%';`

> 鍵值類型：`range`

### 使用 `OR` 條件

在一個 `OR` 條件中不使用索引欄位

### MySQL 認為掃描整個表比使用索引更快

* 搜尋結果的數量與整個表的數量相對接近
    * 使用索引需要額外的時間來從索引中獲取數據，這是 I/O 讀取。

### 使用不等於 / 不為空 / 不包含 / 不存在

`explain select * from target_tag where name != 'testname';`