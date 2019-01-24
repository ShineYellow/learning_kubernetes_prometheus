

ServiceMonitor

https://itnext.io/kubernetes-monitoring-with-prometheus-in-15-minutes-8e54d1de2e13

Prometheus-operator uses a Custom Resource Definition (CRD), named ServiceMonitor, to abstract the configuration to target. As an example below, let’s see how to monitor a NGINX pod with ServiceMonitor. The ServiceMonitor will select the NGINX pod, using the matchLabels selector. The prometheus-operator will search for the pods based on the label selector and creates a prometheus target so prometheus will scrape the metrics endpoint.

```
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  namespaceSelector:
    matchNames:
    - default
  endpoints:
  - port: web
    interval: 30s
```


----


## ServiceMonitor

https://akomljen.com/get-kubernetes-cluster-metrics-with-prometheus-in-5-minutes/


From the picture above you can see that you can create a `ServiceMonitor` resource which will `scrape` the Prometheus `metrics` from the defined set of pods. It instructs Prometheus to `watch` on a new target. For example, if you have a frontend app which exposes Prometheus metrics on port web, you can create a service monitor which will configure the Prometheus server automatically:

```
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: frontend-app
  labels:
    app: frontend-app
    release: prom
spec:
  namespaceSelector:
    any: true
  selector:
    matchLabels:
      app: frontend-app
  endpoints:
  - port: web
    interval: 10s
```  

## example