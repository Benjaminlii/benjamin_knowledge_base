# Note Of Read Kafka Docs

## 1. GETTING STARTED

### 1.1 Introduction

> https://kafka.apache.org/documentation/#introduction

#### 什么是事件流?

事件流相当于数字版的人体中枢神经系统.

在未来, 业务将越来越多的由软件定义,且自动化,并且软件的用户也是软件.事件流是实现这一目标的技术基础.

从技术上讲,事件流是实时的从事件源中捕获数据的一种实践, 存储这些数据供后续的检索, 实时地回顾性地对事件流进行操控、处理以及响应. 根据需要讲事件流路由到不同的目标技术中. 因此事件流保证了数据的连续流动和解析,从而让正确的信息出现在正确的时间和正确的地点.

#### 我们可以使用事件流来做什么?

事件流被广泛的应用与大量的行业和组织的各种各样的用例中. 举例:

1. 实时的处理支付和金融事务, 如股票交易、银行和保险业.
1. 实时的跟踪和监控汽车、车队以及货运, 如物流和汽车行业.
1. 持续的捕获和分析来自物联网设备或者其他设备中的传感器数据.
1. 收集并立即响应用户的活动以及指令.
1. 连接、存储和提供由公司不同部门产生的数据.
1. 是数据平台,事件驱动架构以及微服务的基础.

#### kafka是一个事件流平台, 意味着什么?

kafka提供了三种能力来实现端到端的事件流用例:

1. 发布和订阅事件流
1. 持久并且可靠的存储事件流
1. 事件发生时或者回溯性的处理事件流

所有这些功能都以分布式,可拓展,伸缩,容错,安全的方式提供. kafka可以部署在物理机,虚拟机,容器,本机以及云中.你可以自己管理kafka环境也可以使用提供商提供的完整托管服务.