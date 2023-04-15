---
title: "eloquent-relation-lastOfMany.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["laravel"]
---



# Translated by ChatGTP

## Eloquent 關聯中的最後一個

在某些情況下，我們需要搜索“只有最後一個記錄”的關聯，例如，我想搜索每位最後一筆訂單價格大於1000元的客戶。

如果我們像這樣搜索：

```php
Member::whereHas('lastOrder', function ($query) {
    $query->where('money', '>', 1000);
})


function lastOrder() {
    return $this->hasOne(Order::class)->orderByDesc('id');
}
```

帶有條件的關聯會告訴sql去搜索`最後一筆訂單的價格大於1000`，這不是我們想要的。

因此，為了實現這個搜索條件，我們可以讓關聯表自己加入自己以選擇最大(id)來使用它，並且在laravel8之後，有一個更方便的功能叫做[latesetOfMany](https://laravel.com/api/9.x/Illuminate/Database/Eloquent/Relations/Concerns/CanBeOneOfMany.html#method_latestOfMany)，這使您不需要自己編寫連接。

### 效率

如果您dump出它的sql如何工作，您會意識到這個解決方案在數據量巨大時可能會很慢，因此也許我們可以嘗試其他解決方案，例如修改db架構以在主表中存儲最後一筆記錄的id，雖然有很多事情需要考慮。