---
title: "only-select-columns-you-need.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



---

## Don't select all columns

Selecting all columns from a table is a bad practice. It is better to select only the columns you need.

## Buy why

In `innoDB`, there are two type of indexes,

primary index and secondary index.

If we ask the columns not in the secondary index, 

event we found the row we want by secondary index,

we still need to read the primary index to get other columns,

that's why if you `explain` the query with `select *` and `select index_columns`,

and why mysql optimizer choose different query plan. 


---

