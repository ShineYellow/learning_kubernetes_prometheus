# Histogram

摘自prometheus up & running -- part IV Promql， Charpter13.introduction to PromQL， Aggregation Basics


Histogram metrics allow you to track the `distribution` of the size of events, allowing
you to calculate quantiles from them. 

For example, you can use histograms to calcu‐late the 0.9 quantile (which is also known as the 90 th percentile) latency.

Prometheus 2.2.1 exposes a histogram metric called  `prometheus_tsdb_compaction_duration_seconds` that tracks how many seconds compaction takes for the time series database. This histogram metric has time series with a  `_bucket` suffix called prometheus_tsdb_compaction_duration_seconds_bucket . Each bucket has a `le label`, which is a counter of how many events have a size less than or equal to the bucket boundary. 

This is an implementation detail you largely need not worry about, as the  `histogram_quantile` function takes care of this when calculating quantiles.

For example, the 0.90 quantile would be:
```
histogram_quantile(
0.90,
rate(prometheus_tsdb_compaction_duration_seconds_bucket[1d]))
```
As  prometheus_tsdb_compaction_duration_seconds_bucket is a counter you must
first take a  rate . Compaction usually only happens every two hours, so a one-day
time range is used here so you will see a result in the expression browser such as:

`{instance="localhost:9090",job="prometheus"} 7.720000000000001`

This indicates that the 90 th percentile latency of compactions is around `7.72 seconds`.
As there will usually only be 12 compactions in a day, the 90 th percentile says that
`10% of compactions take longer than this`, which is to say one or two compactions.
This is something to be aware of when using quantiles. 

For example, if you want to calculate a 0.999 quantile you should have several thousand data points to work with in order to produce a reasonably accurate answer. If you have fewer than that, single outliers could greatly affect the result, and you should consider using lower quantiles to avoid making statements about your system that you have insufficient data to back up.

## 我的理解

Histogram 是对你感兴趣的样本作分布的统计

Histogram中需要提供_bucket这个metrcis

这里通过举例普罗米修斯每次做压缩所要花费的时间分布举例，统计出，90%的压缩时间只要7.7，也就是说，10%的压缩大于了这个数

Histogram对样本数量有要求，更高的精度，需要更高的样本数量