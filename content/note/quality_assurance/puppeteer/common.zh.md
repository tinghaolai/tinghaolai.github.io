---
title: "common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["quality_assurance","quality_assurance-puppeteer"]
---



# Translated by ChatGTP

## 核心-全尺寸

使用 `--start-maximized` 參數

```javascript
    const browser = await puppeteer.launch({
        executablePath:
        chromePath,
        headless: false,
        defaultViewport: {
            width: 1618,
            height: 787
        },
        args: [
            '--start-maximized'
        ]
    });
```