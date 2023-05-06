---
title: "start-container-after-reboot.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["docker"]
---



---

## Start container after reboot

* `vim etc/systemd/system/sentry.service`
  * ```yaml
    [Unit]
    Description=Start sentry_line_notify_webhook_app container
    Requires=docker.service
    After=docker.service
    
    [Service]
    Restart=always
    ExecStart=/usr/bin/docker start -a sentry_line_notify_webhook_app_1
    
    [Install]
    WantedBy=multi-user.target
    ``` 
* `sudo systemctl daemon-reload`
* `sudo systemctl enable sentry`
* `sudo systemctl start sentry`

Then you can check by `docker ps` after reboot.

---

