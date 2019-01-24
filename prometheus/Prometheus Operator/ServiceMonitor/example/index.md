
prometheus example for metrics

## define ServiceMonitor
https://github.com/coreos/prometheus-operator/blob/master/example/user-guides/getting-started/example-app-service-monitor.yaml

## curl /metrics

 curl http://10.244.2.48:8080/metrics

 ```
# HELP codelab_api_http_requests_in_progress The current number of API HTTP requests in progress.
# TYPE codelab_api_http_requests_in_progress gauge
codelab_api_http_requests_in_progress 1
# HELP codelab_api_request_duration_seconds A histogram of the API HTTP request durations in seconds.
# TYPE codelab_api_request_duration_seconds histogram
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.0001"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.00015000000000000001"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.00022500000000000002"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.0003375"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.00050625"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.000759375"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.0011390624999999999"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.0017085937499999998"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.0025628906249999996"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.0038443359374999994"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.00576650390625"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.008649755859375"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.0129746337890625"} 1312
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.01946195068359375"} 17835
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.029192926025390625"} 18032
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.043789389038085935"} 18547
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.0656840835571289"} 19382
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.09852612533569335"} 19399
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.14778918800354002"} 19408
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.22168378200531003"} 19417
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.33252567300796504"} 19419
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.49878850951194753"} 19420
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="0.7481827642679213"} 19420
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="1.122274146401882"} 19420
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="1.683411219602823"} 19420
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="200",le="+Inf"} 19420
codelab_api_request_duration_seconds_sum{method="GET",path="/api/bar",status="200"} 341.1830826420011
codelab_api_request_duration_seconds_count{method="GET",path="/api/bar",status="200"} 19420
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.0001"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.00015000000000000001"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.00022500000000000002"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.0003375"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.00050625"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.000759375"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.0011390624999999999"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.0017085937499999998"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.0025628906249999996"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.0038443359374999994"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.00576650390625"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.008649755859375"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.0129746337890625"} 7
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.01946195068359375"} 46
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.029192926025390625"} 48
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.043789389038085935"} 64
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.0656840835571289"} 81
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.09852612533569335"} 81
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.14778918800354002"} 81
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.22168378200531003"} 81
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.33252567300796504"} 81
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.49878850951194753"} 81
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="0.7481827642679213"} 81
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="1.122274146401882"} 81
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="1.683411219602823"} 81
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/bar",status="500",le="+Inf"} 81
codelab_api_request_duration_seconds_sum{method="GET",path="/api/bar",status="500"} 2.1995930530000005
codelab_api_request_duration_seconds_count{method="GET",path="/api/bar",status="500"} 81
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.0001"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.00015000000000000001"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.00022500000000000002"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.0003375"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.00050625"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.000759375"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.0011390624999999999"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.0017085937499999998"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.0025628906249999996"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.0038443359374999994"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.00576650390625"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.008649755859375"} 2149
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.0129746337890625"} 32170
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.01946195068359375"} 32556
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.029192926025390625"} 33477
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.043789389038085935"} 34818
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.0656840835571289"} 34837
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.09852612533569335"} 34859
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.14778918800354002"} 34865
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.22168378200531003"} 34874
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.33252567300796504"} 34878
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.49878850951194753"} 34880
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="0.7481827642679213"} 34880
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="1.122274146401882"} 34880
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="1.683411219602823"} 34880
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="200",le="+Inf"} 34880
codelab_api_request_duration_seconds_sum{method="GET",path="/api/foo",status="200"} 407.3304108320013
codelab_api_request_duration_seconds_count{method="GET",path="/api/foo",status="200"} 34880
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.0001"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.00015000000000000001"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.00022500000000000002"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.0003375"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.00050625"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.000759375"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.0011390624999999999"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.0017085937499999998"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.0025628906249999996"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.0038443359374999994"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.00576650390625"} 0
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.008649755859375"} 6
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.0129746337890625"} 154
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.01946195068359375"} 154
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.029192926025390625"} 198
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.043789389038085935"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.0656840835571289"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.09852612533569335"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.14778918800354002"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.22168378200531003"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.33252567300796504"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.49878850951194753"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="0.7481827642679213"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="1.122274146401882"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="1.683411219602823"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/foo",status="500",le="+Inf"} 255
codelab_api_request_duration_seconds_sum{method="GET",path="/api/foo",status="500"} 4.604977193000002
codelab_api_request_duration_seconds_count{method="GET",path="/api/foo",status="500"} 255
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.0001"} 1191
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.00015000000000000001"} 1191
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.00022500000000000002"} 1191
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.0003375"} 1191
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.00050625"} 1191
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.000759375"} 1191
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.0011390624999999999"} 1191
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.0017085937499999998"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.0025628906249999996"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.0038443359374999994"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.00576650390625"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.008649755859375"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.0129746337890625"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.01946195068359375"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.029192926025390625"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.043789389038085935"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.0656840835571289"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.09852612533569335"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.14778918800354002"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.22168378200531003"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.33252567300796504"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.49878850951194753"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="0.7481827642679213"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="1.122274146401882"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="1.683411219602823"} 1192
codelab_api_request_duration_seconds_bucket{method="GET",path="/api/nonexistent",status="404",le="+Inf"} 1192
codelab_api_request_duration_seconds_sum{method="GET",path="/api/nonexistent",status="404"} 0.030786657
codelab_api_request_duration_seconds_count{method="GET",path="/api/nonexistent",status="404"} 1192
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.0001"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.00015000000000000001"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.00022500000000000002"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.0003375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.00050625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.000759375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.0011390624999999999"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.0017085937499999998"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.0025628906249999996"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.0038443359374999994"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.00576650390625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.008649755859375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.0129746337890625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.01946195068359375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.029192926025390625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.043789389038085935"} 387
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.0656840835571289"} 3937
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.09852612533569335"} 3956
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.14778918800354002"} 4109
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.22168378200531003"} 4323
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.33252567300796504"} 4324
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.49878850951194753"} 4325
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="0.7481827642679213"} 4325
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="1.122274146401882"} 4325
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="1.683411219602823"} 4325
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="200",le="+Inf"} 4325
codelab_api_request_duration_seconds_sum{method="POST",path="/api/bar",status="200"} 256.15550441500017
codelab_api_request_duration_seconds_count{method="POST",path="/api/bar",status="200"} 4325
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.0001"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.00015000000000000001"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.00022500000000000002"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.0003375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.00050625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.000759375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.0011390624999999999"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.0017085937499999998"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.0025628906249999996"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.0038443359374999994"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.00576650390625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.008649755859375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.0129746337890625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.01946195068359375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.029192926025390625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.043789389038085935"} 5
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.0656840835571289"} 44
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.09852612533569335"} 44
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.14778918800354002"} 60
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.22168378200531003"} 76
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.33252567300796504"} 76
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.49878850951194753"} 76
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="0.7481827642679213"} 76
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="1.122274146401882"} 76
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="1.683411219602823"} 76
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/bar",status="500",le="+Inf"} 76
codelab_api_request_duration_seconds_sum{method="POST",path="/api/bar",status="500"} 7.055961711999999
codelab_api_request_duration_seconds_count{method="POST",path="/api/bar",status="500"} 76
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.0001"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.00015000000000000001"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.00022500000000000002"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.0003375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.00050625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.000759375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.0011390624999999999"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.0017085937499999998"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.0025628906249999996"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.0038443359374999994"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.00576650390625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.008649755859375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.0129746337890625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.01946195068359375"} 1106
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.029192926025390625"} 3286
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.043789389038085935"} 3302
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.0656840835571289"} 3463
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.09852612533569335"} 3513
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.14778918800354002"} 3514
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.22168378200531003"} 3514
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.33252567300796504"} 3514
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.49878850951194753"} 3514
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="0.7481827642679213"} 3514
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="1.122274146401882"} 3515
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="1.683411219602823"} 3515
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="200",le="+Inf"} 3515
codelab_api_request_duration_seconds_sum{method="POST",path="/api/foo",status="200"} 80.9558719219997
codelab_api_request_duration_seconds_count{method="POST",path="/api/foo",status="200"} 3515
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.0001"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.00015000000000000001"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.00022500000000000002"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.0003375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.00050625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.000759375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.0011390624999999999"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.0017085937499999998"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.0025628906249999996"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.0038443359374999994"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.00576650390625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.008649755859375"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.0129746337890625"} 0
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.01946195068359375"} 13
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.029192926025390625"} 48
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.043789389038085935"} 48
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.0656840835571289"} 103
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.09852612533569335"} 116
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.14778918800354002"} 116
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.22168378200531003"} 116
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.33252567300796504"} 116
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.49878850951194753"} 117
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="0.7481827642679213"} 117
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="1.122274146401882"} 117
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="1.683411219602823"} 117
codelab_api_request_duration_seconds_bucket{method="POST",path="/api/foo",status="500",le="+Inf"} 117
codelab_api_request_duration_seconds_sum{method="POST",path="/api/foo",status="500"} 5.454652688999998
codelab_api_request_duration_seconds_count{method="POST",path="/api/foo",status="500"} 117
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0.000136789
go_gc_duration_seconds{quantile="0.25"} 0.00015635600000000002
go_gc_duration_seconds{quantile="0.5"} 0.000164284
go_gc_duration_seconds{quantile="0.75"} 0.000178393
go_gc_duration_seconds{quantile="1"} 0.0006090560000000001
go_gc_duration_seconds_sum 0.038570464000000006
go_gc_duration_seconds_count 217
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 28
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 4.064144e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 6.16456512e+08
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.479423e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 5.657144e+06
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 239360
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 4.064144e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 2.4576e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 4.390912e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 31430
# HELP go_memstats_heap_released_bytes_total Total number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes_total counter
go_memstats_heap_released_bytes_total 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.848512e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.5477051282917853e+19
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 29011
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 5.688574e+06
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 1208
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 50848
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 65536
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 466681
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 491520
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 491520
# HELP go_memstats_sys_bytes Number of bytes obtained by system. Sum of all system allocations.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 9.607416e+06
# HELP http_request_duration_microseconds The HTTP request latencies in microseconds.
# TYPE http_request_duration_microseconds summary
http_request_duration_microseconds{handler="api",quantile="0.5"} 11659.394
http_request_duration_microseconds{handler="api",quantile="0.9"} 30918.367
http_request_duration_microseconds{handler="api",quantile="0.99"} 59512.376
http_request_duration_microseconds_sum{handler="api"} 1.105346250682003e+09
http_request_duration_microseconds_count{handler="api"} 63861
http_request_duration_microseconds{handler="prometheus",quantile="0.5"} 9303.543
http_request_duration_microseconds{handler="prometheus",quantile="0.9"} 13014.876
http_request_duration_microseconds{handler="prometheus",quantile="0.99"} 16499.223
http_request_duration_microseconds_sum{handler="prometheus"} 1.1571763920000007e+06
http_request_duration_microseconds_count{handler="prometheus"} 108
# HELP http_request_size_bytes The HTTP request sizes in bytes.
# TYPE http_request_size_bytes summary
http_request_size_bytes{handler="api",quantile="0.5"} 71
http_request_size_bytes{handler="api",quantile="0.9"} 109
http_request_size_bytes{handler="api",quantile="0.99"} 109
http_request_size_bytes_sum{handler="api"} 4.848921e+06
http_request_size_bytes_count{handler="api"} 63861
http_request_size_bytes{handler="prometheus",quantile="0.5"} 214
http_request_size_bytes{handler="prometheus",quantile="0.9"} 214
http_request_size_bytes{handler="prometheus",quantile="0.99"} 214
http_request_size_bytes_sum{handler="prometheus"} 21456
http_request_size_bytes_count{handler="prometheus"} 108
# HELP http_requests_total Total number of HTTP requests made.
# TYPE http_requests_total counter
http_requests_total{code="0",handler="api",method="get"} 54300
http_requests_total{code="0",handler="api",method="post"} 7840
http_requests_total{code="200",handler="prometheus",method="get"} 108
http_requests_total{code="404",handler="api",method="get"} 1192
http_requests_total{code="500",handler="api",method="get"} 336
http_requests_total{code="500",handler="api",method="post"} 193
# HELP http_response_size_bytes The HTTP response sizes in bytes.
# TYPE http_response_size_bytes summary
http_response_size_bytes{handler="api",quantile="0.5"} 0
http_response_size_bytes{handler="api",quantile="0.9"} 0
http_response_size_bytes{handler="api",quantile="0.99"} 35
http_response_size_bytes_sum{handler="api"} 58119
http_response_size_bytes_count{handler="api"} 63861
http_response_size_bytes{handler="prometheus",quantile="0.5"} 3403
http_response_size_bytes{handler="prometheus",quantile="0.9"} 3425
http_response_size_bytes{handler="prometheus",quantile="0.99"} 3430
http_response_size_bytes_sum{handler="prometheus"} 364118
http_response_size_bytes_count{handler="prometheus"} 108
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 28.98
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1.048576e+06
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 14
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.4647296e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.54770402492e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.40926976e+08
```