Recording rules

Recording rules allow you to precompute frequently needed or computationally expensive expressions and save their result as a new set of time series. Querying the precomputed result will then often be much faster than executing the original expression every time it is needed. This is especially useful for dashboards, which need to query the same expression repeatedly every time they refresh.

ALERTING RULES

Defining alerting rules
Inspecting alerts during runtime
Sending alert notifications
Alerting rules allow you to define alert conditions based on Prometheus expression language expressions and to send notifications about firing alerts to an external service


https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/
https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/
