Prometheus Dashboard Queries

Storage by container
  sum(container_fs_usage_bytes{image!=""}) by (image)

Running container
  count_scalar(container_last_seen{name=~"prom|cadvisor"})

Top 3 Memory
  topk(3, sum(container_memory_usage_bytes{image!=""}) by (image)) 

FS Usage
  sum(container_fs_usage_bytes{name!=""}) by (name)

Network 
	sum(container_network_receive_bytes_total{name=""}) by (name)
	sum(container_network_transmit_bytes_total{name=""}) by (name)

Source Scrape Time 
  scrape_duration_seconds