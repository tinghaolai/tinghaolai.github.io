---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["git"]
---



# Translated by ChatGTP

## 忽略 Grep
git grep -i user-token -- ':!*.js'

## 在資料夾中搜尋
git grep -i user-token resources/js

## 排除資料夾
git grep -i identity -- ':!public'

## 將輸出寫入檔案
git grep -i identity -- ':!public' * > output-file

## 隱藏資料夾
git stash push -- public

## 軟重置
git reset --soft head^

## 儲存憑證資訊
git config credential.helper store