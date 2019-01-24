DirectXMan12/k8s-prometheus-adapter

https://github.com/DirectXMan12/k8s-prometheus-adapter

An implementation of the `custom.metrics.k8s.io` API using Prometheus


This repository contains an implementation of the Kubernetes `resource metrics API` and `custom metrics API`.

This adapter is therefore suitable for use with the `autoscaling/v2 Horizontal Pod Autoscaler` in Kubernetes 1.6+.
It can also replace the `metrics server` on clusters that already run Prometheus and collect the appropriate metrics.


## Example
It can be found at https://github.com/luxas/kubeadm-workshop. Pay special attention to:

- Deploying the Prometheus Operator
- Setting up the custom metrics adapter and sample app


## Background
Suppose we have some `frontend webserver`, and we're trying to write a `configuration` for the Prometheus adapter so that we can `autoscale` it based on the `HTTP requests per second` that it receives.