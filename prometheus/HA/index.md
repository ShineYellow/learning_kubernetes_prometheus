Prometheus Server 的高可靠方案

思路：使用 remote_read 来实现 Prometheus 数据的读写分离的集群方案， 从而达到其高可用的目的，下面我将具体讲解。

从 Prometheus 1.8 开始，增加了一个叫做 remote_read 的配置，详细信息如下：

remote_read 参数用于远程读取数据，采用 http 协议。

https://songjiayang.gitbooks.io/prometheus/content/ha/prometheus.html


## HA 官方文档
https://prometheus.io/docs/prometheus/latest/configuration/configuration/#%3Cremote_read%3E
https://prometheus.io/docs/operating/integrations/#remote-endpoints-and-storage

## 如何配置
https://songjiayang.gitbooks.io/prometheus/content/ha/prometheus.html


exampel

本地运行三个 Prometheus Server, 它们分别运行在 9090, 9091, 9092 端口。 其中 9091 和 9092 主要用来收集 node_exporter 数据， 9090 用来读取 9091, 9092 收集的数据。


```
# 数据读取的 9090 的配置
remote_read:
  - url: 'http://localhost:9091/api/v1/read'
    remote_timeout: 8s
  - url: 'http://localhost:9092/api/v1/read'
    remote_timeout: 8s
# 数据收集 9091 的配置
- job_name: 'node'
  static_configs:
  - targets: ["localhost:9100"]
# 数据收集 9092 的配置
- job_name: 'node'
  static_configs:
  - targets: ["localhost:9100"]
```