---
title: "golang-filepath.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["golang"]
---



## Golang檔案路徑 (Translated by ChatGTP)

當我們嘗試載入外部檔案時，

我們可以選擇使用相對路徑或絕對路徑。

但是如果我們使用絕對路徑，

我們無法在不同的環境中使用相同的程式碼。

因此我傾向使用相對路徑。

## 相對路徑問題

我真正想要的「相對路徑」是相對於當前檔案。

但事實上，相對路徑是相對於當前工作目錄。

因此我們必須限制資料夾，以在呼叫程式碼時使用，例如專案根目錄。

## 載入靜態檔案

如果我們想要載入靜態檔案，

在不同的環境中不會改變，

我們可以使用當前檔案的絕對路徑。

然後，如果我們運行 `go run`，

或結合二進位檔案，

我們始終可以找到目標檔案，

因為目標檔案也編譯到二進位檔案中。

`getPackageDir` 獲取當前檔案的絕對路徑，

如果在呼叫程式碼時，無論身處哪個資料夾，都不會改變。

### 範例程式碼

```go
package govel_migration

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

func Make(fileName string) {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	migrationPath := path.Join(cwd, "migrations")

	if err := os.MkdirAll(migrationPath, os.ModePerm); err != nil {
		panic(err)
	}

	location, _ := time.LoadLocation("Asia/Taipei")
	prefix := time.Now().In(location).Format("2006_01_02_150405")

	filePath := path.Join(migrationPath, prefix+"_"+fileName+".go")
	file, err := os.Create(filePath)

	if err != nil {
		fmt.Println("migration file generate fail: " + filePath)
		return
	}

	content := loadStub(path.Join(getPackageDir(), "migration.stub"))
	content = strings.Replace(content, "{UpFunctionName}", "Up"+toCamelCase(fileName), -1)
	content = strings.Replace(content, "{DownFunctionName}", "Down"+toCamelCase(fileName), -1)

	file.WriteString(content)

	fmt.Println(successMessage(filePath + " created"))
}

func loadStub(fileName string) string {
	content, _ := os.ReadFile(fileName)

	return string(content)
}

func getPackageDir() string {
	_, file, _, ok := runtime.Caller(0)

	if !ok {
		panic("failed to get file path")
	}

	return path.Dir(file)
}

```

讓我們這樣想，

當我們運行 composer create-project laravel/laravel 時，

該文件將在當前文件夾中創建，

因此，將 go run 或二進制文件視為應用程序，

它將從當前文件夾讀取相關路徑。