## quantile
The  quantile aggregator returns the specified quantile of the values of the group as
the group’s return value. As with  topk ,  quantile takes a parameter.
So, 

for example, if I wanted to know across the different CPUs in each of my
machines what the 90 th percentile of the system mode CPU usage is I could use:
quantile without(cpu)(0.9, rate(node_cpu_seconds_total{mode="system"}[5m]))
which produces a result like:
{instance="localhost:9100",job="node",mode="system"} 0.024558620689654007
This means that 90% of my CPUs are spending at least 0.02 seconds per second in the
system mode. 

This would be a more useful query if I had tens of CPUs in my
machine, rather than the four it actually has.
In addition to the mean, you could use  quantile to show the median, 25 th , and 75 th
percentiles 12 on your graphs. For example, for process CPU usage the expressions
would be:
```
# average, arithmetic mean
avg without(instance)(0.5, rate(process_cpu_seconds_total[5m]))
# 0.25 quantile, 25th percentile, 1st or lower quartile
quantile without(instance)(0.25, rate(process_cpu_seconds_total[5m]))
# 0.5 quantile, 50th percentile, 2nd quartile, median
quantile without(instance)(0.5, rate(process_cpu_seconds_total[5m]))
# 0.75 quantile, 75th percentile, 3rd or upper quartile
quantile without(instance)(0.75, rate(process_cpu_seconds_total[5m]))
```
This would give you a sense of how your different instances for a job are behaving,
without having to graph each instance individually. This allows you to keep your
dashboards readable as the number of underlying instances grows. Personally I find
that per-instance graphs break down somewhere around three to five instances.



## quantile, histogram_quantile, and quantile_over_time
As you may have noticed by now, there is more than one PromQL function or opera‐
tor with quantile in the name.
The  quantile aggregator works across an instant vector in an aggregation group.
The  quantile_over_time function works across a single time series at a time in a
range vector.
The  histogram_quantile function works across the buckets of one histogram metric
child at a time in an instant vector.
