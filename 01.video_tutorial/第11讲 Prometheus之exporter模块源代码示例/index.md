
# 第十⼀讲：Prometheus 之 exporter模块源代码示例例


本节课内容
- 编写⼀个exporter的流程
- ⽤go开发⼀个exporter的源代码 及讲解
- 对于编写exporter的个⼈建议

## 编写⼀个exporter的流程 官⽹有介绍 编写exporter的详细内容 我们这⾥给⼤家做⼀个 简单的归纳

⾸先 不同于pushgateway, exporter是⼀个独⽴运⾏ 的采集程序 其中的功能需要有这三个部分

1. ⾃⾝是HTTP 服务器，可以响应 从外发过来的 HTTP GET请求
2. ⾃⾝需要运⾏在后台，并可以定期触发 抓取本地的监控 数据
3. 返回给prometheus_server 的内容 是需要符合 prometheus规定的metrics 类型（Key-Value）
4. key-value -> prometheus(TS) values( float int ) return string *nil


接下来 给⼤家⼀份我之前⽤ go编写⼀份 exporter源代码 (⼆) ⽤go开发⼀个exporter的源代码 及讲解

[⽤go开发⼀个exporter的源代码](export.go)


// 前⾯的 scrape describe Collect 是struct类型的成员函数，这⼏ 个函数并没有直接在这个go⾥被调⽤，⽽是 MustRegister注册 进去了它们。


// http.Handle ⾥的prometheus.Handler 将上⼀部2个Mustregister 的注册 关联进⼊http.handle. 也就进⼀步注册进⼊了httpserver- listerner


// http.lister是组赛程序，只有启动后 被curl的时候 才会执⾏。 并且由于之前经过2步的注册，跟exporter的成员函数建⽴了 关联（其实是指针关联） 所以，每次被curl的时候 所有上⾯ 的成员函数都会被调⽤

// 例如上⾯的up.set(1) 就是只有被curl的时候才会被调⽤，成

员函数是作为 ⽣成metrics的⼊⼜


注： 本Go exporter 源代码 是半年前写的了 其中⼀些地⽅ 可 能需要进⼀步修改才可以正常使⽤ （本⼈go开发能⼒也⽐较 菜 - _ -）


##  个⼈建议


通过上⾯的对⼀个exporter源代码的讲解 我们也看得出来

编写⼀个 exporter 远远⽐ 写⼀个pushgateway脚本要复杂的多 的多


个⼈建议：除⾮是⼯作中真的有需要（例如：社区的exporters 都不能满⾜需求，且对于监控客户端的规范化⽐较严格） 且 对⾃⼰的编程能⼒有⾃信

那么可以⾃⾏开发new exporter


exporter-prometheus 5s GET => exporter Handler-> handler-lister 开发的过程中 要特别注意 各种⽂件句柄等等资源的 完整回 收

因为⼀个⾮常被频繁调⽤的后台程序中 ⼀旦循环中出现资源 泄漏 那么后果是很严重的

（之前其实就遇到过⼀次 资源没有正确回收 导致每⼀次 promethue_server访问过来 都造成僵⼫进程 最终服务器被拖垮 的窘境 …囧 当然了 如果是 开发⾼⼿的话 相信我的担⼼也就 多余啦～～）

所以 如果企业中 运维开发的实⼒不够过硬，我还是尽量建议

⼤家 使⽤现成的社区exporters, 配合⾃⾏开发pushgateway脚本


