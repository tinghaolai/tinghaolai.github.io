---
title: "mysql-isolate-level.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



# Translated by ChatGTP

## MySQL 隔離級別

### 尋找當前的隔離級別

5.7
> select @@tx_isolation;

8.0
> SELECT @@TRANSACTION_ISOLATION;

預設值
> REPEATABLE-READ

### 交易操作

```sql
begin;
select * from test.members where id = 3;
update test.members set type = 9999 where id = 3;
rollback;
commit;
```

### 測試交易級別

```sql
set session transaction isolation level read uncommitted;
```

### 隔離級別

* read uncommitted
  * 查看另一個 session 未 commit 或 rollback 時的資料更改。
  * 需要在讀取資料的 session 上更改隔離級別，而非更新資料的 session 上更改。
* read committed
  * 能夠讀取資料，但資料將不會更改。
* repeatable read
  * 只能在交易開始時讀取資料，即使其他 session 提交更新，也不能讀取到更改的資料。

> 所有隔離級別均無法解決非同步資料操作問題。

* serializable
  * 鎖定其他 session 資料更改，即使只選擇資料。

> 可以解決非同步操作問題，但可能引起大量鎖等待。