---
title: "less-join-as-you-can.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



---

## Conclusion

Use less joins as you can, use max 2 joins, 

check if small table drive big table is better than big table drive small table.

also check index used in the join relation. 

Don't join the tables they're having different responsibility.

## But why?

### Performance issue

More join use usage, more data scan, 

causing performance issue,

and operation like this need lots of memory,

will cause other data memory cache been deleted,

which will influence other business.

#### Mainly issue

* Hard to optimize query
* Hard to modify query if one of the table changed into such as: `table-split`, `database-split`, `table-structure-modify`


---

