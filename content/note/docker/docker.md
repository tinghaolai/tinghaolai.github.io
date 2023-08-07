---
title: "docker.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["docker"]
---



---

## CMD vs RUN cs ENTRYPOINT

### CMD

```dockerfile
CMD ["executable","param1","param2"]
CMD command param1 param2
```

* Can only have one CMD, use last one
* Execute when container start

### ENTRYPOINT

* ENTRYPOINT should be defined when using the container as an executable.
  * instead of cmd
* Executed when container start

### RUN

* Can have multiple RUN
* Execute when building image
  * Will create a temporary container, execute command, and commit to image

## Access container using bash

After enter container terminal, type `/bin/bash` to access bash.

If u want to make it convenient, execute it with container script.

> e.g. In Tabby terminal add new cmd config, and set `cmd.exe /c 'docker exec -it php-74 bash'` command when config open. (`php-74` is container name or container alias)


## Call container uri in local

Each container can be access with container's external port with domain: `host.docker.internal`

For example:

* Api url `http://host.docker.internal:8016/api/member/create`
* .env connection
    ```
    REDIS_HOST=host.docker.internal
    REDIS_PASSWORD=
    REDIS_PORT=6379
    REDIS_CACHE_DB=1
    ```

## Call container inside another container

[Solution](https://stackoverflow.com/questions/42385977/accessing-a-docker-container-from-another-container)

Create network and connect containers u want them connected, and inspect their setting.

> Note that ip could be changed after docker restarted.

### Open exists container

docker start broker

### Create new container (with network and static IP)

docker run --net developer-network --ip 172.25.0.9 -it confluentinc/cp-kafka:5.3.1 bash

## Docker bridge network fixed ip

As I know, can't assign static IP to exists container, have to create new one.

[Solution](https://stackoverflow.com/questions/27937185/assign-static-ip-to-docker-container)

## Change password

This happened to my side project, because it's testing at beginning,

so I use default password.

But then I think it's not safe, 

so I change it in `docker-compose.yml`,

but `docker-compose up -d` not work,

since the relate data is in volume,

so I decide to remove the volume and recreate it,

But then I realize that I should just change the account setting in container.


```bash





---

