---
title: "mysql-non-strict-mode.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["google-cloud"]
---



# Translated by ChatGTP

## 非嚴格模式

轉換預設值並僅禁用完全的 GROUP BY。

```
SELECT @@sql_mode;
```
> STRICT_TRANS_TABLES，NO_ENGINE_SUBSTITUTION