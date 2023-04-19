---
title: "handle_cross_time_zone.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["programming"]
---



---

## Handle cross time zone

### Solution

If we have clients in different timezone, 
how should we store the time data?

I think the best way is that client side way convert to timestamp, 
or give timezone info in API.

### Other approach

But what if another team member suggest just store origin datetime or timestamp?

Would that cause problems? and how to prove that?

### Store datetime

If current time is `GTM +0 2023-02-09 00:00:00`,
and we have two client in two timezones, 
which are `GMT +2` and `GMT +3`.

If we just store the datetime from client-side,
We will store two different datetime even both of them are the same time.

### Store timestamp by serverside

If current time is `GMT +0 2023-02-09 00:00:00`, which timestamp is `1675900800`

The client timezone is `GMT +2 2023-02-09 02:00:00`, which timestamp is also `1675900800`

And the server timezone is `GMT +8 2023-02-09 08:00:00`, which timestamp is also `1675900800`

If we convert the client-side to timestamp, in GMT + 0 would be `1675908000`, and GMT + 8 would be `1675879200`

When the client send another API want to get data stored in `GMT +2 2023-02-09 02:00:00`,
In server parse this datetime with GMT +8,

example: php 
```php
dd(Carbon::parse('2023-02-09 02:00:00')->timestamp);
```

Result timestamp would be `1675879200`, which is correct timestamp stored.

But, we'll still have two problems,
1. We can't display the correct datetime unless we know both server and client timezone and how data stored. It can be complex.
2. Still have wrong time if client have multiple timezones.
   * if we have another client with GMT + 3, the timestamp would be `1675882800`, which timestamp won't match if get data API called from GMT + 2.







---

