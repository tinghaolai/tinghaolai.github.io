---
title: "increase-volume-size.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["aws"]
---



---

## Increase volume size

How can we increase the storage size of an EC2 instance?

* Modify the volume size in the EC2 console
* Don't need to restart the instance
* check by `df -h`, still the old size
* `lsblk` to check the new size
  * For my case, name: `xvda`, root partition: `xvda1`
* `sudo growpart /dev/xvda 1`
  * > CHANGED: partition=1 start=262144 old: size=16515039 end=16777183 new: size=104595423 end=104857567
* `sudo resize2fs /dev/xvda1`
* `df -h` to check the new size success

---

