---
title: "mysql-count-timerange-in-different-timezone.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



---

## Group by time range

Let's say we have a table called `target_table`,

and we want to calculate the total records of different time range from the `target_table` table,

What we want to group by is the `created_at` column (which is store the unix time), 

and convert to its own group.

For example,

if we want to every 10 minutes into one group and count the total records,

`10:00:00` to `10:09:59` is one group,

`10:10:00` to `10:19:59` is another group,

What to do is that we can do something like this

`floor(created_at / 60 / 10) * 60 * 10)`

which `10` to be the minutes we want to group by.

you can also calculate the whole day,

for example,

`floor(created_at / 60 / 1440) * 60 * 1440)`

## Timezone problem

If your system have multiple timezone,

for example,

this customer's data is in `+8:00` timezone,

and another customer's data is in `+0:00` timezone,

What's the problem if we use all the same query?

Let's say we want to count the total records of all the day,

and there's two table,

`custom_1.target_table`, 

`custom_2.target_table`,

the timezone of `custom_1.target_table` is `+8:00`,

and the timezone of `custom_2.target_table` is `+0:00`,

At the `0` of unix time, 

`+0:00` is `1970-01-01 00:00:00`,

`+8:00` is `1970-01-01 08:00:00`,

so in this query,

`floor(created_at / 60 / 1440) * 60 * 1440)`

the time range we calculate for timezone in `+0:00` is 

`1970-01-01 00:00:00` to `1970-01-01 23:59:59`,

but for timezone in `+8:00` is 

`1970-01-01 08:00:00` to `1970-01-01 23:59:59`,

That's not what we want.

## Solution

Add bias base on timezone,

in query below,

`28800` is the bias of `+8:00` timezone,

and '+8:00' is the timezone we want to convert to.


```sql
SELECT 
created_at as origin,
date_format(
    CONVERT_TZ(
        FROM_UNIXTIME(
    floor((created_at + 28800) / 60 / 1440) * 60 * 1440), 
        '+8:00', '+0:00'), 
    '%Y-%m-%d %H:%i:%s') 
as range_start,
FROM target_table
order by id ;
```


---

