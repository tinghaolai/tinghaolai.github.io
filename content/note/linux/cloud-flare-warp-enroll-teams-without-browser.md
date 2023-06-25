---
title: "cloud-flare-warp-enroll-teams-without-browser.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["linux"]
---



---

warp-cli teams-enroll <TeamName>

Will make you open browser `https://<Team>.cloudflareaccess.com/warp`

but what if your server is not a desktop?

## Solution

Open the url within a browser on your desktop which is logged in to the same account,

and copy the token from console.

back to your server, 

run `warp-cli teams-enroll-token com.cloudflare.warp://<Team>.cloudflareaccess.com/auth?token=XXXXXXXXXXXXXXXXX`

---

