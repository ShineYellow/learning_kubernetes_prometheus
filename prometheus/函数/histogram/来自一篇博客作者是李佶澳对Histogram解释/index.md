https://www.lijiaocn.com/%E9%A1%B9%E7%9B%AE/2018/08/03/prometheus-usage.html

Histogram（直方图）

对采集的指标进行分组计数，会生成多个指标，分别带有后缀
- _bucket(仅histogram)、
- _sum、
- _count，

其中_bucket是区间内计数

`<basename>_bucket{le="<upper inclusive bound>"}`

举例

名为rpc_durations_secondshistogram生成的metrics：
```
# TYPE rpc_durations_histogram_seconds histogram
rpc_durations_histogram_seconds_bucket{le="-0.00099"} 0
rpc_durations_histogram_seconds_bucket{le="-0.00089"} 0
rpc_durations_histogram_seconds_bucket{le="-0.0007899999999999999"} 0
rpc_durations_histogram_seconds_bucket{le="-0.0006899999999999999"} 1
rpc_durations_histogram_seconds_bucket{le="-0.0005899999999999998"} 1
rpc_durations_histogram_seconds_bucket{le="-0.0004899999999999998"} 1
rpc_durations_histogram_seconds_bucket{le="-0.0003899999999999998"} 10
rpc_durations_histogram_seconds_bucket{le="-0.0002899999999999998"} 26
rpc_durations_histogram_seconds_bucket{le="-0.0001899999999999998"} 64
rpc_durations_histogram_seconds_bucket{le="-8.999999999999979e-05"} 117
rpc_durations_histogram_seconds_bucket{le="1.0000000000000216e-05"} 184
rpc_durations_histogram_seconds_bucket{le="0.00011000000000000022"} 251
rpc_durations_histogram_seconds_bucket{le="0.00021000000000000023"} 307
rpc_durations_histogram_seconds_bucket{le="0.0003100000000000002"} 335
rpc_durations_histogram_seconds_bucket{le="0.0004100000000000002"} 349
rpc_durations_histogram_seconds_bucket{le="0.0005100000000000003"} 353
rpc_durations_histogram_seconds_bucket{le="0.0006100000000000003"} 356
rpc_durations_histogram_seconds_bucket{le="0.0007100000000000003"} 357
rpc_durations_histogram_seconds_bucket{le="0.0008100000000000004"} 357
rpc_durations_histogram_seconds_bucket{le="0.0009100000000000004"} 357
rpc_durations_histogram_seconds_bucket{le="+Inf"} 357
rpc_durations_histogram_seconds_sum -0.000331219501489902
rpc_durations_histogram_seconds_count 357
```
Summary同样产生多个指标，分别带有后缀_bucket(仅histogram)、_sum、_count，可以直接查询分位数：

`<basename>{quantile="<φ>"}`
名为rpc_durations_secondssummary生成到metrics：
```
# TYPE rpc_durations_seconds summary
rpc_durations_seconds{service="exponential",quantile="0.5"} 7.380919552318622e-07
rpc_durations_seconds{service="exponential",quantile="0.9"} 2.291519677915514e-06
rpc_durations_seconds{service="exponential",quantile="0.99"} 4.539723552933882e-06
rpc_durations_seconds_sum{service="exponential"} 0.0005097984764772547
rpc_durations_seconds_count{service="exponential"} 532
```

Histogram和Summary都可以获取分位数。

通过Histogram获得分位数，要将直方图指标数据收集prometheus中， 然后用prometheus的查询函数histogram_quantile()计算出来。 Summary则是在应用程序中直接计算出了分位数。

Histograms and summaries中阐述了两者的区别，特别是Summary的的分位数不能被聚合。

注意，这个不能聚合不是说功能上不支持，而是说对分位数做聚合操作通常是没有意义的。

LatencyTipOfTheDay: You can’t average percentiles. Period中对“分位数”不能被相加平均的做了很详细的说明：分位数本身是用来切分数据的，它们的平均数没有同样的分位效果。