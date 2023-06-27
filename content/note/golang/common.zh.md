---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["golang"]
---



# Translated by ChatGTP

## go.mod 和 go.sum

為什麼我們需要 `go.sum`，而已經有了 `go.mod` 呢？

因為 `go.mod` 中的套件相依性並不是集中管理的。

舉個例子，我可以從 GitHub 發布一個 Go 套件，

版本是 `1.0.1`，

有人下載了該套件，但後來我刪除了它。

然後，我可以發布另一個分支代碼，版本仍然是 `1.0.1`，

這可能會引起問題。

因此，解決方法是通過在 `go.sum` 中檢查哈希值。

## UML 類別圖工具

* https://www.dumels.com/