---
title: "golang-build-so.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["golang"]
---



---

## Golang build plugin

Normally, if we want to write a function to another file, 

we can use `package` to do that.

But in some case,

The function and file we're going to use is dynamic,

In this case, we can build these file's into `.so` file.

More specific: `go build -buildmode=plugin`,

and use this file in our main file.

```go
plug, err := plugin.Open("extra.so")
runLib, err := plug.Lookup("targetFunctionName")
runLib.(func())()
```

## CI/CD problem

When we're building main command into binary file,

and the binary file reading the dynamic `.so` file,

But if the platform we build the binary file is different 

from the platform we're build the `.so` file,

we will get the error `plugin not implemented`.

## Solution

Make sure the platform we're build the binary file is 

the same as the platform we're build the `.so` file,

such as Docker

---

