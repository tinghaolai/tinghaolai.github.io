---
title: "fair-queue-job-handle.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["backend-random-thing"]
---



---


### Queue jobs

Let's say one kind of job, 

that may very time-consuming, 

depend on that specific job,

so we can put it into a queue, 

and handle it later.

### Worker handle jobs

Worker will handle the jobs in the queue,

that's say we have three worker,

and one member calling 100 jobs that are time-consuming,

then other member calling jobs,

even the job is small, 

won't take much time,

but still need to wait for the first 100 jobs finished.

## Fair job handle

Split the jobs into small pieces,

and put them into the queue again.

So in this case,

the other member's job only need to wait for little slice job of the first 100 jobs.

We can improve the mechanism by many ways,

depends on your system,

but any this is the core concept.

## More thoughts

Other than this approach, 

which have only one machine,

there might be other better ways to handle this,

for example: horizontal scaling.

And what is the better way?

I think it depends on your system,

maybe it's totally not necessary to worry about in your system,

but can you handle it if it happens?



---

