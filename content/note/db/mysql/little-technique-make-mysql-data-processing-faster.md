---
title: "little-technique-make-mysql-data-processing-faster.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



---

## Static Data Calculation

Let's say we have a table called `raw`. 

We want to calculate some static data from the `raw` table,

and write the result into the `result` table every day at midnight.

During the calculation process, 

if some result is reusable, 

we can store the intermediate result in the `temp` table,

so that we can calculate the result faster the next time.

## Synchronization Diff Problem

Every time we recalculate the result, 

we need to sync the difference from the `raw` table to the `temp` table,

and then continue to calculate the result. 

However, the `diff sync` may be very slow if there is a lot of data in the `raw` table 

when there is an index in the `temp` table. 

Because every operation will cause the index to be updated, 

and it will update the record one by one.

## Solution

Before syncing the diff, 

we can drop the index in the `temp` table, 

sync the diff, 

and then recreate the index again after syncing. 

In this way, we can calculate the new index much faster.

> This approach is not suitable for all situations, 
> 
> it fits in this situation since only this feature is using the table.
> 
> If there is another feature using the table,
> 
> we can't drop the index in the table, 
> 
> or the searching speed will be very slow.

---

