# manual alert trigger to debug/test alert manager congfiguration



Here's a sample curl command to trigger an alert.

You can see the full tags of a message from prometheus by stopping the alertmanager and running "nc -k -l 9093" in it's place to grab incoming messages.

`curl -H "Content-Type: application/json" -d '[{"labels":{"alertname":"TestAlert1"}}]' localhost:9093/api/v1/alerts`


## reference
https://github.com/prometheus/alertmanager/issues/437