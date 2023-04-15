---
title: "ignore.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["git"]
---



# Translated by ChatGTP

## 忽略文件

如果你在進行一個大型專案的開發，你可能會有許多不想提交的文件或更改，例如測試代碼或其他使開發更輕鬆的東西。當你想進行提交時，你必須挑選你想要提交的文件，而不是開發所做的更改。

這可能需要很長時間，特別是如果你沒有像JetBrain的提交UI這樣強大的IDE。這裡有一種我自己的方法：手動忽略。

### 手動忽略

1. 找出你不想提交的文件。
2. 把重置命令保存在一個地方，以便在需要時複製它
    * `git reset App\TestA.php App\TestB.php public\index.js`
3. 當你需要提交所有當前更改時，運行`git add .`
4. 運行複製的重置命令。