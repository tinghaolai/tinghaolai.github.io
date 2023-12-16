---
title: "arch-linux-install-cloudflare-zero-trust.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["linux"]
---



---


## arch linux install cloudflare zero trust

After install `warp-cli`, if you get `Cloudflare warp daemon error when register the client`, run `systemctl enable --now warp-svc.service`.

### Client register team

* run `warp-cli teams-enroll {teamName}`
* open url with firefox and copy token from console, which token start with "com.cloudflare.warp://".
* `warp-cli status`
* `warp-cli start`
* `warp-cli status`

### Client register team v2 (Linux without GUI)

* run `warp-cli teams-enroll {teamName}`
* open url generate from last command in browser (any computer with GUI).
* open `network`, after login, copy `Request URL`
* `warp-cli teams-enroll-token {Request URL}`
* Check result by
  * `warp-cli start`
  * `warp-cli connect`
  * `warp-cli disconnect`
  * `warp-cli account`

---

