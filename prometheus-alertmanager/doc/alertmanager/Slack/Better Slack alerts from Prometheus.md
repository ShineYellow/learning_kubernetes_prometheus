Better Slack alerts from Prometheus

https://medium.com/quiq-blog/better-slack-alerts-from-prometheus-49125c8c672b


Set title and text fields of slack_config of your Alertmanager’s config to use custom_title and custom_slack_message definitions as follow:
```
receivers:
- name: slack-channel
  slack_configs:
  - channel: #monitoring
    icon_url: https://avatars3.githubusercontent.com/u/3380462
    send_resolved: true
    title: '{{ template "custom_title" . }}'
    text: '{{ template "custom_slack_message" . }}'
Also add the following snippet in the end assuming Alertmanger can access this file:

templates:
- /alertmanager/notifications.tmpl
```
2. Create notifications.tmpl:
```
{{ define "__single_message_title" }}{{ range .Alerts.Firing }}{{ .Labels.alertname }} @ {{ .Annotations.identifier }}{{ end }}{{ range .Alerts.Resolved }}{{ .Labels.alertname }} @ {{ .Annotations.identifier }}{{ end }}{{ end }}
{{ define "custom_title" }}[{{ .Status | toUpper }}{{ if eq .Status "firing" }}:{{ .Alerts.Firing | len }}{{ end }}] {{ if or (and (eq (len .Alerts.Firing) 1) (eq (len .Alerts.Resolved) 0)) (and (eq (len .Alerts.Firing) 0) (eq (len .Alerts.Resolved) 1)) }}{{ template "__single_message_title" . }}{{ end }}{{ end }}
{{ define "custom_slack_message" }}
{{ if or (and (eq (len .Alerts.Firing) 1) (eq (len .Alerts.Resolved) 0)) (and (eq (len .Alerts.Firing) 0) (eq (len .Alerts.Resolved) 1)) }}
{{ range .Alerts.Firing }}{{ .Annotations.description }}{{ end }}{{ range .Alerts.Resolved }}{{ .Annotations.description }}{{ end }}
{{ else }}
{{ if gt (len .Alerts.Firing) 0 }}
*Alerts Firing:*
{{ range .Alerts.Firing }}- {{ .Annotations.identifier }}: {{ .Annotations.description }}
{{ end }}{{ end }}
{{ if gt (len .Alerts.Resolved) 0 }}
*Alerts Resolved:*
{{ range .Alerts.Resolved }}- {{ .Annotations.identifier }}: {{ .Annotations.description }}
{{ end }}{{ end }}
{{ end }}
{{ end }}
```
3. Ensure each of your Prometheus alerts contains identifier and description tags in the annotations block, e.g.
```
- alert: instance_down
  expr: up == 0
  for: 5m
  annotations:
    identifier: '{{ $labels.instance }}'
    description: '{{ $labels.job }} exporter job has been down for more than 5 minutes.'
```