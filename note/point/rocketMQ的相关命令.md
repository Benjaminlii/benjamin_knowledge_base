# 启动服务
```shell
nohup sh bin/mqnamesrv &
nohup sh bin/mqbroker -c ./conf/broker.conf -n localhost:9876 autoCreateTopicEnable=true &
cat nohup.out
```
# 关闭服务
```shell
sh bin/mqshutdown namesrv
sh bin/mqshutdown broker
eval "" > nohup.out
```
