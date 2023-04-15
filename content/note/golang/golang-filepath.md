---
title: "golang-filepath.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["golang"]
---



---


## Golang file path

When we're trying to load external files,

we can choose using relative path or absolute path.

but if we're using absolute path, 

we can't use the same code in different environment.

So I prefer to use relative path.

## Relative path problem

What I really want "relative" is relative to the current file.

but in fact, the relative path is relative to the current working directory.

So we have to constrain the folder when calling the code,

such as root of the project.

## Load static file

If we want to load static file, 

which won't change in different environment,

we can use getting absolute path of the current file.

then if we're run with `go run`, 

or combined binary file, 

we can always find the target file,

since the target file is also compiled into the binary file.

`getPackageDir` getting the absolute path of the current file,

if won't change no matter what folder we're in when calling the code.


### Example code

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

## Load dynamic file

After many ways trying, the final solution is using relative path and limit the code-calling folder.



```go
func loadEnv() {
	wd, _ := os.Getwd()

	envFile, _ := filepath.Abs(filepath.Join(wd, ".env"))

	if err := godotenv.Load(envFile); err != nil {
		panic(err)
	}
}

```

Let's think this way,

when we run `composer create-project laravel/laravel`,

the file create in the current folder,

so think `go run` or binary file as an application,

it read the related path from the current folder.

---

