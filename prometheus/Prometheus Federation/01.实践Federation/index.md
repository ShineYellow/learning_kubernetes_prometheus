
搭建了一套 Prometheus Federation

```
Name				Role				Port
Prometheus-global	Master				9090
Prometheus-node		Collector			9091
Prometheus-docker	Collector			9092

```

新增一個檔案prometheus-docker.yml，並加入以下內容:
```
global:
  scrape_interval:     15s
  evaluation_interval: 15s
  external_labels:
      server: 'docker-monitor'
scrape_configs:
  - job_name: 'docker'
    scrape_interval: 5s
    static_configs:
      - targets: ['localhost:9323']
```
新增一個檔案prometheus-node.yml，並加入以下內容:
```
global:
  scrape_interval:     15s
  evaluation_interval: 15s
  external_labels:
      server: 'node-monitor'
scrape_configs:
  - job_name: 'node'
    scrape_interval: 5s
    static_configs:
      - targets: ['localhost:9100']
```
新增一個檔案prometheus-global.yml，並加入以下內容:
```
global:
  scrape_interval:     15s
  evaluation_interval: 15s
  external_labels:
      server: 'global-monitor'
scrape_configs:
  - job_name: 'federate'
    scrape_interval: 15s
    honor_labels: true
    metrics_path: '/federate'
    params:
      'match[]':
        - '{job=~"prometheus.*"}'
        - '{job="docker"}'
        - '{job="node"}'
    static_configs:
      - targets:
        - 'localhost:9091'
        - 'localhost:9092'
```

具体参考
02.了解 Prometheus Federation 功能
