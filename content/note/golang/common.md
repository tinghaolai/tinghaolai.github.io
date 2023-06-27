---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["golang"]
---



---

## go.mod and go.sum

Why need `go.sum` when we have `go.mod`?

Because package dependencies in `go.mod` is not centralized. 

For example, I can release a go package from github,

and version `1.0.1`, 

someone downloaded it but then I deleted it.

Then I can release another branch code with same  version `1.0.1`,

This can cause problem.

so the solution is by checking hash in `go.sum`.

## UML class diagram tools

* https://www.dumels.com/


---

