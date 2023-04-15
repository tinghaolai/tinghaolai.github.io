---
title: "golang-build-so.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["golang"]
---



## Golang建置插件 (Translated by ChatGTP)

通常，如果我們想要將一個函數寫到另一個檔案中，可以使用 `package`。

但在某些情況下，

我們要使用的函數和檔案是動態的，

此時，我們可以將這些檔案建置成 `.so` 檔案。

更具體來說，使用 `go build -buildmode=plugin`，

然後在主要檔案中使用這個 `.so` 檔案。

```go
plug, err := plugin.Open("extra.so")
runLib, err := plug.Lookup("targetFunctionName")
runLib.(func())()
```

## CI/CD問題

當我們將主命令建置為二進位檔案，

並且該二進位檔案讀取動態的 .so 檔案時，

但如果我們建置二進位檔案的平台與我們建置 .so 檔案的平台不同，

我們會得到 plugin not implemented 錯誤。

## 解決方案

確保我們建置二進位檔案的平台與我們建置 .so 檔案的平台相同，

例如使用 Docker。
