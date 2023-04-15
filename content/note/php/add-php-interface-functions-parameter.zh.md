---
title: "add-php-interface-functions-parameter.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["php"]
---



# Translated by ChatGTP

## 新增php接口函數參數

假設您有100個類別擴展同一個接口，某個類別的函數需要新增參數以實現該函數，這可能是一件很痛苦的事情，因為您需要修改接口函數並向每個類別擴展該接口的函數添加參數。

這裡有一種解決方法，您只需要在函數中添加 `$parameter=null` 參數即可，但是這樣做會失去接口的意義，使用者不知道如何僅通過接口函數使用此函數，每個類別函數的參數含義和參數類型可能不同。

因此，我的結論是，只需修改每個類別的函數。