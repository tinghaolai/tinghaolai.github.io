---
title: "when-will-index-lose-effect.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



---

## When will index column lose effect

### Operations on index column

`explain select * from target_tag where CONCAT('t', name) = 'etname';`

#### Fuzzy searching at beginning

`explain select * from target_tag where name like '%etnam%';`

> All

`explain select * from target_tag where name like 'tetnam%';`

> key type: `range`

### Using `OR` Condition

Not using index column in one of `OR` condition

### MySQL think scan all table is faster than using index

* searching result amount is relatively close to the whole table amount
    * using index need extra time to get data from index, which is I/O reading.

### Using not equal / is not null / not in / not exists

`explain select * from target_tag where name != 'testname';`



---

