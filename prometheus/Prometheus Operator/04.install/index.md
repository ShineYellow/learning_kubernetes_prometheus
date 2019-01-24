## install by yaml
```
git clone https://github.com/coreos/prometheus-operator.git

cd prometheus-operator/contrib/kube-prometheus 

kubectl apply -f manifests/

```

## install by helm
address

https://github.com/coreos/prometheus-operator/tree/master/helm

helm repo add coreos https://s3-eu-west-1.amazonaws.com/coreos-charts/stable/

helm install coreos/kube-prometheus --name kube-prometheus --namespace monitoring  


