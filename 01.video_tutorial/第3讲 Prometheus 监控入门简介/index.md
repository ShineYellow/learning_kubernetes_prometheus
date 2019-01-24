


image


image


image

第三讲 Prometheus 监控入门简介


第三讲内容
• prometheus是什么？

• prometheus能为我们带来些什么

• prometheus对于运维的要求

• prometheus多图效果展⽰



1) Prometheus是什么


我们还是来先看看 官⽅对它的原始定义

prometheus is an open-source systems monitoring and alerting toolkit originally built at soundcloud , Since its inception in 2012, many companies and organizations have adopted Prometheus, and the project has a very active developer and user community. It is now a standalone open source project and maintained independently of any company. To emphasize this, and to clarify the project's governance structure, Prometheus joined the Cloud Native Computing Fondation in 2016 as the second hosted project, after Kubernets.


⼤大⽶米来帮翻译⼀一下哈=》


prometheus 是⼀一个开源系统监控和报警的⼯工具集合，由SoundCloud创建（http:// soundcloud.com/ ） ， ⾃自从2012诞⽣生之后，⾄至今已经有许多公司和组织开始使⽤用它了了，这个 开源项⽬目拥有⼤大量量的积极参与开发和建设的 研发⼈人员以及社区⽤用户. ⽬目前 已经是⼀一个独⽴立运

⾏行行的 开源的 由各公司⾃自⾏行行维护的监控项⽬目。 为了了让项⽬目更更充实更更清晰 2016年年 prometheus 加⼊入了了 Cloud Native Computing Fondation （CNCF）， 并且成为继Kubernets(结合 容器器/ docker)之后 第⼆二个加⼊入该组织的成员


这个就是来⾃于官⽅的介绍


其中更多突出的 还是这个项⽬基于开源的 和 各种社区组织维护 多重联合开发 的这样⼀个 特性

也就决定了 这个项⽬ 必然是越来越好


当前最新的版本Prometheus 2.0 也是近期刚推出不久，经过本⼈最近⼏个⽉的对最新版本的 实践 各种之前遇到的问题 差不多都被修复了 ⽽且性能更加稳定

（⼤⽶⽼师当前所在的企业 也是以proemtheus 作为核⼼监控，新版本运⾏ ⼏个⽉来 没有出 现过任何稳定和性能上的问题）


2） prometheus能给我们带来什么


之前在第⼆讲最后的部分 我们给⼤家推出了⼀个 最终的未来理想化监控

那么prometheus 可以针对未来监控 对于 准确性 和 精确性的 要求 极⼤的贡献⾃⼰的⼀份⼒

量

为什么这么说呢？ 我们不能凭空判断，咱们⽤真实数据来说话 认识promethues监控的优质特性

• 基于time series 时间序列列模型 (数字 数学)


时间序列列（time series X,Y）是⼀一系列列有序的数据。通常是等时间间隔的采样数据。


• 基于K/V 的数据模型

Key/value 这个 键值的概念 咱们很熟悉


{ disk_size : 80 }


最⼤的好处 就是数据格式 简单 速度快 易维护开发


• 采样数据的查询 完全基于数学运算 ⽽不是其他的表达式 并提供专有的查询输⼊

console

这个特点很独特，所有的查询都基于数学运算公式 例如 (增量（A） + 增量（B） ) / 总增量

（C） > 固定百分⽐ =>


• 采⽤ HTTP pull / push 两种对应的数据采集传输⽅式

所有的数据采集 都基本采⽤ HTTP ，⽽且分为 pull / push 推和拉两种⽅式去写采集程序 ⽅ 便


• 开源，且⼤量的社区成品插件

这个⾮常厉害，https://prometheus.io 很多prometheus 社区开发的差距已经异常强⼤和完善 如果 公司对监控要求的不是特别⾼的话，默认的⼏个成品插件 装上就可以⽤到底了 监控成型速度太快了


• push的⽅法 ⾮常⾮常的灵活

pull push的⽅法 我们往后会具体介绍，这⾥提⼀句

push的这种采集⽅法 灵活程度超过你的想象 ⼏乎任何形式的数据 都可以实现


• 本⾝⾃带图形调试(sql) prometheus(查询语句)本⾝的就再带了 现成的图形成型界⾯

虽然最终肯定不能跟grafana的效果相⽐，但是 这种⾃带图形成图 可以⼤⼤帮助运维做 调试


• 最精细的数据采样

⼤多数市⾯上的开源监控 采样也就能精确到 半分钟 ⼀分钟的程度

商品化监控产品 就更别提了 （为了缩⼩数据存储的成本） 有的甚⾄ 5分钟 就是采样最⼩间 隔

prometheus 理论上 可以达到 每1秒采集！！⽽且可以⾃⾏定制频率 （不过 强⼤的同时 其实

⼤⽶ 不太建议 细到这个程度 因为数据量太⼤了 如果1s采样⼀次）


• 最新的⾼⼤上监控 很霸⽓ ⾯试很有料

呵呵 这个就⽐较个⼈主义了 ， 运维⾯试的时候 监控⽅⾯ 如果你懂得这么个⾼⼤上的⼯ 具，并且对各种监控算法 系统底层相关都很熟悉（prometheus逼着你必须熟悉底层） 的 话，优势很⼤哦


⽐较显著的特点就介绍如上了

另外 ：咱们也不能盲⽬崇拜 也需要客观的对待prometheus prometheus 还是又⼀些不⾜ 有待于改进

• 不⽀持集群化 （这个是当前最迫切的需求）

• 被监控集群过⼤后 本⾝性能有⼀定瓶颈（如果有集群 就可以解决这个问题）

• 偶尔发⽣数据丢失（这个问题 在2.0之前 会偶尔发⽣⼏次， 2.0之后貌似已经彻底解 决 ⾄少⼤⽶这⼏个⽉ 没有看到丢失）

• 中⽂⽀持不好 中⽂资料也很少（这个问题 也是⽼⽣常谈了 往往新的 很⽜的国外⼯具 都不太⽀持中⽂）

3) prometheus 对于运维的要求
• 要求对操作系统有很深⼊扎实的知识 不能只是浮在表⾯ /proc/ nagios zabbix

• 对数学思维有⼀定的要求 因为它基本的内核就是数学公式组成 T-S (四则运算，算法 -

> 微积分， 代数 数论)

• 对监控的经验有很⾼的要求 很多时候 监控项需要很细的定制


4） prometheus 各种图形展⽰

• prometheus主界⾯



image


• promtheus 数学查询命令⾏展⽰

image


(1-((sum(increase(node_cpu{mode="idle"}[1m])) by (instance)) /(sum(increase(node_cpu[1m])) by (instance)))) * 100 (较复杂数学公式查询展⽰)


image


• prometheus 配置展⽰


image


• prometheus targets展⽰（被监控节点）


image


• prometheus + grafana 监控CPU展⽰


image


image


• promtheus ⽂件描述符监控 展⽰


image


image



image
