Prometheus Operator — 為 Kubernetes 設定及管理 Prometheus

https://medium.com/getamis/kubernetes-operators-prometheus-3584edd72275


## kube-prometheus
kube-prometheus 有點像是一個 wrapper，安裝它的同時，會順便把 Prometheus/Alertmanager/Grafana 以及各種 K8S 內要用到的 Exporter 都一次安裝完畢！


## helm repo add coreos 
helm repo add coreos https://s3-eu-west-1.amazonaws.com/coreos-charts/stable/

## helm install
```
# 按照預設組態安裝
~$ helm install coreos/kube-prometheus --name kube-prometheus --namespace monitoring  
# 或是讀取客製化組態安裝 (可略過)
~$ helm install coreos/kube-prometheus --name kube-prometheus --namespace monitoring --values kube-prometheus.yaml
NAME:   kube-prometheus
LAST DEPLOYED: Sun May 27 22:16:11 2018
NAMESPACE: monitoring
STATUS: DEPLOYED
RESOURCES:
==> v1/Secret
NAME                          TYPE    DATA  AGE
alertmanager-kube-prometheus  Opaque  1     1s
kube-prometheus-grafana       Opaque  2     1s
...
```