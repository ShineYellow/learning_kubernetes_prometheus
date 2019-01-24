kubernetes_sd_config

https://prometheus.io/docs/prometheus/latest/configuration/configuration/#kubernetes_sd_config

Kubernetes SD configurations allow retrieving scrape targets from `Kubernetes' REST API` and always staying synchronized with the cluster state.
See below for the configuration options for Kubernetes discovery:

```yml


# The information to access the Kubernetes API.

# The API server addresses. If left empty, Prometheus is assumed to run inside
# of the cluster and will discover API servers automatically and use the pod's
# CA certificate and bearer token file at /var/run/secrets/kubernetes.io/serviceaccount/.
[ api_server: <host> ]

# The Kubernetes role of entities that should be discovered.
role: <role>

# Optional authentication information used to authenticate to the API server.
# Note that `basic_auth`, `bearer_token` and `bearer_token_file` options are
# mutually exclusive.
# password and password_file are mutually exclusive.

# Optional HTTP basic authentication information.
basic_auth:
  [ username: <string> ]
  [ password: <secret> ]
  [ password_file: <string> ]

# Optional bearer token authentication information.
[ bearer_token: <secret> ]

# Optional bearer token file authentication information.
[ bearer_token_file: <filename> ]

# TLS configuration.
tls_config:
  [ <tls_config> ]

# Optional namespace discovery. If omitted, all namespaces are used.
namespaces:
  names:
    [ - <string> ]
```

One of the following role types can be configured to discover targets

Where <role> must be endpoints, service, pod, node, or ingress.

- Available meta labels
    - node
        - __meta_kubernetes_node_name: The name of the node object.
        - __meta_kubernetes_node_label_<labelname>: Each label from the node object.
        - __meta_kubernetes_node_annotation_<annotationname>: Each annotation from the node object.
        - __meta_kubernetes_node_address_<address_type>: The first address for each node address type, if it exists.
    - service
        - __meta_kubernetes_namespace: The namespace of the service object.
        - __meta_kubernetes_service_name: The name of the service object.
        - __meta_kubernetes_service_label_<labelname>: The label of the service object.
        - __meta_kubernetes_service_annotation_<annotationname>: The annotation of the service object.
        - __meta_kubernetes_service_port_name: Name of the service port for the target.
        - __meta_kubernetes_service_port_number: Number of the service port for the target.
        - __meta_kubernetes_service_port_protocol: Protocol of the service port for the target.
    - pod
        - __meta_kubernetes_namespace: The namespace of the pod object.
        - __meta_kubernetes_pod_name: The name of the pod object.
        - __meta_kubernetes_pod_ip: The pod IP of the pod object.
        - __meta_kubernetes_pod_label_<labelname>: The label of the pod object.
        - __meta_kubernetes_pod_annotation_<annotationname>: The annotation of the pod object.
        - __meta_kubernetes_pod_container_name: Name of the container the target address points to.
        - __meta_kubernetes_pod_container_port_name: Name of the container port.
        - __meta_kubernetes_pod_container_port_number: Number of the container port.
        - __meta_kubernetes_pod_container_port_protocol: Protocol of the container port.
        - __meta_kubernetes_pod_ready: Set to true or false for the pod's ready state.
        - __meta_kubernetes_pod_phase: Set to Pending, Running, Succeeded, Failed or Unknown in the lifecycle.
        - __meta_kubernetes_pod_node_name: The name of the node the pod is scheduled onto.
        - __meta_kubernetes_pod_host_ip: The current host IP of the pod object.
        - __meta_kubernetes_pod_uid: The UID of the pod object.
        - __meta_kubernetes_pod_controller_kind: Object kind of the pod controller.
        - __meta_kubernetes_pod_controller_name: Name of the pod controller.
    - endpoints
        - __meta_kubernetes_namespace: The namespace of the endpoints object.
        - __meta_kubernetes_endpoints_name: The names of the endpoints object.
        - __meta_kubernetes_endpoint_ready: Set to true or false for the endpoint's ready state.
        - __meta_kubernetes_endpoint_port_name: Name of the endpoint port.
        - __meta_kubernetes_endpoint_port_protocol: Protocol of the endpoint port.
        - __meta_kubernetes_endpoint_address_target_kind: Kind of the endpoint address target.
        - __meta_kubernetes_endpoint_address_target_name: Name of the endpoint address target.
    - ingress
        - __meta_kubernetes_namespace: The namespace of the ingress object.
        - __meta_kubernetes_ingress_name: The name of the ingress object.
        - __meta_kubernetes_ingress_label_<labelname>: The label of the ingress object.
        - __meta_kubernetes_ingress_annotation_<annotationname>: The annotation of the ingress object.
        - __meta_kubernetes_ingress_scheme: Protocol scheme of ingress, https if TLS config is set. Defaults to http.
        - __meta_kubernetes_ingress_path: Path from ingress spec. Defaults to /.
