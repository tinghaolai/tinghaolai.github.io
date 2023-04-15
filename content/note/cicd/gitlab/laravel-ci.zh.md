---
title: "laravel-ci.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["cicd","cicd-gitlab"]
---



# Translated by ChatGTP

## Laravel ci

範例專案：[Laravel-crawler](https://github.com/tinghaolai/laravel-react-crawler)


### docker-compose.yml

```yaml
version: '3.8'

services:
    app:
        build:
            context: .
            dockerfile: Dockerfile
        command: bash -c "php artisan serve --host=0.0.0.0 --port=8998 & npm run dev"
        ports:
            - 5174:5174
            - 8998:8998
        networks:
            - crawler_network
        volumes:
            - .:/application
        depends_on:
            - mysql

    mysql:
        image: mysql
        ports:
            - 3308:3306
        networks:
            - crawler_network
        environment:
            MYSQL_ROOT_PASSWORD: root
            MYSQL_USER: laravel
            MYSQL_PASSWORD: laravel
            MYSQL_DATABASE: laravel
        volumes:
            - mysql_data:/var/lib/mysql

networks:
    crawler_network:
volumes:
    mysql_data:
```

### Dockerfile

```yaml
FROM php:8.1.0

# Add docker php extension installer script
ADD https://github.com/mlocati/docker-php-extension-installer/releases/latest/download/install-php-extensions /usr/local/bin/

RUN chmod +x /usr/local/bin/install-php-extensions \
    && install-php-extensions @composer \
    calendar \
    exif \
    ffi \
    mysqli \
    pdo \
    pdo_mysql \
    gd

ENV NODE_VERSION=16.13.0
RUN apt install -y curl
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
ENV NVM_DIR=/root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"

WORKDIR /application
```


### .gitlab-ci.yml

```yaml
image: docker

services:
  - docker:dind

variables:
  CI_IMAGE: $CI_REGISTRY_IMAGE/dave-test-image:latest
  GITLAB_REGISTRY_USERNAME: gitlab+deploy-token-1611692
  MYSQL_DATABASE: laravel
  MYSQL_ROOT_PASSWORD: laravel

stages:
  - build
  - test

build_job:
  stage: build
  only:
    variables:
      - $CI_COMMIT_MESSAGE =~ /build-image/
      - $CI_COMMIT_TAG =~ /build-image/
  before_script:
    - docker login $CI_REGISTRY -u $GITLAB_REGISTRY_USERNAME -p $GITLAB_REGISTRY_PASSWORD
  script:
    - docker build -t $CI_IMAGE .
    - docker push $CI_IMAGE

test_job:
  stage: test
  image: $CI_IMAGE
  services:
    - mysql:8.0
  variables:
    MYSQL_ROOT_PASSWORD: laravel
    MYSQL_DATABASE: laravel
    MYSQL_USER: laravel
    MYSQL_PASSWORD: laravel
  script:
    - bash scripts/gitlab-ci.sh
    - composer install
    - php artisan key:generate
    - php artisan optimize
    - php artisan migrate:install
    - php artisan migrate
    - php artisan test
```

### gitlab-ci.sh

執行需要部署此專案所需的任何內容，例如這裡我需要為 Laravel 專案建立一個 `.env` 檔案。

```shell
cat <<EOT >> .env
APP_NAME=Laravel
APP_ENV=local
APP_KEY=
APP_DEBUG=true
APP_URL=http://localhost

LOG_CHANNEL=stack
LOG_DEPRECATIONS_CHANNEL=null
LOG_LEVEL=debug

DB_CONNECTION=mysql
DB_HOST=mysql
DB_PORT=3306
DB_DATABASE=laravel
DB_USERNAME=laravel
DB_PASSWORD=laravel

BROADCAST_DRIVER=log
CACHE_DRIVER=file
FILESYSTEM_DISK=local
QUEUE_CONNECTION=sync
SESSION_DRIVER=file
SESSION_LIFETIME=120

MEMCACHED_HOST=127.0.0.1

REDIS_HOST=127.0.0.1
REDIS_PASSWORD=null
REDIS_PORT=6379

MAIL_MAILER=smtp
MAIL_HOST=mailhog
MAIL_PORT=1025
MAIL_USERNAME=null
MAIL_PASSWORD=null
MAIL_ENCRYPTION=null
MAIL_FROM_ADDRESS="hello@example.com"
MAIL_FROM_NAME="${APP_NAME}"

AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_DEFAULT_REGION=us-east-1
AWS_BUCKET=
AWS_USE_PATH_STYLE_ENDPOINT=false

PUSHER_APP_ID=
PUSHER_APP_KEY=
PUSHER_APP_SECRET=
PUSHER_HOST=
PUSHER_PORT=443
PUSHER_SCHEME=https
PUSHER_APP_CLUSTER=mt1

VITE_PUSHER_APP_KEY="${PUSHER_APP_KEY}"
VITE_PUSHER_HOST="${PUSHER_HOST}"
VITE_PUSHER_PORT="${PUSHER_PORT}"
VITE_PUSHER_SCHEME="${PUSHER_SCHEME}"
VITE_PUSHER_APP_CLUSTER="${PUSHER_APP_CLUSTER}"

EOT

```


> 可以複製 Git 中的檔案（例如 `.env-gitlab`），並使用腳本來複製它，如果需要放置令牌和密碼等機密信息，應該放置於 Gitlab CI 變數中。


### Gitlab runner 設定

使用 Docker runner 執行 Git 專案，並使用 Docker in Docker 建置映像檔，安裝相關設置並運行測試，如果測試通過，則可以看到管線成功，如果測試未通過，則管線失敗。