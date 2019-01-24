Prometheus中的服务发现和relabel
http://ylzheng.com/2018/01/17/prometheus-sd-and-relabel/


在默认情况下，我们从所有环境的Node Exporter中采集到的主机指标如下：

`node_cpu{cpu="cpu0",instance="172.21.0.3:9100",job="consul_sd",mode="guest"}`
其中instance为target的地址，通过instance我们可以区分主机，但是无法区分环境。

我们希望采集的指标应该是如下形式：

`node_cpu{cpu="cpu0",instance="172.21.0.3:9100",dc="dc1",job="consul_sd",mode="guest"}`


官方文档中是这样解释relabel能力的

> Relabeling is a powerful tool to dynamically rewrite the label set of a target before it gets scraped. Multiple relabeling steps can be configured per scrape configuration. They are applied to the label set of each target in order of their appearance in the configuration file.

简单理解的话，就是Relabel可以在Prometheus采集数据之前，通过Target实例的Metadata信息，动态重新写入Label的值。除此之外，我们还能根据Target实例的Metadata信息选择是否采集或者忽略该Target实例。


基于Consul动态发现的Target实例，具有以下Metadata信息：

- __meta_consul_address: consul地址
- __meta_consul_dc: consul中服务所在的数据中心
- __meta_consul_metadata_: 服务的metadata
- __meta_consul_node: 服务所在consul节点的信息
- __meta_consul_service_address: 服务访问地址
- __meta_consul_service_id: 服务ID
- __meta_consul_service_port: 服务端口
- __meta_consul_service: 服务名称
- __meta_consul_tags: 服务包含的标签信息

```
scrape_configs:
  - job_name: consul_sd
    relabel_configs:
    - source_labels:  ["__meta_consul_dc"]
      regex: "(.*)"
      replacement: $1
      action: replace
      target_label: "dc"
```