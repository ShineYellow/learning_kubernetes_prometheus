```yaml
env:
- name: GF_INSTALL_PLUGINS
    value: "grafana-kubernetes-app,grafana-clock-panel,grafana-simple-json-datasource"  
- name: GF_AUTH_BASIC_ENABLED
    value: "true"
- name: GF_AUTH_ANONYMOUS_ENABLED
    value: "true"
- name: GF_SECURITY_ADMIN_USER
    valueFrom:
    secretKeyRef:
        name: grafana
        key: user
- name: GF_SECURITY_ADMIN_PASSWORD
    valueFrom:
    secretKeyRef:
        name: grafana
        key: password    
```