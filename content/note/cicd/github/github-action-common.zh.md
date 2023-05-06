---
title: "github-action-common.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["cicd","cicd-github"]
---



# Translated by ChatGTP

### 輸入設備不是 TTY

範例：
```bash
docker-compose exec app /app/migrate migrate
```

解決方案：在指令中加入 `-T`
```bash
docker-compose exec -T app /app/migrate migrate
```

### 將註冊圖片公開化

非常感謝 Github 的 UI/UX 團隊，

他們讓我撰寫此筆記，

只是為了修改可見度。

1. 前往組織頁面
    * 不是從存儲庫頁面
    * 注意來自套件頁面
2. 前往設置頁面
    * 如果沒有 `packages` 標籤，那麼你就在錯誤的地方
3. 然後修改可見度

或者只需像這樣替換 URL

https://github.com/organizations/streaming-data-architecture/settings/packages