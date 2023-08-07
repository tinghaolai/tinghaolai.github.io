---
title: "docker.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["docker"]
---



# Translated by ChatGTP

## CMD與RUN與ENTRYPOINT的比較

### CMD

```dockerfile
CMD ["executable", "param1", "param2"]
CMD command param1 param2
```

* 只能有一個CMD，使用最後的
* 容器啟動時執行

### ENTRYPOINT

* 在把容器當作可執行文件時需要定義ENTRYPOINT
  * 而不是CMD
* 容器啟動時執行

### RUN

* 可以有多個RUN
* 在構建映像時執行
  * 會創建臨時容器，執行命令，然後提交到映像

## 使用bash訪問容器

在進入容器終端後，輸入`/bin/bash`以訪問bash。

如果希望更方便，可以使用容器腳本執行它。

> 例如，在Tabby終端中添加新的cmd配置，並在打開配置時設置`cmd.exe /c' docker exec -it php-74 bash'`命令（`php-74`是容器名稱或容器別名）


## 在本地呼叫容器URI

每個容器可以通過域`host.docker.internal`訪問容器的外部端口。

例如：

* API網址`http://host.docker.internal:8016/api/member/create`
* .env連接
    ```
    REDIS_HOST=host.docker.internal
    REDIS_PASSWORD=
    REDIS_PORT=6379
    REDIS_CACHE_DB=1
    ```

## 在另一個容器內呼叫容器

[解決方案](https://stackoverflow.com/questions/42385977/accessing-a-docker-container-from-another-container)

創建網絡並連接要連接的容器，並檢查它們的設置。

> 請注意，IP地址在Docker重新啟動後可能會更改。

### 打開現有容器

docker start broker

### 創建新容器（帶有網絡和靜態IP）

docker run --net developer-network --ip 172.25.0.9 -it confluentinc/cp-kafka:5.3.1 bash

## Docker橋接網絡固定IP

據我所知，無法為現有容器分配靜態IP，必須創建新容器。

[解決方案](https://stackoverflow.com/questions/27937185/assign-static-ip-to-docker-container)

## 更改密碼

這發生在我的副項目上，因為一開始正在測試，所以我使用的是默認密碼。

但是後來我認為這並不安全，

所以我在`docker-compose.yml`中修改它，

但是`docker-compose up -d`不起作用，

因為相關數據存儲在卷中，

所以我決定刪除該卷並重新創建它，

但後來我意識到我只需要修改容器中的帳戶設置即可。