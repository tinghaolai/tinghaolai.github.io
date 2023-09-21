---
title: "ab-testing.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["backend-random-thing"]
---



# Translated by ChatGTP

## AB測試

寫一個簡單的PHP範例。

**錯誤的演算法**

```php
public function abWrong()
{
    $rates = [
        'a' => 5,
        'b' => 70,
        'c' => 25,
    ];

    $totalrate = 100;
    $counts = [
        'a' => 0,
        'b' => 0,
        'c' => 0,
    ];

    $totalTestCase = 10000000;
    for ($i = 0; $i < $totalTestCase; $i++) {
        foreach ($rates as $key => $rate) {
            $rollNumber = mt_rand(1, $totalrate);
            if (
                $key !== 'c' &&
                $rollNumber > $rate
            ) {
                continue;
            }

            $counts[$key]++;
            break;
        }
    }

    $countPercent = [];
    foreach ($counts as $key => $count) {
        $countPercent[$key] = $count / $totalTestCase;
    }

    dd($countPercent);
}
```

你可以打印出`$countPercent`的結果來看百分比是錯誤的。

### 思考1.

如果a和b都是10%的機率，c是80%的機率

正確的百分比應該要有20%的機率選到a或b，

但是在這個演算法中，它計算的是如果選到10%的機率兩次，

百分比將會是(100% - (沒有選到a和沒有選到b)) = 100 - 0.9 * 0.9 = 19.19%。

### 思考2.

如果a的機率是90%，

並且b和c都是5%的機率，

當a沒有被選到時，

b和c的機率應該是各50%，

而不是各5%。

**正確的演算法**

```php
public function abRight()
{
    $rates = [
        'a' => 5,
        'b' => 70,
        'c' => 25,
    ];

    $totalrate = array_sum($rates);
    $counts = [
        'a' => 0,
        'b' => 0,
        'c' => 0,
    ];

    $totalTestCase = 10000000;
    for ($i = 0; $i < $totalTestCase; $i++) {
        $accumulatedWeight = 0;
        $rollNumber = mt_rand(1, $totalrate);
        foreach ($rates as $key => $rate) {
            $accumulatedWeight += $rate;
            if (
                $rollNumber <= $accumulatedWeight
            ) {
                $counts[$key]++;
                break;
            }
        }
    }

    $countPercent = [];
    foreach ($counts as $key => $count) {
        $countPercent[$key] = $count / $totalTestCase * 100;
    }

    dd($counts, $countPercent);
}
```