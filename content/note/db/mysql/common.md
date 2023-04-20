---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



---

## Null value

In conclusion, don't recommend to use `NULL` in `MySQL` database.

There's a lot of problem with `NULL` value,

but various in different version/situation.

There are some problems might happen.


* Causing performance issue
* Causing wrong sort
* Causing wrong result for calculation such as `count()`, `max()`, `min()`

## Kill process that execute too long

run `Show processlist;` and `kill {processID}`

## count(*) vs count(1) speed

same in `innodb`




---

