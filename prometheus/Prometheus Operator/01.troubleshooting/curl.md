
## get the serviceaccount token from prometheus pods
- kubectl exec prometheus-k8s-0  -n monitoring  -- cat /var/run/secrets/kubernetes.io/serviceaccount/token
- kubectl exec prometheus-k8s-1  -n monitoring  -- cat /var/run/secrets/kubernetes.io/serviceaccount/token

## make a curl to kubelet to get the metrics

```
curl -kvi -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJtb25pdG9yaW5nIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InByb21ldGhldXMtazhzLXRva2VuLXNqNXQ1Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InByb21ldGhldXMtazhzIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiMTFkNmZiMTItZjIxNi0xMWU4LTkxNjYtMDA1MDU2YjJmMjBjIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Om1vbml0b3Jpbmc6cHJvbWV0aGV1cy1rOHMifQ.eidlanB5ZuUcRJCIPmDGsgq5_6Da0-nIw0Dw7h4NZhrupT1yvYS5CFbB2azcAJnm3Vq0qnNiw7_wT1nKgP06WvgCQvihgoMYTLXMgoXF5w7pUGJvC3mbHDRUa5FjZRx99WZOAzSOJc5lrxEnMFDRq3yrM1mSz4QdbFL7yTnPRAmJ_o1A26DTNcoPelxnLqfJ7mN-s49YRel6Eo81uLIZwm5ju44zy_rj7rJx32dbg5zkSkf7GzM9VXrvow4qm-B5hBJ0lW78Op6K0YSqVS_gcHW1-LhsneodYJJnVNe-Y9e7yIJS6yqj5IZEuQPhvL-jDkxOfcFX4VlSGN1zrTgGBQ" https://10.242.105.204:9100/metrics
*   Trying 10.242.105.204...
* Connected to 10.242.105.204 (10.242.105.204) port 9100 (#0)
* found 148 certificates in /etc/ssl/certs/ca-certificates.crt
* found 596 certificates in /etc/ssl/certs
* ALPN, offering http/1.1
* SSL connection using TLS1.2 / ECDHE_RSA_AES_128_GCM_SHA256
* 	 server certificate verification SKIPPED
* 	 server certificate status verification SKIPPED
* 	 common name: @1543303754 (does not match '10.242.105.204')
* 	 server certificate expiration date OK
* 	 server certificate activation date OK
* 	 certificate public key: RSA
* 	 certificate version: #3
* 	 subject: CN=@1543303754
* 	 start date: Tue, 27 Nov 2018 07:29:14 GMT
* 	 expire date: Wed, 27 Nov 2019 07:29:14 GMT
* 	 issuer: CN=@1543303754
* 	 compression: NULL
* ALPN, server accepted to use http/1.1
> GET /metrics HTTP/1.1
> Host: 10.242.105.204:9100
> User-Agent: curl/7.47.0
> Accept: */*
> Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJtb25pdG9yaW5nIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InByb21ldGhldXMtazhzLXRva2VuLXNqNXQ1Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InByb21ldGhldXMtazhzIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiMTFkNmZiMTItZjIxNi0xMWU4LTkxNjYtMDA1MDU2YjJmMjBjIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Om1vbml0b3Jpbmc6cHJvbWV0aGV1cy1rOHMifQ.eidlanB5ZuUcRJCIPmDGsgq5_6Da0-nIw0Dw7h4NZhrupT1yvYS5CFbB2azcAJnm3Vq0qnNiw7_wT1nKgP06WvgCQvihgoMYTLXMgoXF5w7pUGJvC3mbHDRUa5FjZRx99WZOAzSOJc5lrxEnMFDRq3yrM1mSz4QdbFL7yTnPRAmJ_o1A26DTNcoPelxnLqfJ7mN-s49YRel6Eo81uLIZwm5ju44zy_rj7rJx32dbg5zkSkf7GzM9VXrvow4qm-B5hBJ0lW78Op6K0YSqVS_gcHW1-LhsneodYJJnVNe-Y9e7yIJS6yqj5IZEuQPhvL-jDkxOfcFX4VlSGN1zrTgGBQ
>
< HTTP/1.1 200 OK
HTTP/1.1 200 OK
< Content-Type: text/plain; version=0.0.4; charset=utf-8
Content-Type: text/plain; version=0.0.4; charset=utf-8
< Date: Mon, 21 Jan 2019 10:47:23 GMT
Date: Mon, 21 Jan 2019 10:47:23 GMT
< Transfer-Encoding: chunked
Transfer-Encoding: chunked

<
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 1.6446e-05
go_gc_duration_seconds{quantile="0.25"} 1.7337e-05
go_gc_duration_seconds{quantile="0.5"} 1.8007e-05
go_gc_duration_seconds{quantile="0.75"} 1.9271e-05
go_gc_duration_seconds{quantile="1"} 0.001180204
go_gc_duration_seconds_sum 15.745453771
go_gc_duration_seconds_count 317582
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 6
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.9.6"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 2.310352e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 7.56126539584e+11
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 2.232928e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 6.650014204e+09
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 0.0004826345810965034
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 405504
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 2.310352e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 2.53952e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 3.35872e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 17111
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 5.89824e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.5480676426936586e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 1.778467e+07
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 6.650031315e+09
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 1736
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 40128
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 65536
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 497304
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 393216
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 393216
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 9.509112e+06
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 7
# HELP node_arp_entries ARP entries by device
# TYPE node_arp_entries gauge
node_arp_entries{device="eth0"} 4
node_arp_entries{device="flannel.1"} 2
# HELP node_boot_time_seconds Node boot time, in unixtime.
# TYPE node_boot_time_seconds gauge
node_boot_time_seconds 1.541587864e+09
# HELP node_context_switches_total Total number of context switches.
# TYPE node_context_switches_total counter
node_context_switches_total 4.846728311e+10
# HELP node_cpu_guest_seconds_total Seconds the cpus spent in guests (VMs) for each mode.
# TYPE node_cpu_guest_seconds_total counter
node_cpu_guest_seconds_total{cpu="0",mode="nice"} 0
node_cpu_guest_seconds_total{cpu="0",mode="user"} 0
# HELP node_cpu_seconds_total Seconds the cpus spent in each mode.
# TYPE node_cpu_seconds_total counter
node_cpu_seconds_total{cpu="0",mode="idle"} 5.48760586e+06
node_cpu_seconds_total{cpu="0",mode="iowait"} 2863.37
node_cpu_seconds_total{cpu="0",mode="irq"} 0
node_cpu_seconds_total{cpu="0",mode="nice"} 1317.39
node_cpu_seconds_total{cpu="0",mode="softirq"} 10740.13
node_cpu_seconds_total{cpu="0",mode="steal"} 0
node_cpu_seconds_total{cpu="0",mode="system"} 236797.39
node_cpu_seconds_total{cpu="0",mode="user"} 624302.37
# HELP node_disk_io_now The number of I/Os currently in progress.
# TYPE node_disk_io_now gauge
node_disk_io_now{device="dm-0"} 0
node_disk_io_now{device="dm-1"} 0
node_disk_io_now{device="dm-2"} 0
node_disk_io_now{device="dm-3"} 0
node_disk_io_now{device="dm-4"} 0
node_disk_io_now{device="sda"} 0
node_disk_io_now{device="sr0"} 0
# HELP node_disk_io_time_seconds_total Total seconds spent doing I/Os.
# TYPE node_disk_io_time_seconds_total counter
node_disk_io_time_seconds_total{device="dm-0"} 58.568
node_disk_io_time_seconds_total{device="dm-1"} 0.06
node_disk_io_time_seconds_total{device="dm-2"} 11012.52
node_disk_io_time_seconds_total{device="dm-3"} 54.616
node_disk_io_time_seconds_total{device="dm-4"} 0.34800000000000003
node_disk_io_time_seconds_total{device="sda"} 10468.508
node_disk_io_time_seconds_total{device="sr0"} 0
# HELP node_disk_io_time_weighted_seconds_total The weighted # of seconds spent doing I/Os. See https://www.kernel.org/doc/Documentation/iostats.txt.
# TYPE node_disk_io_time_weighted_seconds_total counter
node_disk_io_time_weighted_seconds_total{device="dm-0"} 686.22
node_disk_io_time_weighted_seconds_total{device="dm-1"} 0.06
node_disk_io_time_weighted_seconds_total{device="dm-2"} 25517.564000000002
node_disk_io_time_weighted_seconds_total{device="dm-3"} 255.184
node_disk_io_time_weighted_seconds_total{device="dm-4"} 0.368
node_disk_io_time_weighted_seconds_total{device="sda"} 22119.128
node_disk_io_time_weighted_seconds_total{device="sr0"} 0
# HELP node_disk_read_bytes_total The total number of bytes read successfully.
# TYPE node_disk_read_bytes_total counter
node_disk_read_bytes_total{device="dm-0"} 2.10732032e+08
node_disk_read_bytes_total{device="dm-1"} 3.33824e+06
node_disk_read_bytes_total{device="dm-2"} 3.9428608e+08
node_disk_read_bytes_total{device="dm-3"} 2.66752e+06
node_disk_read_bytes_total{device="dm-4"} 2.405376e+06
node_disk_read_bytes_total{device="sda"} 6.2040064e+08
node_disk_read_bytes_total{device="sr0"} 0
# HELP node_disk_read_time_seconds_total The total number of milliseconds spent by all reads.
# TYPE node_disk_read_time_seconds_total counter
node_disk_read_time_seconds_total{device="dm-0"} 55.480000000000004
node_disk_read_time_seconds_total{device="dm-1"} 0.06
node_disk_read_time_seconds_total{device="dm-2"} 30.028000000000002
node_disk_read_time_seconds_total{device="dm-3"} 0.14
node_disk_read_time_seconds_total{device="dm-4"} 0.152
node_disk_read_time_seconds_total{device="sda"} 84.956
node_disk_read_time_seconds_total{device="sr0"} 0
# HELP node_disk_reads_completed_total The total number of reads completed successfully.
# TYPE node_disk_reads_completed_total counter
node_disk_reads_completed_total{device="dm-0"} 17537
node_disk_reads_completed_total{device="dm-1"} 136
node_disk_reads_completed_total{device="dm-2"} 6478
node_disk_reads_completed_total{device="dm-3"} 210
node_disk_reads_completed_total{device="dm-4"} 109
node_disk_reads_completed_total{device="sda"} 24696
node_disk_reads_completed_total{device="sr0"} 0
# HELP node_disk_reads_merged_total The total number of reads merged. See https://www.kernel.org/doc/Documentation/iostats.txt.
# TYPE node_disk_reads_merged_total counter
node_disk_reads_merged_total{device="dm-0"} 0
node_disk_reads_merged_total{device="dm-1"} 0
node_disk_reads_merged_total{device="dm-2"} 0
node_disk_reads_merged_total{device="dm-3"} 0
node_disk_reads_merged_total{device="dm-4"} 0
node_disk_reads_merged_total{device="sda"} 94
node_disk_reads_merged_total{device="sr0"} 0
# HELP node_disk_write_time_seconds_total This is the total number of seconds spent by all writes.
# TYPE node_disk_write_time_seconds_total counter
node_disk_write_time_seconds_total{device="dm-0"} 630.72
node_disk_write_time_seconds_total{device="dm-1"} 0
node_disk_write_time_seconds_total{device="dm-2"} 25309.036
node_disk_write_time_seconds_total{device="dm-3"} 255.04
node_disk_write_time_seconds_total{device="dm-4"} 0.216
node_disk_write_time_seconds_total{device="sda"} 22133.496
node_disk_write_time_seconds_total{device="sr0"} 0
# HELP node_disk_writes_completed_total The total number of writes completed successfully.
# TYPE node_disk_writes_completed_total counter
node_disk_writes_completed_total{device="dm-0"} 160135
node_disk_writes_completed_total{device="dm-1"} 0
node_disk_writes_completed_total{device="dm-2"} 7.467788e+07
node_disk_writes_completed_total{device="dm-3"} 5.345373e+06
node_disk_writes_completed_total{device="dm-4"} 208
node_disk_writes_completed_total{device="sda"} 6.7211544e+07
node_disk_writes_completed_total{device="sr0"} 0
# HELP node_disk_writes_merged_total The number of writes merged. See https://www.kernel.org/doc/Documentation/iostats.txt.
# TYPE node_disk_writes_merged_total counter
node_disk_writes_merged_total{device="dm-0"} 0
node_disk_writes_merged_total{device="dm-1"} 0
node_disk_writes_merged_total{device="dm-2"} 0
node_disk_writes_merged_total{device="dm-3"} 0
node_disk_writes_merged_total{device="dm-4"} 0
node_disk_writes_merged_total{device="sda"} 1.2980174e+07
node_disk_writes_merged_total{device="sr0"} 0
# HELP node_disk_written_bytes_total The total number of bytes written successfully.
# TYPE node_disk_written_bytes_total counter
node_disk_written_bytes_total{device="dm-0"} 1.923858432e+09
node_disk_written_bytes_total{device="dm-1"} 0
node_disk_written_bytes_total{device="dm-2"} 4.9628891136e+11
node_disk_written_bytes_total{device="dm-3"} 2.1905743872e+10
node_disk_written_bytes_total{device="dm-4"} 851968
node_disk_written_bytes_total{device="sda"} 5.20119391232e+11
node_disk_written_bytes_total{device="sr0"} 0
# HELP node_entropy_available_bits Bits of available entropy.
# TYPE node_entropy_available_bits gauge
node_entropy_available_bits 841
# HELP node_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which node_exporter was built.
# TYPE node_exporter_build_info gauge
node_exporter_build_info{branch="HEAD",goversion="go1.9.6",revision="d42bd70f4363dced6b77d8fc311ea57b63387e4f",version="0.16.0"} 1
# HELP node_filefd_allocated File descriptor statistics: allocated.
# TYPE node_filefd_allocated gauge
node_filefd_allocated 1600
# HELP node_filefd_maximum File descriptor statistics: maximum.
# TYPE node_filefd_maximum gauge
node_filefd_maximum 399213
# HELP node_filesystem_avail_bytes Filesystem space available to non-root users in bytes.
# TYPE node_filesystem_avail_bytes gauge
node_filesystem_avail_bytes{device="/dev/mapper/rootvg-home",fstype="ext4",mountpoint="/host/root/home"} 9.39188224e+09
node_filesystem_avail_bytes{device="/dev/mapper/rootvg-root",fstype="ext4",mountpoint="/host/root"} 7.075168256e+09
node_filesystem_avail_bytes{device="/dev/mapper/rootvg-tmp",fstype="ext4",mountpoint="/host/root/tmp"} 9.39194368e+09
node_filesystem_avail_bytes{device="/dev/mapper/rootvg-var",fstype="ext4",mountpoint="/etc/hostname"} 6.692274176e+09
node_filesystem_avail_bytes{device="/dev/mapper/rootvg-var",fstype="ext4",mountpoint="/etc/hosts"} 6.692274176e+09
node_filesystem_avail_bytes{device="/dev/mapper/rootvg-var",fstype="ext4",mountpoint="/etc/resolv.conf"} 6.692274176e+09
node_filesystem_avail_bytes{device="/dev/mapper/rootvg-var",fstype="ext4
```