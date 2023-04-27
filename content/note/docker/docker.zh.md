---
title: "docker.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["docker"]
---



# Translated by ChatGTP

## 透過 bash 存取容器 

進入容器終端機後，輸入 `/bin/bash` 以存取 bash 。

如果您想方便些，可以使用容器指令來執行。

> 例如 在 Tabby 終端機中增加新的指令配置，在設定打開時以 `cmd.exe /c 'docker exec -it php-74 bash'` 命令配置。( `php-74` 是容器名稱或容器別名 ) 


## 在本地環境中調用容器 URI

每個容器都可以通過其埠號與域名進行存取： `host.docker.internal`

例如：

* API 位址 `http://host.docker.internal:8016/api/member/create`
* .env 連接
    ```
    REDIS_HOST=host.docker.internal
    REDIS_PASSWORD=
    REDIS_PORT=6379
    REDIS_CACHE_DB=1
    ```

## 在另一個容器內部呼叫容器

[解決方法](https://stackoverflow.com/questions/42385977/accessing-a-docker-container-from-another-container)

建立網路，連接您想要連接的容器，並檢查其設置。

> 請注意，IP 可能會在 Docker 重新啟動後更改。

### 開啟現有的容器

docker start broker

### 建立新的容器(帶有網路和固定 IP)

docker run --net developer-network --ip 172.25.0.9 -it confluentinc/cp-kafka:5.3.1 bash

## Docker 橋接網路固定 IP

據我所知，無法為現有容器分配固定 IP，必須創建新的容器。

[解決方法](https://stackoverflow.com/questions/27937185/assign-static-ip-to-docker-container)

## 更改密碼

這發生在我的副業項目中，因為一開始是進行測試，所以我使用預設密碼。

但後來我覺得不安全，就在 `docker-compose.yml` 中修改了密碼。

但是， `docker-compose up -d` 沒有發揮作用，因為相關的數據是在卷中，

所以我決定刪除卷並重新創建它，但後來我意識到應該只需更改容器中的帳戶設置。