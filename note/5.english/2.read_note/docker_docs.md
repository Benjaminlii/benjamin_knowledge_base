# Note Of Read Docker Docs

## Get started

### Part 1: Overview

> https://docs.docker.com/get-started/

什么是容器？
	是一个沙盒程序在本地机器上运行，和其他本地机器上运行的程序相互隔离。这种隔离通过linux内核的命名空间和cgroups实现。这是Linux很久前就实现的功能，docker旨在将这些能力变得更可接触和更可用。
什么是容器的镜像？
	当容器运行时，它会使用一个隔离的文件系统。这个用户自定义的文件系统就是由容器镜像提供的。
	由于这个镜像包含容器的文件系统，所以它必须包含容器运行应用所需要的所有依赖，配置项，脚本以及二进制文件等等。除此之外还需要包含容器运行所需的其他配置，比如环境变量，要执行的命令以及其他元数据

### Pare 2: Containerize an application

> https://docs.docker.com/get-started/02_our_app/

构建镜像并启动一个容器的例子：
获取一个应用：
	略：
构建镜像：
	Docker通过Dockerfile构建镜像，在当前目录下有Dockerfile文件，执行docker build命令可以进行镜像打包。
	Docker下载FROM指令后的镜像为基础镜像，执行RUN指令后的命令进行镜像的定制。CMD指令指定了从这个镜像启动的容器所执行的命令。
	`docker build -t <image_name> .` 命令的-t选项可以对镜像进行tag，可以起名字。最后的点指定了Dockerfile存放的路径。
启动容器：
	`docker run -dp 3000:3000 <image_name>` 该命令可以从一个镜像启动一个容器，-d选项为后台启动，-p指定宿主机到容器内的端口映射，后面需要跟镜像名。
	`docker ps` 该命令可以查看当前机器上所运行的容器，-a表示包含已停止的容器。

### Pare 3: Update the application

> https://docs.docker.com/get-started/03_updating_app/

如何更新工程?
可以在更新代码后,重新构建镜像,然后`docker run`启动.
但是需要先停止并删除旧的容器,因为会涉及端口抢占的问题.

```
docker ps
docker stop <container_id>
docker rm <container_id>
```

然后重新执行`docker run`启动即可

### Pare 4: Share the application

> https://docs.docker.com/get-started/04_sharing_app/

在docker hub上可以创建一个仓库,用来存放某个工程的镜像

```shell
docker login -u <username> # 登陆到docker hub
docker tag <image_name> <username/image_name> #给镜像重命名,username在推送时必须
docker push <username/image_name> # 推送镜像到远端仓库
```

执行`docker run <username/image_name>`即可在另一台机器上运行实例

### Pare 5: Persist the DB

> https://docs.docker.com/get-started/05_persisting_data/

在两个docker container之间，即使他们的启动镜像是同一个，他们的文件系统也都会是隔离的。这由docker的隔离机制决定。
但可以通过挂载docker卷宗来将容器内的存储映射到宿主机上，从而实现多个容器数据的共享或复用。

```shell
# 创建docker卷宗
docker volume create <volume_name>

# 启动容器时进行挂载
docker run --mount type=volume,src=<volume_name>,target=<path_in_container> <image_name>

# 查看卷宗属性 挂载点
docker volume inspect <volume_name>
eg:
[
    {
        "CreatedAt": "2019-09-26T02:18:36Z",
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/todo-db/_data",
        "Name": "todo-db",
        "Options": {},
        "Scope": "local"
    }
]

# 可以在docker run中使用 `-v` 来直接创建并使用挂载卷
docker run -v <volume_name>:<path_in_container> <image_name>
```

### Pare 6: Use bind mounts

> https://docs.docker.com/get-started/06_bind_mounts/

除了上节提到的挂载卷之外，还可以通过绑定挂载的方式将容器内外的存储联系起来。其好处是更加灵活的选择主机中path的位置。（与之相比volume可以启动容器时向新的卷写入数据，并且支持更换挂载驱动，以支持将数据挂载到云端）

```shell
docker run -it --mount type=bind,src=<path_in_host>,target=<path_in_container> <image_name> 
```

### Para 7: Multi container apps

> https://docs.docker.com/get-started/07_multi_container/

在多应用的场景下，将不同的应用放置在不同的容器中保证容器的单一职责会是比较好的选择。那么又会引入一个问题：不同容器之间如何通信？

Docker提供了容器网络来解决这一问题。

```shell
# 首先创建容器网络
docker network create <network_name>
# 启动容器
docker run -d --network <network_name> --network-alias <host_name_in_network> <image_name>
```

运行上面的命令后，连接到同一容器网络的容器之间就可以通过 `host_name_in_network` 这个host来访问该容器

### Para 8: Use Docker Compose

> https://docs.docker.com/get-started/08_using_compose/

Docker Compose 是一种定义多容器程序的工具，可以通过一个yml文件来定义一个程序中所需要启动的多个 Docker Container 。

大致格式如下：

```yml
# docker-compose.yml
version: '3'

services:
    app1:
        image: <image_name>
        command: sh -c "<your_cmd>"
        build: ./dir
        ports:
            - 3000:3000
        working_dir: /app
        volumes:
            - ./:/app
        environment:
            env_name: xxx
    
    app2:
        volumes:
        - <volume_name>:/xxx

volumes:
  <volume_name>:
```

1. version
指定本 yml 依从的 compose 哪个版本制定的。

1. image
当前容器的启动镜像

1. command
执行docker run后的命令
```yml
command: ["bundle", "exec", "thin", "-p", "3000"]
```
1. build
定义构建工程所需要的相关参数
    1. context：设置工程所在目录
    1. dockerfile：指定构建镜像的 Dockerfile 文件名。（在context：设置工程所在目录下）
    1. args：添加构建参数，这是只能在构建过程中访问的环境变量。

1. ports
指定端口开放及映射

1. working_dir
指定工作目录，相当于 docker run 的 -w

1. volumes
指定挂载卷

1. environment
向容器内添加环境变量

通过命令 `docker compose up` 即可在当前目录下寻找 `docker-compose.yml` 启动程序。
`docker compose down` 命令可以停止由 `docker-compose.yml` 启动的所有容器。

### Para 9: Image-building best practices

> https://docs.docker.com/get-started/09_image_best/

通过 `docker image history <image_id>` 命令可以查看镜像在构建过程中的每一个命令。

#### 层缓存
如下的一个dockerfile文件:

```dockerfile
# syntax=docker/dockerfile:1
FROM node:18-alpine
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["node", "src/index.js"]
```

在每次构建都会拷贝代码，并安装依赖。那么每次更新代码重新打包镜像时，过程中生成的每一层镜像都是不一样的。这种情况下docker镜像在构建过程中最为耗时的操作是安装镜像，但是大多数情况下依赖项并不会发生变化。

在这种情况下，将dockerfile做以下改动：

1. 在构建一开始并不拷贝整个仓库进镜像，而是只拷贝依赖定义文件 `package.json`。
1. 在 `yarn install --production` 安装好依赖后再拷贝仓库。
1. 在dockerfile文件错在目录下创建一个 `.dockerignore` 文件，在这个文件中定义在构建过程中拷贝文件时忽略 `node_modules` 文件，这样就可以忽略掉依赖文件的拷贝。

做完以上的改动，docker在构建镜像的过程中，会先去拷贝依赖定义文件，然后安装命令，然后再拷贝除依赖项的其他所有文件。

这样做的好处是在拷贝仓库之前（这一步得到的结果大概率每一次都不同），每一步运行的结果可以复用，就可以避免重复的进行依赖下载。

#### 多阶段构建

在dockerfile中可以有如下的定义

```dockerfile
FROM image_id_1 AS build
DO_SOMETHING

FROM image_id_2
COPY --from=build xxx xxx
DO_SOMETHING
```

上面的dockerfile在构建镜像的过程中会先根据镜像 `image_id_1` 去进行第一阶段的构建，并在后续的构建过程中可以使用 `COPY --from=build` 从第一阶段构建的结果中拷贝文件。除最后一阶段构建的结果之外，前面的构建的所有文件都不会被保存在最终的输出镜像中。

### Para 10: What next

> https://docs.docker.com/get-started/11_what_next/

集群的编排：
在生产环境中，在物理机器上直接操控容器是很困难的，无论是使用 `docker run` 或者 `docker-compose`。想要跨多台机器中对多个容器进行编排，kubernetes等等提供了方案。
一般的方案是会有一个接受预期状态的管理器，例如“运行web应用的两个实例，公开80端口”，然后管理器就会查看集群中的的所有机器并做出一些动作，然后返回执行后的集群状态。

云原生计算基金会（CNCF）：
CNCF是一个中立的开源项目，包括Kubernetes等项目，有很多项目可以解决监控、日志、安全等等问题，

从头创建容器：mark
Aqua Security的Liz Rice有一个演讲，演示使用go从头开始创建容器，深入探讨了工作原理。
