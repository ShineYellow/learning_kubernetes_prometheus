## kube-state-metrics 介绍

https://www.kancloud.cn/huyipow/prometheus/521184

kube-state-metrics is a simple service that listens to the Kubernetes API server and generates metrics about the state of the objects. (See examples in the Metrics section below.) It is not focused on the health of the individual Kubernetes components, but rather on the health of the various objects inside, such as deployments, nodes and pods.

The metrics are exported through the Prometheus golang client on the HTTP endpoint /metrics on the listening port (default 8080). They are served either as plaintext or protobuf depending on the Accept header. They are designed to be consumed either by Prometheus itself or by a scraper that is compatible with scraping a Prometheus client endpoint. You can also open /metrics in a browser to see the raw metrics.

- CronJob Metrics
- DaemonSet Metrics
- Deployment Metrics
- Job Metrics
- LimitRange Metrics
- Node Metrics
- PersistentVolume Metrics
- PersistentVolumeClaim Metrics
- Pod Metrics
- ReplicaSet Metrics
- ReplicationController Metrics
- ResourceQuota Metrics
- Service Metrics
- StatefulSet Metrics
- Namespace Metrics
- Horizontal Pod Autoscaler Metrics
- Endpoint Metrics
- kube-state-metricsgithub 项目地址

查看kube-state配置文件
```
    # tree kube-state-metrics/
    kube-state-metrics/
    ├── kube-state-metrics-cluster-role-binding.yaml
    ├── kube-state-metrics-cluster-role.yaml
    ├── kube-state-metrics-deployment.yaml
    ├── kube-state-metrics-role-binding.yaml
    ├── kube-state-metrics-role.yaml
    ├── kube-state-metrics-service-account.yaml
    ├── kube-state-metrics-service.yaml
    └── prometheus-k8s-service-monitor-kube-state-metrics.yaml
```
    0 directories, 8 files
执行创建 kube-state-metrics

kubectl apply -f kube-state-metrics/
检查服务启动状态
```
 [root@app553 ~]# kubectl get pod,svc -n monitoring
NAME                                      READY     STATUS    RESTARTS   AGE
po/kube-state-metrics-678d94f978-nv4g5    2/2       Running   0          52m
po/prometheus-k8s-0                       2/2       Running   0          4d
po/prometheus-k8s-1                       2/2       Running   0          4d
po/prometheus-operator-68589bfbfd-pf4tk   1/1       Running   0          4d
po/rbd-provisioner-698ddf4568-w7s29       1/1       Running   0          4d
```