---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



---

## IP allow in 8.0

* `GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;`
* `FLUSH PRIVILEGES;`

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

## Paginate

There's a problem with common pagination using limit and offset,

if size of offset is too large,

it will cause performance issue,

since it will scan all the rows before the offset.

One of solution is to use last of `id` to paginate.





---

