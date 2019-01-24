


image



image



image


第六讲 Prometheus 初探和配置（安装测试）


第六讲内容

• Prometheus 官⽹下载

• Prometheus 开始安装

• Prometheus 启动运⾏

• Prometheus 基本配置⽂件讲解

• 安装第⼀个exporter =》 node_exporter

• Prometheus 连接exporter获取数据

• Prometheus 命令⾏⼊门 第⼀个数学查询公式

之前我们说过，相对于较难的使⽤⽅法，prometheus其实安装的过程 反⽽是异常的简单


安装Prometheus之前 我们必须先安装 ntp时间同步

（prometheus T_S 对系统时间的准确性要求很⾼，必须保证本机时间实时同步） 以Centos7 为例

timedatectl set-timezone Asia/Shanghai


• * * * * * ntpdate -u cn.pool.ntp.org



1） Prometheus下载


⾸先 我们去到http://prometheus.io 官⽹


下载最新版本 prometheus-2.1.0.linux-amd64.tar.gz


wget https://github.com/prometheus/prometheus/releases/download/v2.1.0/prometheus-2.1.0.linux- amd64.tar.gz


2） Prometheus的安装 ⾮常简单

[root@server01 download]# tar -xvzf prometheus-2.0.0.linux-amd64.tar.gz prometheus-2.0.0.linux-amd64/

prometheus-2.0.0.linux-amd64/consoles/

prometheus-2.0.0.linux-amd64/consoles/index.html.example prometheus-2.0.0.linux-amd64/consoles/node-cpu.html prometheus-2.0.0.linux-amd64/consoles/node-disk.html prometheus-2.0.0.linux-amd64/consoles/node-overview.html prometheus-2.0.0.linux-amd64/consoles/node.html

prometheus-2.0.0.linux-amd64/consoles/prometheus-overview.html prometheus-2.0.0.linux-amd64/consoles/prometheus.html prometheus-2.0.0.linux-amd64/console_libraries/

prometheus-2.0.0.linux-amd64/console_libraries/menu.lib prometheus-2.0.0.linux-amd64/console_libraries/prom.lib prometheus-2.0.0.linux-amd64/prometheus.yml prometheus-2.0.0.linux-amd64/LICENSE

prometheus-2.0.0.linux-amd64/NOTICE prometheus-2.0.0.linux-amd64/prometheus prometheus-2.0.0.linux-amd64/promtool


cp -rf prometheus-2.0.0.linux-amd64 /usr/local/prometheus


3) Prometheus 启动 和 后台运⾏

启动也很简单


./prometheus


image


之后默认运⾏在 9090


浏览器可以直接打开访问 ⽆账号密码验证 （如果希望加上 验证 ，可以使⽤类似apache httppass ⽅式添加）


image

4） 接下来 我们来简单看⼀下 Prometheus的主配置⽂件


其实prometheus解压安装之后，就默认⾃带了⼀个 基本的配置⽂件如下

./prometheus —yml


image


我们来⼤致讲解⼀下 配置⽂件的内容


# my global config global:

scrape_interval: 15s # Set the scrape interval to every 15 seconds. Default is every 1 minute. evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.

# scrape_timeout is set to the global default (10s).


前两个全局变量

scrape_interval. 抓取采样数据的 时间间隔， 默认 每15秒去被监控机上 采样⼀次 => 5s


这个就是我们所说的 prometheus的⾃定义 数据采集频率了 evaluation_interval. 监控数据规则的评估频率 grafana 这个参数是prometheus多长时间 会进⾏⼀次 监控规则的评估

举个例： 假如 我们设置 当 内存使⽤量 > 70%时 发出报警 这么⼀条rule（规则）

那么prometheus 会默认 每15秒来执⾏⼀次这个规则 检查内存的情况


# Alertmanager configuration alerting:

alertmanagers:

- static_configs:

- targets:

# - alertmanager:9093


Alertmanager 是prometheus的⼀个⽤于管理和发出报警的 插件


我们这⾥对 Alertmanger 暂时先不做介绍 暂时也不需要 （我们采⽤ 4.0最新版的 Grafana , 本

⾝就已经⽀持报警发出功能了 往后我们会学习到）


再往后 从这⾥开始 进⼊prometheus重要的 配置采集节点的设置


# Here it's Prometheus itself. scrape_configs:

# The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.

- job_name: 'prometheus'


# metrics_path defaults to '/metrics'

# scheme defaults to 'http'.


static_configs:

- targets: ['localhost:9090']


先定义⼀个 job的名称


image


然后 定义监控节点 targets


image

- targets的设定

以这种形式设定 默认带了⼀个 prometheus本机的 static_configs:

- targets: ['localhost:9090']


这⾥可以继续 扩展加⼊ 其他需要被监控的节点 如下是⼀个 ⽣产配置例⼦

- job_name: 'aliyun' static_configs:

- targets: [‘server04:9100’,'IP:9100’,’nginx06:9100','web7:9100’,'redis1:9100','log: 9100','redis2:9100']

prometheuserver _ /etc/hosts, local_dns server


可以看到 targets可以并列写⼊ 多个节点


⽤逗号隔开， 机器名+端⼜号


端⼜号：通常⽤的就是 exporters 的端⼜ 在这⾥ 9100 其实是 node_exporter 的默认端⼜


如此 prometheus就可以通过配置⽂件 识别监控的节点，持续开始采集数据 prometheus到此就算初步的搭建好了 （后续课程 还有各种扩展的 优化）


5） 光搭建好prometheus_server 是不够的，我们需要给监控节点 搭建第⼀个exporter ⽤来采 样数据


我们就选⽤ 企业中最常⽤的 node_exporter 这个插件


node_exporter 之前介绍过，是⼀个以http_server⽅式运⾏在后台，并且持续不断采集 Linux

系统中 各种操作系统本⾝相关的监控参数的程序 其采集量是很⼤很全的，往往默认的采集项⽬ 就远超过你的实际需求 接下来我们来看下 node_exporter是怎么回事

⼀样 先下载node_exporter 从官⽹

https://prometheus.io/download/#node_exporter


下载之后 解压缩 然后直接运⾏即可


image


node_exporter的运⾏更加简单 如上图所⽰


运⾏起来以后 我们使⽤netstats -tnlp 可以来看下 node_exporter进程的状态


image


这⾥就可以看出 node_exporter默认⼯作在9100端⼜ 可以响应 prometheus_server发过来的 HTTP_GET请求


也可以响应其他⽅式的 HTTP_GET请求 我们⾃⼰就可以发送 测试


image


image


执⾏curl之后，我们看到 node_exporter 给我们返回了 ⼤量的这种 metrics类型 K/V数据 关于 metrics 和 k/v 之前咱们介绍过了 就不再重复了


⽽这些 返回的 K/V数据 ，其中的Key的名称 就可以直接复制黏贴 在prometheus的查询命令

⾏ 来查看结果了 我们来试⼀试

就⽤这⼀项看看 node_memory_MemFree


image



image

直接就可以看到曲线了 对吗？

其实这个 就是 最简单的 来查看⼀下 服务器的 空闲内存状态的⽅式


觉得很简单对吗？ prometheus不也就如此⽽已吗 没什么难的 我暗中呵呵。。。。 -_- 好吧 那么我们接下来 再来看下⼀个 难⼀些的例⼦ CPU使⽤率 的获取⽅式 node_cpu这个 key 也是node_exporter返回的⼀个 ⽤来统计 CPU使⽤率的


image


image


输⼊查询框之后， 咱们来看⼀下结果


image


疑？？？ 怎么回事 CPU不是应该是使⽤率吗 ？ 类似 百分 50% 80%这样的数据才对啊


image


怎么返回是⼀个持续不断累加的 近似于直线的 庞⼤的数字呢？


image


说到这⼉，我们就不得不给⼤家提⼀句了 这个其实就关系到 prometheus对Linux数据的采集 的精细的特性

其实 prometheus对Linux CPU的采集 并不是 直接给我们返回⼀个 现成的CPU百分⽐ ⽽是 返

回 Linux中 很底层的 cpu时间⽚ 累积数值 的这样⼀个数据（


我们平时⽤惯了 top / uptime这种简便的⽅式 看CPU使⽤率， 往往浅尝辄⽌ 根本没有好好深

⼊理解 所谓的CPU使⽤率 在linux中到底是怎么回事） 当懒⼈当的时间长了 ⾃⼰问问⾃⼰对Linux到底了解多少呢？


⼤⽶给⼤家补补课吧 :)


其实 如果想真的弄明⽩CPU的使⽤率这个概念 在Linux中 要先从CPU时间 这个概念开始建

⽴


Linux中 CPU时间 实际是指 ： 从操作系统开启算起 CPU就开始⼯作了 并记录⾃⼰在⼯作 中 总共使⽤的 “时间”的累积量 把它保存在系统中


⽽累积的CPU使⽤时间 还会分成⼏个重要的状态类型


⽐如 CPU time => 分成 CPU user time / sys time / nice time / idle time / irq / 等等。。。 翻译过来就是 CPU ⽤户态使⽤时间， 系统/内核态使⽤时间， nice值分配使⽤时间， 空闲 时间， 中断时间 等等


那么所谓的 CPU使⽤率 是什么意思呢？


CPU使⽤率 最准确的定义 其实是 CPU各种状态中 除了idle(空闲) 这个状态外， 其他所有的

CPU状态的 加合 / 总CPU时间


得出来的 就是我们所说的 CPU使⽤率 (运⾏的任务 ⽤户 内核)


那么 回到我们刚才 使⽤ node_cpu 这个key 如果直接输⼊进去


他返回的 其实是 CPU各个核CPU各个核 各个状态下 从开机开始 ⼀直累积下来的 CPU使⽤ 时间 的累积值

所以我们才 看到这么⼀个 image 如下 截取⼀段⽹上 对各个CPU状态的时间单位的解释


image


所以： 如果在prometheus中 想对CPU的使⽤率 准确的来查询 正确的⽅法如下：（难的开始 要来了。。。。） node_cpu

(1-((sum(increase(node_cpu{mode="idle"}[1m])) by (instance)) / (sum(increase(node_cpu[1m])) by (instance)))) * 100


image


这才是 真正的CPU使⽤率 V%


image

呵呵… 现在还觉得promethues 和 Linux 很简单吗？


其实话说回来 ， prometheus的这种精细的 底层的计算特性 虽然学起来难 不过带来的 好处也是 显⽽易见


1） prometheus 这种底层数据采集 所形成的监控 其实是最准确 最可信的

2） prometheus 本⾝也逼着使⽤它的运维同学 你不踏实下来 好好的真正把Linux技术学过关 的话 你就没有办法使⽤好 这个超强⼒的监控⼯具了


所以说呢 我们还是踏踏实实的 ⼀步⼀个脚印的来学习吧 加油～


最后再额外提两句


1） 最后使⽤的 这个复杂的cpu计算 promtheus公式 ， 我们从下⼀讲开始 会给⼤家 详细讲解

（不⽤着急哈）

2） 对于Linux操作系统 底层的各种深⼊知识学习 请关注 ⼤⽶运维课堂的 系列课程 第三 /

第四阶段 ⾼级=>专级课程 (⾯向 5-10年经验⼯程师)

（今天随⼜提的这个CPU时间计算 只不过是Linux开始深⼊的 冰⼭⼀⾓⽽已）


本节课结束 更多内容 请关注后续
