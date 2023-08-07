---
title: "index-condition-pushdown.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



---

## Index Condition Pushdown (ICP)

ICP opened automatically

## What does ICP do?

example:

Table:
```sql
create table t(
    a int,
    b int,
    key(a,b)
);
```

Query: 
```sql
select * from t where a like 'A%' and b=2;
```

### Without ICP
MySQL scan rows match `a like 'A%'` first using `secondary index`,
then scan rows match `b=2` using `primary index`.

### With ICP
MySQL scan rows match `a like 'A%' and b=2` using `secondary index`,
since `b` is the second column of `key(a,b)`.

---

