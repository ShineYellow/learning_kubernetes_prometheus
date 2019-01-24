kubernetes/kube-state-metrics
https://github.com/kubernetes/kube-state-metrics


Overview

Build Status  Go Report Card

kube-state-metrics is a simple service that listens to the Kubernetes API server and generates metrics about the state of the objects. (See examples in the Metrics section below.) It is not focused on the health of the individual Kubernetes components, but rather on the health of the various objects inside, such as deployments, nodes and pods.

That kube-state-metrics is about generating metrics from Kubernetes API objects without modification. This ensures, that features provided by kube-state-metrics have the same grade of stability as the Kubernetes API objects themselves. In turn this means, that kube-state-metrics in certain situation may not show the exact same values as kubectl, as kubectl applies certain heuristics to display comprehensible messages. kube-state-metrics exposes raw data unmodified from the Kubernetes API, this way users have all the data they require and perform heuristics as they see fit.

The metrics are exported on the HTTP endpoint /metrics on the listening port (default 80). They are served as plaintext. They are designed to be consumed either by Prometheus itself or by a scraper that is compatible with scraping a Prometheus client endpoint. You can also open /metrics in a browser to see the raw metric
