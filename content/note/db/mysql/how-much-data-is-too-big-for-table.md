---
title: "how-much-data-is-too-big-for-table.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["db","db-mysql"]
---



---

## Conclusion

Limit about 200 million rows for recommendation,

the actual number of rows depends on the size of the data.

## Reason

Due to the B+ tree,

the more rows, 

the more layers b+ tree has,

and it means more I/O reading.



---

