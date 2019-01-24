


讲：Grafana 超实⽤用企业级监控绘图⼯工具的结合


第⼗⼆讲 内容

第⼗十⼆二

- Grafana 这款软件的介绍
- Grafana 下载/安装/配置/运⾏
- Grafana 设置数据源 连接prometheus_server
- Grafana 建⽴ Dashboard
- Grafana 创建 graph 成图
- Grafana graph进阶内容
- Grafana Dashboard 全局部署优化
- Grafana json备份 和 还原
- Grafana 实现报警功能 连接（4.0新功能 重要！）
- Grafana 这款软件的介绍
定义： Grafana是⼀款近⼏年新兴的 开源数据绘图⼯具平台

默认⽀持如下这么多钟 数据源作为输⼊




现如今在各⼤企业 被⼴泛使⽤中


其强⼤⽽美观的绘图能⼒ 灵活的数据链接 以及4.0之后的报警 功能 让⼈爱不释⼿


接下来 咱们就⼀起来实践⼀下吧


• Grafana 下载/安装/配置/运⾏ 官⽅⽹站： grafana.com


image


安装更简单了 直接 rpm就装完了


运⾏： 也更简单 更规范化 直接service就可以启 动


image


访问：

默认运⾏在 3000端⼜ 直接 ip:3000 即可


image


• Grafana 设置数据源 连接prometheus_server

进⼊Grafana主界⾯后

第⼀件事就是要： 设置新的数据源

让Grafana链接上 我们的prometheus_server


设置⽅法

找到左上⾓的开始按钮 并选择 Data Sources


image


image


设置连接的类型 => prometheus


设置 prometheus的URL地址 ip:9090 即可


image


• Grafana 建⽴ Dashboard

image


之后 就进⼊ new dashboard的 主界⾯


image


• Grafana 创建 graph 成图
接下来 我们来创建第⼀个Graph图 点击Graph


image


image


这个是空Graph , 我们开始给它添加内容


image


点击 Edit


image


之后 我们来到了 graph的 Edit编辑模式


⾸先 我们给这个Graph 起⼀个名字


image


点击 General 选项卡之后


image


之后 我们进⼊右边下⼀个 选项卡 Metrics


image


Data Source 选择 prometheus


image

之后 在下⾯ 输⼊ Key 名称


image

然后 我们就得到了 咱们的⼀副Graph监控图


image


• Grafana graph进阶内容
接下来 咱们来详细讲 ⼀个 graph的编辑 第⼀个选项卡 general的部分

可以设置 整个graph的分辨率

如果希望图形更清晰 可以设置 Height的部分 到200-400px


image


接下来我们看 Metrics选项卡的部分


image

默认图像下⽅的输出信息 是输出全部的标签 看着很乱 我们可以定制⾃⼰的 标签别名

使⽤ {{ 标签名称}}


image

接下来 我们来到 Legend选项卡


image


在这⾥ 我们可以进⼀步 对输出值 进⾏规划

可以讲显⽰数值 额外设置 最⼩ 最⼤ 平均 当前 数值进⾏显⽰ 最终可以做成这种样⼦的显⽰


image

然后 我们可以对 曲线图的颜⾊进⾏设置（保存后 永久⽣效） 最后我们来看下 最终效果


image


• Grafana json备份 和 还原

SAVE as


当我们的dashboard设置的graph越来越多时 我们除了即时保存 还要掌握 导出的⽅法 以便永久保存


image


在最上⾯ 设置的部分 选择 View Json


image


将全部的json copy出来 保存到第⼀个地⽅ 以备不时之需

Grafana⽀持将 json导⼊ 还原成 dashboard


image

• Grafana 实现报警功能 连接（4.0新功能 重要！）

Grafana 4.0之后 ⽀持报警功能


所有的graph图形都可以 设置连接上多个报警平台 并设置阈 值并报警


这个是⼀个⾮常重要的新功能 有了这个 新功能后 prometheus_server(alertmanager插件)-Pagerduty 仅仅作为数据 源即可（没有必要再使⽤prometheus做报警设置了）


整个的流程如下


image


接下来我们来看下 如何设置grafana的报警功能


1） 新建⼀个报警平台 连接渠道（Alerting channel） 我们先去Pagerduty 获取⼀个 Integration Key 这个Integration Key 实际上是⼀串数字，作为让其他软件连接到⾃⼰的 认证码 xxxxxxxxxxxxxxxx


然后 我们进⼊Grafana的 Alerting选项卡 点击 Notification channels


image

之后 我们新建⼀个 channel


image


然后 我们在新建channel中 设置如下


image


如此设置 让这个Channel 连接Pagerduty报警平台

并黏贴上 那⼀串从Pagerduty申请的 注册码（这个码如何申请，我们会在Pagerduty的章节 进

⾏讲解）


image


保存


2） 回到我们的 Dashboard => Graph

我们来设置 报警阈值

进⼊编辑模式后 如下 我们选择 Alert 选项卡


image


点击 image

之后我们就可以开始设置我们 对于这⼀个graph监控的 阈值设置了


image


Conditions 是作为 报警阈值的 设置项

WHEN之后 选择阈值的选取形式

可以是 avg() 也可以是 max() min() 等等

对⼀段数据的数值 进⾏估算 ，可以⾛ 平均 也可以取最⼤值 最⼩值


然后 OF 之后 选择我们的 之前的查询公式

Query(A,1m,now)

其中的 A 就是我们在创建 Graph Metrics选项卡的 设置的第⼀ 个查询公式

1m 代表的是 取1分钟内的数据


IS ABOVE 后 就是报警的触发阈值了 当⼤于2000的时候 就会发出报警到 Pagerduty 保存之后 我们的Graph图 如下


image

可以看到 图上⽅ 多了⼀条红⾊报警阈值线 当我们曲线 达到这个阈值时 就是触发报警的时刻

可以看到 在过去的12个⼩时中，只有 ⼀个时刻 真的到达了 触发点


3） 继续设置 我们的 notification


image


这⾥ 是设置 报警发送到什么地⽅


选择 我们在第⼀步 刚刚创建的 First Channel (PagerDuty的链接)


下⾯的Message 我们⾃⾏定义 报警中会显⽰的额外 内容


之后 我们的Grafana 针对这⼀张graph的阈值和报警 就设置完毕了


image


image


image



image


image


