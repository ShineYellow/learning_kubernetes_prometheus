
第四讲 Prometheus 运行框架介绍




第四讲内容




• 简单了解⼀下 promethues的框架结构




• 介绍⼀下 prometheus 的各种组件




1） 框架结构的展⽰图




image




• 我们先来看下这个部分




image




这⾥是 prometheus的服务端 也就是核⼼

prometheus本⾝是⼀个以进程⽅式启动，之后以多进程和多线程实现监控数据收集 计算 查

询 更新 存储 的这样⼀个C/S模型运⾏模式 本⾝的启动很简单




image

如果不带参数的 不考虑后台运⾏的问题 那么 prometheus 默认启动更简单




解压缩之后

./prometheus 即可 之后默认监听在9090端⼜ ⽤来访问




• 接下来我们来看看prometheus的存储形式




image




这是⼀段来⾃官⽅的 prometheus存储数据的介绍




Prometheus includes a local on-disk time series database, but also optionally integrates with remote storage systems.

Local storage

Prometheus's local time series database stores time series data in a custom format on disk.

Ingested samples are grouped into blocks of two hours. Each two-hour block consists of a directory containing one or more chunk files that contain all time series samples for that window of time, as well as a metadata file and index file (which indexes metric names and labels to series in the chunk files). The block for currently incoming samples is kept in memory and not fully persisted yet. It is secured against crashes by a write-ahead-log (WAL) that can be replayed when the Prometheus server restarts after a crash. When series are deleted via the API, deletion records are stored in separate tombstone files (instead of deleting the data immediately from the chunk files).




⼤⽶来给⼤家翻译和简化⼀下定义 ， 以上介绍主要包括如下⼏个点




• prometheus 采⽤的是 time-series （时间序列）的⽅式 以⼀种⾃定义的格式存储在 本 地硬盘上

• prometheus的本地T-S（time-series）数据库以每两⼩时为间隔来分block（块）存储， 每⼀个块中 又分为



多个chunk⽂件（我们以后会介绍chunk的概念），chunk⽂件是⽤来存放 采集过来的数据的
T-S 数据，metadata 和 索引⽂件（index）
• index⽂件 是对metrics(prometheus中 ⼀次K/V 采集数据 叫做⼀个metric) 和 labels(标 签) 进⾏索引 之后存储在 chunk中

chunk 是作为存储的基本单位，index and metadata是作为⼦集

• prometheus平时是将采集过来的数据 先都存放在内存之中（prometheus对内存的消耗 还是不⼩的）以类似缓存的⽅式 ⽤于加快搜索和访问

• 当出现当机时，prometheus有⼀种保护机制 叫做WAL 可以讲数据定期存⼊硬盘中 以

chunk来表⽰，并在重新启动时 ⽤以恢复进⼊内存




— 然后我们来看下⼤图中的 这个部分




image

这⾥⾯讲的是 prometheus 可以集成的 服务发现功能 例如 Consul

prometheus本⾝跟其他的开源软件类似 也是通过定义配置⽂件 来给prometheus本⾝规定需要 被监控的项⽬和被监控节点




我们看下配置⽂件的模版




image




配置⽂件中 规定了⼀个 ⼤的Job的名字

之后 在这个Jobs的名字下⾯ 具体来定义 要被监控的节点 以及节点上具体的端⼜信息等等




那么如果prometheus 配合了 例如consul这种服务发现软件

prometheus的配置⽂件 就不再需要⼈⼯去 ⼿⼯定义出来，⽽是能⾃动发现集群中 有哪些新 机器 以及新机器上出现了哪些新服务 可以被监控




— 接下来我们来看看 采集客户端的部分




image




prometheus 的客户端 之前给⼤家介绍过 主要有两种 ⽅式采集

pull 主动拉取的形式

push 被动推送的形式




pull: 指的是 客户端（被监控机器）先安装各类已有exporters(由社区组织 或企业 开发的监控 客户端插件)在系统上

之后，exporters 以守护进程的模式 运⾏ 并 开始采集数据

exporter 本⾝也是⼀个http_server 可以对http请求作出响应 返回数据（K/V metrics）




prometheus ⽤ pull 这种主动拉的⽅式（HTTP get） 去访问每个节点上exporter 并采样回需要 的数据




push ：




image

指的是 在客户端（或者服务端）安装这个 官⽅提供的pushgateway插件 然后，使⽤我们运维⾃⾏开发的各种脚本 把监控数据组织成k/v的形式 metrics形式 发送给 pushgateway

之后 pushgateway会再推送给prometheus




这种是⼀种被动的数据采集模式




— 最后 咱们来看下这个部分 报警的部分




image




prometheus 本⾝并不具备报警的功能 只能通过第三⽅ 开源或者商业软件实现报警




image

另外

这⾥指的是 prometheus 可以依赖很多 ⽅式⼆次绘制监控图形 本课程中 ⾸推Grafana




image




image