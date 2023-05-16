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
