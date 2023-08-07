---
title: "data-store.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["backend-random-thing"]
---



# Translated by ChatGTP

## 浮動貨幣

永遠不要使用浮點數/雙精度來存儲金錢，或任何需要精確的東西。

相反，使用 'string' 類型，例如 'decimal(10,2)'。