---
title: "fair-queue-job-handle.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["backend-random-thing"]
---



# Translated by ChatGTP

### 佇列作業

假設有一種作業，

可能非常耗時，

需要依靠特定作業，

因此我們可以將其放入佇列中，

稍後再處理。

### 工作者處理作業

工作者將處理佇列中的作業，

假設我們有三個工作者，

有一名成員呼叫100個耗時的作業，

然後其他成員呼叫作業，

即使該作業很小，

不需要花太多時間，

仍然需要等待前100個作業完成。

## 公平處理作業

將作業分為小塊，

再次放入佇列中。

所以在這種情況下，

其他成員的工作只需等待前100個作業的一小塊作業即可。

我們可以通過多種方式改善機制，

取決於您的系統，

但這是核心概念。

## 更多想法

除此方法外，

假如只有一台機器，

可能有其他更好的方法來處理這個問題，

例如：水平擴展。

什麼是更好的方法？

我認為這取決於您的系統，

也許在您的系統中完全不需要擔心，

但如果發生了，您能處理嗎？