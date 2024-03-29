# Note Of Read Kafka Docs

## 1. 开始

### 1.1 引言

> https://kafka.apache.org/documentation/#introduction

#### 1.1.1 什么是事件流?

事件流相当于数字版的人体中枢神经系统.

在未来, 业务将越来越多的由软件定义, 且自动化, 并且软件的用户也是软件.事件流是实现这一目标的技术基础.

从技术上讲, 事件流是实时的从事件源中捕获数据的一种实践, 存储这些数据供后续的检索, 实时地回顾性地对事件流进行操控、处理以及响应. 根据需要讲事件流路由到不同的目标技术中. 因此事件流保证了数据的连续流动和解析, 从而让正确的信息出现在正确的时间和正确的地点.

#### 1.1.2 我们可以使用事件流来做什么?

事件流被广泛的应用与大量的行业和组织的各种各样的用例中. 举例:

1. 实时的处理支付和金融事务, 如股票交易、银行和保险业.
1. 实时的跟踪和监控汽车、车队以及货运, 如物流和汽车行业.
1. 持续的捕获和分析来自物联网设备或者其他设备中的传感器数据.
1. 收集并立即响应用户的活动以及指令.
1. 连接、存储和提供由公司不同部门产生的数据.
1. 是数据平台, 事件驱动架构以及微服务的基础.

#### 1.1.3 Kafka是一个事件流平台, 意味着什么?

Kafka提供了三种能力来实现端到端的事件流用例:

1. 发布和订阅事件流
1. 持久并且可靠的存储事件流
1. 事件发生时或者回溯性的处理事件流

所有这些功能都以分布式, 可拓展, 伸缩, 容错, 安全的方式提供. Kafka可以部署在物理机, 虚拟机, 容器, 本机以及云中.你可以自己管理Kafka环境也可以使用提供商提供的完整托管服务.

#### 1.1.4 Kafka是如何工作的

Kafka是服务端和客户端通过高性能TCP协议通信的一个分布式系统。可以部署在物理机, 虚拟机, 容器中。

服务端：Kafka作为集群运行, 可以扩越多个区域。其中一些服务构成存储层, 成为代理。其他的服务器运行Kafka connect, 通过事件流的形式导入导出数据, 从而将Kafka和现有的系统（如关系型数据库）集成。Kafka有高度可扩展行和容错性, 如何任何一个服务器出现错误, 其他的服务器会代替它工作。

客户端：允许你便携分布式的app和微服务, 在并行, 大规模和容错的方式读写和处理事件流。Kafka附带了一些客户端, 在java, scala的客户端中可以支持更高级的 Kafka Streams 库。还提供了go, python, c/c++等等编程语言的客户端以及REST API。

#### 1.1.5 主要概念和术语

一个事件意味着你的业务中“有什么事情发生了”。这在文档中也被称为记录或者消息。当你在Kafka中读写数据时，你就用事件的形式做了这件事。概念上，一个事件包含键，值，时间戳和可选的元数据头。举例如下：

- 键：Alice
- 值：付200给Bob
- 时间戳：2022.1.25 14:06

生产者是发布时间到Kafka的这些客户端应用，消费者是订阅（读并且处理）这些事件的端。在Kafka中，生产者和消费者是完全解耦并且相互不感知的。这是Kafka实现总所周知的改读可拓展性的关键设计元素。举个例子，生产者从不需要等待消费者。Kafka提供了各种各样的保证，比如只处理一次事件的能力。

事件是有组织的并且持久的存储在主题（topic）中。非常简单，一个主题就是一个文件系统中的文件夹，事件则是这个文件夹中的文件。主题在Kafka中总是多生产者和多订阅的：一个主题可以有零至多个生产者想起中写入事件，也可以有零至多个消费者订阅这些事件。主题中的事件可以根据需要随时读取，这和传统消息系统不同，事件在在被消费后不会被删除。相反，你可以定义Kafka按照主题纬度定义保留你的事件多久，在此之后，旧的事件将会被丢弃。Kafka的性能在数据大小方面是有效恒定的，所以长时间的存储数据时完全没问题的。

主题是分段的，这意味着主题分布在不同Kafka代理的许多“桶”上。这种分布式的数据放置对可拓展性是非常重要的，因为这允许客户端应用同时在多个代理中读写数据。当新事件被发布到主题中时，实际上会被附加到这个主题的某个分区中。具有相同事件key的事件会被写入到同一个分区中，Kafka保证指定主题分区的多个消费者总是在用事件写入的顺序来读取该分区的事件。

![主题分区](https://kafka.apache.org/images/streams-and-tables-p1_p4.png)
上图：示例的主题有四个分区P1-P4，两个不同的生产者相互独立地通过网络向主题的分区中写入事件来发布新事件。这些事件拥有同样的key（图中按颜色表示）的会被写入到同一个分区中。两个生产者的事件也可以被写入同一个分区，只要key相同。

为了数据的容错性和高可用性，每一个主题都可以被复制，甚至跨越地理区域或者数据中心，所以总会有多个代理拥有拷贝的数据，以防止出现问题，你可以去维护代理等等。一个公用的生产设置是复制因子为3，也就是说，一份数据总会有三个副本。这种复制是在主题分区级别执行的。

#### 1.1.6 Kafka APIs

除了命令行工具之外，Kafka还为Java和Scala提供了5个核心API：
- 管理api用于管理和检查主题，代理和其他的Kafka实例。
- 生产者api用于向一个或者多个Kafka主题发布时间流消息。
- 消费者api用于订阅一至多个主题的消息并进行处理。
- Kafka流api实现流处理应用和微服务。它提供了高级方法来处理事件流，包括转换，操作状态（聚合和连接），窗口，基于事件时间的处理等等。从一至多个主题中读取输入，为一至多个主题生成输出。有效的实现输入输出的转换。
- Kafka连接api可以构建和运行可复用的导入导出连接器，这些连接器可以让外部系统或应用向生产者或者消费者中读写数据，从而让他们与Kafka集成。例如：一个连接到关系型数据库的连接器可以捕获对表的每一次更改。实际上并不需要实现自己的连接器，Kafka社区已经提供了非常多现成的连接器。

### 1.2 用例

以下是关于Kafka的一些流行用例：

#### 1.2.1 消息

Kafka可以非常好的替代传统的消息代理。使用消息代理的原因有很多（从数据生产者那里解耦消息的处理，缓存未处理的消息，等等）。相对于大多数的消息系统，Kafka拥有更好的吞吐量，内置的分区，复制和容错，这使得Kafka成为大规模数据处理应用的一个好的解决方案。
在我们的经验里，消息的使用通常是低吞吐量的，但可能会需要低的端到端延迟，并且通常依赖Kafka提供的强大的持久化保证。

在这个领域内kafka可以于传统的消息系统相媲美。

#### 1.2.2 网站活动跟踪

Kafka的原始用例是能够重构用户活动跟踪流程为一组实时的发布订阅数据源。这意味着网站活动（页面浏览，搜索或者其他用户可能的活动）会被发布到核心主题，每一个活动类型只有一个主题。这些数据源可以用于订阅一些用例，包括实时处理，实时监控，以及加载到Hadoop或者离线的数据仓库系统以供离线处理和报告。
活动跟踪的量通常非常大，每个用户页面都会生成很多活动消息。

#### 1.2.3 指标

Kafka常常被用来操作监控数据。这涉及到聚合来自分布式应用的统计信息，并生成集中式的操作数据源。

#### 1.2.4 日志聚合

很多人将Kafka作为日志聚合解决方案的替代品。日志聚合通常需要在服务器外收集原始日志文件，并且将它们上传到一个统一的位置（一个文件服务或者Hadoop分布式文件系统）去处理。Kafka抽离了文件细节，并且以更清晰的方式将日志或者事件数据抽象为一系列消息。这个允许更低延迟的处理、更容易支持的多数据源和分布式数据消费。相对比Scribe或者Flume这样的以日志为中心的系统，Kafka提供了同样好的性能，由于复制带来的更强大的持久化保证，更低的端到端延迟。

#### 1.2.5 流式处理

许多 Kafka 的用户使用由多个节点组成的处理流程来处理数据，从 Kafka 主题中消费原始数据，然后进行聚合，丰富，或者以其他方式转化为新的主题进一步被后续的流程处理。例如：一个推荐新文章的处理流程可能会从RSS源来爬取文章内容，然后将其发布到一个‘文章’的主题；进一步的处理可能是对这些内容规范化或者去重，然后发布这个净化后的文章内容到一个新的主题；最终的处理步骤可能是尝试将该内容推荐给用户。这样的处理流程创建了一个基于每一个主题的实时数据流向图。从0.10.0.0版本开始，Kafka 提供了一个名为 Kafka Streams 的轻量级但是强大的流处理库去处理类似上述的数据处理流程。除了 Kafka Streams, 其他的开源流处理工具还包括 Apache Storm 和 Apache Samza

#### 1.2.6 事件溯源

事件溯源是一种应用程序设计风格，状态的更新会被记录为时间顺序排列的记录。Kafka对非常大存储日志数据的支持使得它成为了以这种风格构建应用程序的杰出后端。

#### 1.2.7 提交日志

Kafka可以为分布式系统提供一种外部提交日志的服务。日志可以帮助复制节点之间的数据，然后用于失败节点重新同步数据的机制来帮助其恢复数据。Kafka日志压缩的特性可以帮助支持这种用法。在这个用法中，Kafka类似于Apache BookKeeper项目。

## 4. 设计

#### 4.1 动机

我们设计Kafka是为了能够以一个统一的平台去处理一个大公司可能拥有的全部数据源。为了做到这一点，我们需要思考非常广的用例。

它需要有高吞吐量去支持大容量的时间流，例如实时的日志聚合。

它需要优雅的处理巨量的数据积压，从而支持来自线下系统的周期性数据加载。

同时意味着系统需要处理低延迟的交付去处理更多传统的消息功能。

我们想要支持分区，分布式，实时处理这些数据源，以生成新的衍生数据。这促进了我们的分区和消费者模型。

最后在数据流向其他数据系统用来进行服务的情况下，我们知道系统可能需要能够在机器存在故障的情况下保证容错性。

支持这些用例让我们采用了一种具有很多独特元素的设计，对比传统消息系统更像一个数据库日志。

#### 4.2 持久性

不要害怕文件系统。

Kafka非常重依赖文件系统来存储和缓存消息。有一种普遍看法认为“磁盘是慢的”，这让人们怀疑持久化结构能否提供有竞争力的性能。事实上磁盘比人们的期望慢的多的同时又快得多，这取决于人们怎么使用它，并且一个合适的磁盘结构设计通常可以像网络一样快。

关于磁盘性能的重要事实是在最近的十年中磁盘驱动的吞吐量和磁盘查找的延迟一直在偏差。在同一块磁盘上顺序写入和随机写入的性能差距可能高达6000倍。这些线性的读写在所有的用例中几乎是完全可以预见的，并且很大程度上可以由操作系统进行优化。一个现代操作系统提供了预读和延迟写入技术，可以预读取多个大块的数据，并将小的逻辑写入组合成大的物理写入。关于这个问题的更多的讨论可以在ACM Queue的文章中找到。实际上他们发现顺序的磁盘访问有时会比随机的内存访问更快。

为了补偿这些性能偏差，现代操作系统开始越来越积极的使用内存来进行磁盘缓存。现代操作系统会很乐意使用所有空闲内存进行磁盘缓存，内存回收时的性能开销很小。所有磁盘读写都将通过这个统一缓存。这个性能不容易被关闭，除非直接使用IO。因此即使一个进程维护了进程的缓存数据，这些数据也可能重复存在于操作系统的页缓存，实际上存储了两次。

此外，我们是在JVM之上构建的，任何或时间研究Java内存使用的人都会知道：
  - 内存中对象的开销非常高，通常是存储数据大小的两倍（或者更糟）。
  - 随着对内存中数据的增加，Java垃圾收集器变得越来越繁琐和慢。

由于这些因素，使用文件系统并且信任页缓存要优先于维护一个内存缓存或者其他结构--通过自动访问所有空闲内存，我们最少可以将可用的缓存翻倍，甚至有可能通过存储压缩后的字节而不是独立的结构来再次翻倍。这样做将会在一个32GB内存的机器上产生搞到28-30GB的缓存，并且没有GC惩罚。此外，即使服务重启这个缓存将会保持热状态，反而进程缓存需要在内存中进行重建（10GB的缓存可能会需要10分钟）或者需要通过一个冷的缓存启动服务（这意味着糟糕的初始性能）。这也会非常简化代码，因为所有的维护缓存和文件系统一致性的逻辑都在OS中，相比进程内的一次性活动，操作系统能够更有效且更正确地完成。如果你的磁盘偏向线性读取，那么预读会在每一次磁盘读取之前有效的预先填充这个缓存，与有用的数据。

这暗示了一个非常简单的设计，相比尽可能才内存中保留数据，并在空间不足时将所有数据刷新带文件系统中，我们翻转一下。所有的数据直接写入文件系统上的持久化日志，不一定要刷新到磁盘。实际上意味着数据被传输到内核的页缓存中。

这种页缓存为中心的设计风格在关于Varnish的设计文章中有所描述（并附带了一些自信）。

#### 4.3 性能

我们在提高效率方面下了很大功夫。我们的主要用例之一是处理网页活动数据，它的量非常大。每一个页面都会生成数十个写操作。此外，我们认为每一个被发布的消息都会被至少一个（通常会有多个）消费者读取，因此我们努力使消费更加轻松。

我们还从构建和运行类似的系统发现，性能是有效的多用户操作的关键。如何下游的基础架构系统容易因为程序使用量的小幅增加而成为瓶颈，这种小的改变很容易造成问题。通过非常好的性能，我们帮助确保在负载下程序将在基础架构之前故障。这在尝试在核心集群上运行一个支撑几十乃至上百应用的核心服务时非常重要，因为用户使用方式的变化几乎每天都会发生。

我们在之前的小节讨论了磁盘性能，一旦消除了糟糕的磁盘访问模式，在这种系统中，效率低下有两种常见的原因：太多小型IO操作，以及过多的字节复制。

小IO问题出现即发生在客户端和服务端之间，也发生在服务端本身的持久化操作中。

为了避免这个问题，我们的协议围绕“消息集合”的概念来进行构建，自然的将消息组合在一起。这允许网络请求将消息组合在一起，分摊网络往返的开销，而不是每次发送一条消息。服务端反过来每次加载一组消息到它的日志中，消费者一次获取一个大的线性快。

这个简单的优化产生了数量级的性能增长。批处理大的网络包，大的连续磁盘操作，相邻内存块等等，这些都允许Kafka将突发性的随机消息写入转换为流向消费者的线性写入。

另一个低效率的是字节复制。在低消息速率下这不是个问题，但是在负荷下影响是显著的。为了避免这个问题我们采用了一个生产者代理和消费者都共享的标准二进制消息格式（因此数据块可以不经过修改就在他们之间传输）。

通过代理维护的消息日志本身只是一个文件目录，每一个都通过一系列消息集填充，这些消息集按照与生产者和消费者一样的格式被写入磁盘。维护这个通用格式可以优化最重要的操作：持久化日志块的网络传输。现代unix操作系统为将数据从页缓存传输到socket提供了高优化代码路径；Linux中通过sendfile的系统调用完成。

为了理解sendfile的影响，理解从文件到socket传输数据的通用数据路径是很重要的：
1. 系统从磁盘读取数据到内核空间的页缓存中
2. 应用从内核空间读取数据到用户空间缓冲区
3. 应用回写数据到内核空间的socket缓冲区
4. 操作系统从socket缓冲区到通过网络发送的NIC缓冲区复制数据

这很清楚是低效率的，这里有四次拷贝和两次系统调用。使用sendfile，通过允许系统从页缓存直接发送数据到网络来避免这些重复的拷贝。在这个优化路径下，只需要最后拷贝到NIC缓冲区。

我们期望一个常见用例是多个消费者在一个主题上。使用上面的零拷贝优化，数据一次性被精确地拷贝到页缓存，并且每次消费时服用而不是每次都从用户空间中读取并拷贝出来并且存储到内存中。这允许消息按照接近网络限制的速率被消费。

页缓存和sendfile的结合意味着，在消费者大多数都被赶上的Kafka集群中，你见看不到任何磁盘上的读活动，因为它们将完全从缓存中提供数据。

TLS/SSL 库操作用户空间（内核SSL_sendfile目前Kafka还不支持）。由于这个限制，SSL启用时sendfile不会被使用。

#### 4.4 生产者

##### 4.4.1 负载均衡
生产者直接发送数据到分区领导者的代理不经过任何中间路由层。为了帮助生产者做到这一点，所有Kafka节点都可以在任意给定的时间里回答一个关于哪一个服务可用，一个主题的分区的领导者在哪里的元数据请求，以允许生产者适当地管理其请求。

客户端控制发布消息到那一个分区。这可以随机完成，实现一种随机的负载均衡，或者它可以有一些语义分区函数来完成。我们暴露了语义分区的接口，通过允许用户指定一个key来分区或者用这个key的hash结果来分区。如果需要，也可以选择覆盖分区函数。例如，如果这个key被挑选为user id，那么给定user的所有数据竟会被发送到同样的分区。反过来，这将允许消费者对他们的消费做出本地臆断。这种分区风格明确的被设计用来允许在消费者中进行对位置敏感的处理。

##### 4.4.2 异步发送
批处理是提高效率的重要因素之一，为了支持批处理，Kafka生产者尝试在内存中堆积数据，并且在一个单独的请求中发送数个大的批次。批处理可以被配置堆积不超过固定数量的消息，等待不超过一些固定值的延迟范围。这允许堆积更多的字节再发送，以及几乎没有大的系统IO操作。这个缓冲是可配置的，并且提供了一种机制来权衡少量的额外延迟以获得更好的吞吐量。

#### 4.5 消费者
Kafka消费者通过向想要消费的分区的领导代理发布fetch请求来工作。消费者每一次请求都会在日志中指定它的偏移量，并且从这个位置接收一大块日志。因此消费者具对这个位置有重要的控制权，并且可以按需倒回到这个位置来再消费数据。

##### 推送 vs 拉取
一个我们深思熟虑的最初问题是消费者应该从代理拉取数据，还是代理应该推送数据到消费者。在这方面，Kafka遵循了更传统的设计，像其他大多数消息系统一样，数据从生产者推送到代理，由消费者从代理处拉取。一些以日志为中心的系统，像Scribe和Apache Flume，遵循一个非常不同的基于推送的路径，其数据被向下推送，这两种方法各有利弊。然而，一个基于推送的系统会很难处理不同的消费者，因为代理控制了数据传输的速率。消费者的目标通常是能够按照尽可能大的速率进行消费；不幸的是，对于推送系统而言，当消费者的消费速度低于生产速度时（本质上是拒绝服务攻击），消费者往往会不堪重负。一个基于拉取的系统具有更好的属性，消费者只是落后，并且在可能的时候会赶上。这种情况可以通过某种退让协议来减轻，通过这种协议，消费者可以表明自己已经不堪重负，但是要充分的利用传输速率（但不要过度）比看起来更加棘手。以前以这种方式构建系统的尝试导致我们采用了更传统的拉取模型。
