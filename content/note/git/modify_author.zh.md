---
title: "modify_author.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["git"]
---



# Translated by ChatGTP

## 修改作者

### 使用情境

提交的作者無法與 Github 帳戶連結，即便已經確定提交作者的使用者和電子郵件資訊是正確的，但仍然無法使用。很可能是在切換作者帳戶設置後，我將 Git 設定弄壞了。

### 解決方案

`git rebase -i sha`

`git commit --amend --author="Author <email@gamil.com>" --no-edit`
> 每次重新設置都要使用

`git rebase --committer-date-is-author-date SHA`