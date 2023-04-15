---
title: "repository-patterm.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["design_pattern"]
---



# Translated by ChatGTP

## 儲存庫模式

我們需要一個模型來獲取機密數據。

```php
<?php

namespace App\Models;

class SecretModel
{
    public function find($id)
    {
        return [
            'id'      => $id,
            'title'   => '我的秘密',
            'content' => 'testName',
            'own'     => '秘密擁有者',
        ];
    }
}

```

我們使用一個服務從模型中獲取用戶秘密。

```php

<?php

namespace App\Services;

use App\Models\SecretModel;
use Illuminate\Support\Facades\Log;

class SecretService
{
    public function getSecret($id, $userType)
    {
        return (new SecretModel())->find($id);
    }
}
```


如果有三種類型的用戶，可以是付費用戶、普通用戶和訪客用戶。

* 普通用戶和訪客用戶無法獲取機密內容。
* 如果訪客想要一個秘密，必須記錄事件。

我們可以修改函數，像這樣。

```php
    public function getSecret($id, $userType)
    {
        $secret = (new SecretModel())->find($id);
        if ($userType === 'guest' || $userType === 'normal') {
            $secret['content'] = '***';
        }

        if ($userType === 'guest') {
            Log::info('guest want secret' . $id);
        }

        return $secret;
    }
```

這段代碼有兩個問題。

* 如果邏輯變得更加復雜，維護起來很困難。
* 不僅僅是獲取用戶數據，還需要閱讀代碼以了解會發生什麼事情。

所以，也許我可以嘗試這種方式。

```php
    public function getSecret($id, $hideSensitive, $logEvent)
    {
        $secret = (new SecretModel())->find($id);
        if ($hideSensitive) {
            $secret['content'] = '***';
        }

        if ($logEvent) {
            Log::info('guest want secret' . $id);
        }

        return $secret;
    }

    public function guestGetSecret($id)
    {
        return $this->getSecret($id, true, true);
    }
```

* 我可以通過參數命名知道發生了什麼事情，但仍然很難閱讀，特別是如果這個函數變得越來越大，例如有50個參數。
* 這個函數仍然在做更多的事情，不僅僅是獲取秘密。

讓我們把條件邏輯拆分為用戶類型函數。

```php
<?php

namespace App\Services;

use App\Models\SecretModel;
use Illuminate\Support\Facades\Log;

class SecretService
{
    public function getSecret($id)
    {
        return (new SecretModel())->find($id);
    }

    public function convertSensitive($secret)
    {
        $secret['content'] = '***';

        return $secret;
    }

    public function logEvent($id, $userType)
    {
        Log::info($userType . ' want secret ' . $id);
    }

    public function paidUserGetSecret($id)
    {
        return $this->getSecret($id);
    }

    public function normalUserGetSecret($id)
    {
        $secret = $this->getSecret($id);

        return $this->convertSensitive($secret);
    }

    public function guestGetSecret($id)
    {
        $this->logEvent($id, 'guest');
        $secret = $this->getSecret($id);

        return $this->convertSensitive($secret);
    }
}
```
 
 這樣，我可以輕鬆地知道在什麼情況下，什麼事情會被執行。
 
 這些函數的差異
 * 條件邏輯（業務邏輯） - 知道該做什麼
    * `paidUserGetSecret`、`normalUserGetSecret`、`guestGetSecret`
 * 儲存庫 - 實際代碼是什麼
    * 其他

當然，我可以混合儲存庫和業務邏輯，但是我認為像這樣的編程風格更易於維護。

最後一步：

* 將函數分為儲存庫類和服務類
* 實現界面使儲存庫可替換（服務中使用的功能）
* 減少依賴項
    * 儲存在庫中的代碼，例如SecretModel
    * 創建儲存庫
    
```php
    public function register()
    {
        $this->app->bind(SecretInterface::class, function($app)
        {
            return new SecretRepository(new SecretModel());
        });
    }
```


## 錯誤的範例

這是使用了儲存庫的服務功能，但方法不當。

```php
public function index($memberId, $perPage)
{
        $collectionCondition = $this->feeRepository->subjectCondition('allArrangeDeposit');
        $fee = $this->feeRepository->depositQuery()
            ->validDeposit()
            ->search([
                'memberId' => $memberId,
            ])
            ->search($collectionCondition);

        $fees = $this->paymentRepository->depositQuery()
                    ->search($this->paymentRepository->searchCondition('arrangeDeposit'))
                    ->search([
                        'memberId' => $memberId,
                    ])
                    ->unionAll($fee)
                    ->where('member_id', $memberId)
                    ->orderByDesc('created_at')
                    ->paginate($perPage);

        $fees->getCollection()->transform(function ($fee) {
            $repository = ($fee->unionType === 'payment') ? $this->paymentRepository : $this->feeRepository;
            return $repository->convertApiFormat($fee);
        });

        return $fees;
}
```

#### 依賴問題

```php
        $fee = $this->feeRepository->depositQuery() // 獲取查詢構建器
            ->validDeposit()                 // 在服務中調用模型函數（作用域）
```


* 通過儲存庫接口無法知道可以使用哪些函數
* Model與這個儲存庫有依賴關係，一旦切換儲存庫，儲存庫必須返回具有相同功能的對象

> 解決方案 - `depositQuery()`返回self，將模型存儲在self實例中，調用`validDeposit()`來操作模型

#### 沒有分離業務邏輯

```php
    ->validDeposit()
```

* 我知道它的意思，但我不知道它做了什麼
* 如何確定存款是有效的？
* 如果有效存款確定已更改，我該如何更改？

解決方案

```php

// service function
public function validDeposit()
{
// 可能有一個搜尋條件、回調函數，或者最好的方式是在儲存庫中搜索模型 
    return $this->feeRepository->successStatus();
}

```

>業務邏輯作為一個函數，它們所做的只是搜索`where('status', 'success')`，但專注於業務邏輯，這個函數名稱就是業務邏輯，並且必須知道這個邏輯實際上是什麼，在這種情況下，這是搜索成功狀態，但服務調用儲存庫時不需要知道如何編寫代碼，只需調用即可。


#### 還沒有分離業務邏輯
```php
    return $repository->convertApiFormat($fee);
```

* 儲存庫不應該知道api格式是什麼

解決方案
* `convertApiFormat()`應該是服務函數
* 編寫`convertFullInfo()`或`getPayerName()`函數，服務調用這些函數來結合api信息。