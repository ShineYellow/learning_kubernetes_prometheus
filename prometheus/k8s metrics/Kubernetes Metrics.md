
## Kubernetes Metrics

https://docs.datadoghq.com/agent/kubernetes/metrics/

### Kubernetes
| key                                            | desc                                            |
|------------------------------------------------|-------------------------------------------------|
| kubernetes.cpu.capacity                        | The number of cores in this machine             |
| kubernetes.cpu.usage.total                     | The number of cores used                        |
| kubernetes.cpu.limits                          | The limit of cpu cores set                      |
| kubernetes.cpu.requests                        | The requested cpu cores                         |
| kubernetes.filesystem.usage                    | The amount of disk used                         |
| kubernetes.filesystem.usage_pct                | The percentage of disk used                     |
| kubernetes.memory.capacity                     | The amount of memory (in bytes) in this machine |
| kubernetes.memory.limits                       | The limit of memory set                         |
| kubernetes.memory.requests                     | The requested memory                            |
| kubernetes.memory.usage                        | The amount of memory used                       |
| kubernetes.network.rx_bytes                    | The amount of bytes per second received         |
| kubernetes.network.tx_bytes                    | The amount of bytes per second transmitted      |
| kubernetes.network_errors                      | The amount of network errors per second         |
| kubernetes.diskio.io_service_bytes.stats.total | The amount of disk space the container uses.    |



### Kubelet  
| key                                                | desc                                                                                  |
|----------------------------------------------------|---------------------------------------------------------------------------------------|
| kubernetes.containers.last_state.terminated        | The number of containers that were previously terminated                              |
| kubernetes.pods.running                            | The number of running pods                                                            |
| kubernetes.containers.running                      | The number of running containers                                                      |
| kubernetes.containers.restarts                     | The number of times the container has been restarted                                  |
| kubernetes.containers.state.terminated             | The number of currently terminated containers                                         |
| kubernetes.containers.state.waiting                | The number of currently waiting containers                                            |
| kubernetes.cpu.load.10s.avg                        | Container cpu load average over the last 10 seconds                                   |
| kubernetes.cpu.system.total                        | The number of cores used for system time                                              |
| kubernetes.cpu.user.total                          | The number of cores used for user time                                                |
| kubernetes.cpu.cfs.throttled.period                | Number of throttled period intervals                                                  |
| kubernetes.cpu.cfs.throttled.second                | Total time duration the container has been throttled                                  |
| kubernetes.cpu.capacity                            | The number of cores in this machine                                                   |
| kubernetes.cpu.usage.total                         | The number of cores used                                                              |
| kubernetes.cpu.limits                              | The limit of cpu cores set                                                            |
| kubernetes.cpu.requests                            | The requested cpu cores                                                               |
| kubernetes.filesystem.usage                        | The amount of disk used                                                               |
| kubernetes.filesystem.usage_pct                    | The percentage of disk used                                                           |
| kubernetes.io.read_bytes                           | The amount of bytes read from the disk                                                |
| kubernetes.io.write_bytes                          | The amount of bytes written to the disk                                               |
| kubernetes.memory.capacity                         | The amount of memory (in bytes) in this machine                                       |
| kubernetes.memory.limits                           | The limit of memory set                                                               |
| kubernetes.memory.requests                         | The requested memory                                                                  |
| kubernetes.memory.usage                            | Current memory usage in bytes including all memory regardless of when it was accessed |
| kubernetes.memory.working_set                      | Current working set in bytes - this is what the OOM killer is watching for            |
| kubernetes.memory.rss                              | Size of RSS in bytes                                                                  |
| kubernetes.memory.usage_pct                        | The percentage of memory used                                                         |
| kubernetes.network.rx_bytes                        | The amount of bytes per second received                                               |
| kubernetes.network.rx_dropped                      | The amount of rx packets dropped per second                                           |
| kubernetes.network.rx_errors                       | The amount of rx errors per second                                                    |
| kubernetes.network.tx_bytes                        | The amount of bytes per second transmitted                                            |
| kubernetes.network.tx_dropped                      | The amount of tx packets dropped per second                                           |
| kubernetes.network.tx_errors                       | The amount of tx errors per second                                                    |
| kubernetes.diskio.io_service_bytes.stats.total     | The amount of disk space the container uses                                           |
| kubernetes.apiserver.certificate.expiration.count  | The count of remaining lifetime on the certificate used to authenticate a request     |
| kubernetes.apiserver.certificate.expiration.sum    | The sum of remaining lifetime on the certificate used to authenticate a request       |
| kubernetes.rest.client.requests                    | The number of HTTP requests                                                           |
| kubernetes.rest.client.latency.count               | The count of request latency in seconds broken down by verb and URL                   |
| kubernetes.rest.client.latency.sum                 | The sum of request latency in seconds broken down by verb and URL                     |
| kubernetes.kubelet.runtime.operations              | The number of runtime operations                                                      |
| kubernetes.kubelet.runtime.errors                  | The number of runtime operations errors                                               |
| kubernetes.kubelet.network_plugin.latency.sum      | The sum of latency in microseconds of network plugin operations                       |
| kubernetes.kubelet.network_plugin.latency.count    | The count of network plugin operations by latency                                     |
| kubernetes.kubelet.network_plugin.latency.quantile | The quantiles of network plugin operations by latency                                 |
| kubernetes.kubelet.volume.stats.available_bytes    | The number of available bytes in the volume                                           |
| kubernetes.kubelet.volume.stats.capacity_bytes     | The capacity in bytes of the volume                                                   |
| kubernetes.kubelet.volume.stats.used_bytes         | The number of used bytes in the volume                                                |
| kubernetes.kubelet.volume.stats.inodes             | The maximum number of inodes in the volume                                            |
| kubernetes.kubelet.volume.stats.inodes_free        | The number of free inodes in the volume                                               |
| kubernetes.kubelet.volume.stats.inodes_used        | The number of used inodes in the volume                                               |



###  kube-state-metrics     
        
Note that kubernetes_state.* metrics are gathered from the kube-state-metrics API.     
| key                                                           | desc                                                                         |
|---------------------------------------------------------------|------------------------------------------------------------------------------|
| kubernetes_state.container.ready                              | Whether the containers readiness check succeeded                             |
| kubernetes_state.container.running                            | Whether the container is currently in running state                          |
| kubernetes_state.container.terminated                         | Whether the container is currently in terminated state                       |
| kubernetes_state.container.status_report.count.terminated     | Count of the containers currently reporting a in ...                         |
| kubernetes_state.container.waiting                            | Whether the container is currently in waiting state                          |
| kubernetes_state.container.status_report.count.waiting        | Count of the containers currently reporting a in ...                         |
| kubernetes_state.container.gpu.request                        | The number of requested gpu devices by a container                           |
| kubernetes_state.container.gpu.limit                          | The limit on gpu devices to be used by a container                           |
| kubernetes_state.container.restarts                           | The number of restarts per container                                         |
| kubernetes_state.container.cpu_requested                      | The number of requested cpu cores by a container                             |
| kubernetes_state.container.memory_requested                   | The number of requested memory bytes by a container                          |
| kubernetes_state.container.cpu_limit                          | The limit on cpu cores to be used by a container                             |
| kubernetes_state.container.memory_limit                       | The limit on memory to be used by a container                                |
| kubernetes_state.daemonset.scheduled                          | The number of nodes running at least one daemon pod and that are supposed to |
| kubernetes_state.daemonset.misscheduled                       | The number of nodes running a daemon pod but are not supposed to             |
| kubernetes_state.daemonset.desired                            | The number of nodes that should be running the daemon pod                    |
| kubernetes_state.daemonset.ready                              | The number of nodes that should be running the daemon pod ...                |
| kubernetes_state.deployment.replicas                          | The number of replicas per deployment                                        |
| kubernetes_state.deployment.replicas_available                | The number of available replicas per deployment                              |
| kubernetes_state.deployment.replicas_unavailable              | The number of unavailable replicas per deployment                            |
| kubernetes_state.deployment.replicas_updated                  | The number of updated replicas per deployment                                |
| kubernetes_state.deployment.replicas_desired                  | The number of desired replicas per deployment                                |
| kubernetes_state.deployment.paused                            | Whether a deployment is paused                                               |
| kubernetes_state.deployment.rollingupdate.max_unavailable     | Maximum number of unavailable replicas during a rolling update               |
| kubernetes_state.endpoint.address_available                   | Number of addresses available in endpoint                                    |
| kubernetes_state.endpoint.address_not_ready                   | Number of addresses not ready in endpoint                                    |
| kubernetes_state.endpoint.created                             | Unix creation timestamp                                                      |
| kubernetes_state.job.status.failed                            | Observed number of failed pods in a job                                      |
| kubernetes_state.job.status.succeeded                         | Observed number of succeeded pods in a job                                   |
| kubernetes_state.limitrange.cpu.min                           | Minimum CPU request for this type                                            |
| kubernetes_state.limitrange.cpu.max                           | Maximum CPU limit for this type                                              |
| kubernetes_state.limitrange.cpu.default                       | Default CPU limit if not specified                                           |
| kubernetes_state.limitrange.cpu.default_request               | Default CPU request if not specified                                         |
| kubernetes_state.limitrange.cpu.max_limit_request_ratio       | Maximum CPU limit / request ratio                                            |
| kubernetes_state.limitrange.memory.min                        | Minimum memory request for this type                                         |
| kubernetes_state.limitrange.memory.max                        | Maximum memory limit for this type                                           |
| kubernetes_state.limitrange.memory.default                    | Default memory limit if not specified                                        |
| kubernetes_state.limitrange.memory.default_request            | Default memory request if not specified                                      |
| kubernetes_state.limitrange.memory.max_limit_request_ratio    | Maximum memory limit / request ratio                                         |
| kubernetes_state.node.cpu_capacity                            | The total CPU resources of the node                                          |
| kubernetes_state.node.memory_capacity                         | The total memory resources of the node                                       |
| kubernetes_state.node.pods_capacity                           | The total pod resources of the node                                          |
| kubernetes_state.node.gpu.cards_allocatable                   | The GPU resources of a node that are available for scheduling                |
| kubernetes_state.node.gpu.cards_capacity                      | The total GPU resources of the node                                          |
| kubernetes_state.persistentvolumeclaim.status                 | The phase the persistent volume claim is currently in                        |
| kubernetes_state.persistentvolumeclaim.request_storage        | Storage space request for a given pvc                                        |
| kubernetes_state.persistentvolume.by_phase                    | Number of persistent volumes to sum by phase and storageclass                |
| kubernetes_state.node.cpu_allocatable                         | The CPU resources of a node that are available for scheduling                |
| kubernetes_state.node.memory_allocatable                      | The memory resources of a node that are available for scheduling             |
| kubernetes_state.node.pods_allocatable                        | The pod resources of a node that are available for scheduling                |
| kubernetes_state.node.status                                  | Submitted with a value of 1 for each node and tagged either...               |
| kubernetes_state.nodes.by_condition                           | To sum by `condition` and `status` to get numbe...                           |
| kubernetes_state.hpa.min_replicas                             | Lower limit for the number of pods that can be set by the autoscaler         |
| kubernetes_state.hpa.max_replicas                             | Upper limit for the number of pods that can be set by the autoscaler         |
| kubernetes_state.hpa.desired_replicas                         | Desired number of replicas of pods managed by this autoscaler                |
| kubernetes_state.pod.ready                                    | In association with the `condition` tag, ....                                |
| kubernetes_state.pod.scheduled                                | Reports the status of the scheduling process for the pod with its tags       |
| kubernetes_state.pod.status_phase                             | To sum by `phase` to get number of pods ...                                  |
| kubernetes_state.replicaset.replicas                          | The number of replicas per ReplicaSet                                        |
| kubernetes_state.replicaset.fully_labeled_replicas            | The number of fully labeled replicas per ReplicaSet                          |
| kubernetes_state.replicaset.replicas_ready                    | The number of ready replicas per ReplicaSet                                  |
| kubernetes_state.replicaset.replicas_desired                  | Number of desired pods for a ReplicaSet                                      |
| kubernetes_state.replicationcontroller.replicas               | The number of replicas per ReplicationController                             |
| kubernetes_state.replicationcontroller.fully_labeled_replicas | The number of fully labeled replicas per ReplicationController               |
| kubernetes_state.replicationcontroller.replicas_ready         | The number of ready replicas per ReplicationController                       |
| kubernetes_state.replicationcontroller.replicas_desired       | Number of desired replicas for a ReplicationController                       |
| kubernetes_state.replicationcontroller.replicas_available     | The number of available replicas per ReplicationController                   |
| kubernetes_state.resourcequota.pods.used                      | Observed number of pods used for a resource quota                            |
| kubernetes_state.resourcequota.services.used                  | Observed number of services used for a resource quota                        |
| kubernetes_state.resourcequota.persistentvolumeclaims.used    | Observed number of persistent volume claims used for a resource quota        |
| kubernetes_state.resourcequota.services.nodeports.used        | Observed number of node ports used for a resource quota                      |
| kubernetes_state.resourcequota.services.loadbalancers.used    | Observed number of loadbalancers used for a resource quota                   |
| kubernetes_state.resourcequota.requests.cpu.used              | Observed sum of CPU cores requested for a resource quota                     |
| kubernetes_state.resourcequota.requests.memory.used           | Observed sum of memory bytes requested for a resource quota                  |
| kubernetes_state.resourcequota.requests.storage.used          | Observed sum of storage bytes requested for a resource quota                 |
| kubernetes_state.resourcequota.limits.cpu.used                | Observed sum of limits for CPU cores for a resource quota                    |
| kubernetes_state.resourcequota.limits.memory.used             | Observed sum of limits for memory bytes for a resource quota                 |
| kubernetes_state.resourcequota.pods.limit                     | Hard limit of the number of pods for a resource quota                        |
| kubernetes_state.resourcequota.services.limit                 | Hard limit of the number of services for a resource quota                    |
| kubernetes_state.resourcequota.persistentvolumeclaims.limit   | Hard limit of the number of PVC for a resource quota                         |
| kubernetes_state.resourcequota.services.nodeports.limit       | Hard limit of the number of node ports for a resource quota                  |
| kubernetes_state.resourcequota.services.loadbalancers.limit   | Hard limit of the number of loadbalancers for a resource quota               |
| kubernetes_state.resourcequota.requests.cpu.limit             | Hard limit on the total of CPU core requested for a resource quota           |
| kubernetes_state.resourcequota.requests.memory.limit          | Hard limit on the total of memory bytes requested for a resource quota       |
| kubernetes_state.resourcequota.requests.storage.limit         | Hard limit on the total of storage bytes requested for a resource quota      |
| kubernetes_state.resourcequota.limits.cpu.limit               | Hard limit on the sum of CPU core limits for a resource quota                |
| kubernetes_state.resourcequota.limits.memory.limit            | Hard limit on the sum of memory bytes limits for a resource quota            |
| kubernetes_state.service.count                                | Sum by namespace and type to count active services                           |
| kubernetes_state.statefulset.replicas                         | The number of replicas per statefulset                                       |
| kubernetes_state.statefulset.replicas_desired                 | The number of desired replicas per statefulset                               |
| kubernetes_state.statefulset.replicas_current                 | The number of current replicas per StatefulSet                               |
| kubernetes_state.statefulset.replicas_ready                   | The number of ready replicas per StatefulSet                                 |
| kubernetes_state.statefulset.replicas_updated                 | The number of updated replicas per StatefulSet                               |
### kube-dns        
    
| key                                    | desc                                                                                      |
|----------------------------------------|-------------------------------------------------------------------------------------------|
| kubedns.response_size.bytes.sum        | Size of the returns response in bytes.                                                    |
| kubedns.response_size.bytes.count      | Number of responses on which the kubedns.response_size.bytes.sum metric is evaluated.     |
| kubedns.request_duration.seconds.sum   | Time (in seconds) each request took to resolve.                                           |
| kubedns.request_duration.seconds.count | Number of requests on which the kubedns.request_duration.seconds.sum metric is evaluated. |
| kubedns.request_count                  | Total number of DNS requests made.                                                        |
| kubedns.request_count.count            | Instant number of DNS requests made.                                                      |
| kubedns.error_count                    | Number of DNS requests resulting in an error.                                             |
| kubedns.error_count.count              | Instant number of DNS requests made resulting in an error.                                |
| kubedns.cachemiss_count                | Number of DNS requests resulting in a cache miss.                                         |
| kubedns.cachemiss_count.count          | Instant number of DNS requests made resulting in a cache miss.                            |


###  Kubernetes Proxy
| key                                | desc                                                               |
|------------------------------------|--------------------------------------------------------------------|
| kubeproxy.cpu.time                 | Total user and system CPU time spent in seconds                    |
| kubeproxy.mem.resident             | Resident memory size in bytes                                      |
| kubeproxy.mem.virtual              | Virtual memory size in bytes                                       |
| kubeproxy.client.http.requests     | Number of HTTP requests partitioned by status code method and host |
| kubeproxy.sync_rules.latency.count | SyncProxyRules latency count                                       |
| kubeproxy.sync_rules.latency.sum   | SyncProxyRules latency sum                                         |