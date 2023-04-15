---
title: "Laravel-whereIn-query-slow.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["laravel"]
---



# Translated by ChatGTP

### Laravel whereIn 資料查詢緩慢

當我使用一個簡單查詢時，例如：

```php
$memberIds = $this->getMemberIds();
$members = MemberData::whereIn('member_id', $memberIds)->get();
```

當 `$memberIds` 陣列長度為 `5000` 時，查詢耗費約 30 秒，但若使用 GUI (datagrip) 查詢，則只需 1 秒。

這是 PDO 的問題。 https://bugs.php.net/bug.php?id=80027


### 解決方案

使用原生 SQL，但由於 SQL 注入問題，最好使用整數欄位。