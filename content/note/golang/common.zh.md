---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["golang"]
---



# Translated by ChatGTP

## go.mod 和 go.sum

為什麼我們需要 `go.sum` 呢？`go.mod` 還不夠嗎？

這是因為 `go.mod` 中的套件依賴不是集中管理的。

例如，我可以從 GitHub 上發佈一個名為 `1.0.1` 的 go 套件，

有人下載了它，但然後我刪除了它。

接著我可以發布另一個分支的程式碼，版本號也是 `1.0.1`，

這樣可能就會造成問題。

因此，解決方法是通過在 `go.sum` 中檢查散列值來確定套件版本。