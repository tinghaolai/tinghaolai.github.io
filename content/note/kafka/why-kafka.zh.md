---
title: "why-kafka.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["kafka"]
---



# Translated by ChatGTP

## 為什麼使用Kafka

### 在討論訊息佇列之前

* 輪詢
  * 浪費 - 每次都詢問是否有新任務（即使沒有新任務）

### 為什麼要使用訊息佇列

Kafka是一種訊息佇列，所以我們首先討論為何使用訊息佇列。

* 永久性 / 可靠性 / 交易：防止資料遺失
* 效能 / 負載平衡

### 為什麼將Kafka與其他訊息佇列進行比較

* 效能
  * TPS（每秒交易數）很高
    * Kafka：百萬
    * Rocket MQ：10萬
    * Active MQ：1萬
  * 零拷貝