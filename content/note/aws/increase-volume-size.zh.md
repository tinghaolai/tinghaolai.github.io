---
title: "increase-volume-size.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["aws"]
---



# Translated by ChatGTP

## 增加容量大小

我們如何增加 EC2 實例的存儲大小？

* 在 EC2 控制台中修改卷大小
* 無需重新啟動實例
* 使用 `df -h` 檢查，仍然是舊的大小
* 使用 `lsblk` 檢查新的大小
  * 對於我的情況，名稱：`xvda`，根分區：`xvda1`
* `sudo growpart /dev/xvda 1`
  * > 改變了：分區=1 開始=262144 舊：大小=16515039 結束=16777183 新： 大小=104595423 結束=104857567
* `sudo resize2fs /dev/xvda1` 
* 使用 `df -h` 檢查新的大小成功。