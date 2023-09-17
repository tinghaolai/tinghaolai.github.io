---
title: "build.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["flutter"]
---



# Translated by ChatGTP

# Android

## 安裝 APK 失敗

* 檢查 `android/app/build.gradle` 中的 `versionCode`
  * 最近不太清楚為什麼，我必須將 versionCode 增加到 3 才能在 Play Store 上發佈，
    成功上傳，但無法從 apk 安裝到我的手機上。