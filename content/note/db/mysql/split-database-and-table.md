---
title: "split-database-and-table.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



---

## Why

If the size of table is too large,

the performance will be bad,

and one of solution is to split the table.

## How to split table

* Base on the different data type
* split by id

## Split by id

But how to find which database/table to get the row.

For example: If we have too many users, 

and we want to split it into multiple table.

And we have a user_id: `559342`,

how do we know which table `user_id:559342` in?

### Solution

* Using hash to know which table to store / read.
* Hash methods
  * Snowflake

---

