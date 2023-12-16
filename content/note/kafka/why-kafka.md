---
title: "why-kafka.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["kafka"]
---



---

## Why Kafka

### General purpose

Handle stream data.

### Before message queue

* Polling
  * waste - ask if new mission every time (even if no new mission)

### Why message queue

Kafka is kind of message queue, so that's first discuss why using mq.

* persistence / reliability / transactional: prevent data loss
* performance / load balance

### Why kafka compare to other mq (Result)

* Performance
  * TPS (Transaction per Second) is high
    *  kafka: million
    *  Rocket MQ: 100k
    *  Active MQ: 10k
  * Zero copy (will discuss in another note `why-kafka-is-fast.md`)

### Why kafka compare to other mq (Reason)

We can process data using queue,
but the problem is latency.
Kafka is faster than other mq when data is large.

From note: `how-kafka-works.md` and `why-kafka-is-fast.md`,
we can know kafka read data from log file by fast Polling.

But why queue can't do fast Polling in other mq?
Because other mq don't have optimization like kafka,
such as frequent request, lock mechanism, etc.
This could cause many performance and abnormal problems.

### Key difference between kafka and other mq

People thing message system should support complex message search function,
but kafka don't implement this function,
so kafka can
* Have faster performance
* Keep performance when message backlog is large

But in other hand, 
kafka sacrifice some function like message priority.

---

