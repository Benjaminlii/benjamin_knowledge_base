# local 和 remote 传输文件

在很多跑脚本，部署服务等等很多时候，都需要本地和远端服务上传或者下载文件，这时可以使用scp命令。

```shell
# scp [选项] [源文件] [目标位置]

# 本地上传文件到远端
scp ./* tao_li@xxx.xxx.xxx:/home/tao_li/xxx

# 本地上传文件到远端
scp tao_li@xxx.xxx.xxx:/home/tao_li/xxx ./
```
