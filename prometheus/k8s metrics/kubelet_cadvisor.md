

## kubelet/cadvisor

https://docs.signalfx.com/en/latest/integrations/agent/monitors/kubelet-stats.html

source code 
https://github.com/google/cadvisor/blob/master/metrics/prometheus.go

```
metricsToInclude:
  - metricNames:
    - container_cpu_cfs_periods
    - container_cpu_cfs_throttled_periods
    - container_cpu_cfs_throttled_time
    - container_cpu_percent
    - container_cpu_system_seconds_total
    - container_cpu_usage_seconds_total
    - container_cpu_user_seconds_total
    - container_fs_io_current
    - container_fs_io_time_seconds_total
    - container_fs_io_time_weighted_seconds_total
    - container_fs_limit_bytes
    - container_fs_read_seconds_total
    - container_fs_reads_merged_total
    - container_fs_reads_total
    - container_fs_sector_reads_total
    - container_fs_sector_writes_total
    - container_fs_usage_bytes
    - container_fs_write_seconds_total
    - container_fs_writes_merged_total
    - container_fs_writes_total
    - container_last_seen
    - container_memory_failcnt
    - container_memory_working_set_bytes
    - container_spec_cpu_shares
    - container_spec_memory_swap_limit_bytes
    - container_start_time_seconds
    - container_tasks_state
    - machine_cpu_frequency_khz
    - pod_network_receive_packets_dropped_total
    - pod_network_receive_packets_total
    - pod_network_transmit_packets_dropped_total
    - pod_network_transmit_packets_total
    monitorType: kubelet-stats
```