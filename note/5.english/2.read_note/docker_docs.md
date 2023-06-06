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

```dockerfile
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
