# how to update the alertmanager config

1. create the config

for example

```yaml
config:
  global:
    resolve_timeout: 5m
  route:
    group_by: ['job']
    group_wait: 30s
    group_interval: 5m
    repeat_interval: 12h
    receiver: 'null'
    routes:
    - match:
        alertname: DeadMansSwitch
      receiver: 'null'
  receivers:
  - name: 'lianghuang@ebay.com'
```

2. encode it with base64

for example

```shell
CAT <<EOF |  base64
config:
  global:
    resolve_timeout: 5m
  route:
    group_by: ['job']
    group_wait: 30s
    group_interval: 5m
    repeat_interval: 12h
    receiver: 'null'
    routes:
    - match:
        alertname: DeadMansSwitch
      receiver: 'null'
  receivers:
  - name: 'lianghuang@ebay.com'
EOF
```

Y29uZmlnOgogIGdsb2JhbDoKICAgIHJlc29sdmVfdGltZW91dDogNW0KICByb3V0ZToKICAgIGdyb3VwX2J5OiBbJ2pvYiddCiAgICBncm91cF93YWl0OiAzMHMKICAgIGdyb3VwX2ludGVydmFsOiA1bQogICAgcmVwZWF0X2ludGVydmFsOiAxMmgKICAgIHJlY2VpdmVyOiAnbnVsbCcKICAgIHJvdXRlczoKICAgIC0gbWF0Y2g6CiAgICAgICAgYWxlcnRuYW1lOiBEZWFkTWFuc1N3aXRjaAogICAgICByZWNlaXZlcjogJ251bGwnCiAgcmVjZWl2ZXJzOgogIC0gbmFtZTogJ2xpYW5naHVhbmdAZWJheS5jb20nCg== 



3. replace the value of the key alertmanager.yaml in the file `alertmanager-secret.yaml`

