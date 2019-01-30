## 


https://prometheus.io/docs/concepts/metric_types/

A histogram samples observations (usually things like request durations or response sizes) and counts them in configurable buckets. It also provides a sum of all observed values.

A histogram with a base metric name of <basename> exposes multiple time series during a scrape:

cumulative counters for the observation buckets, exposed as <basename>_bucket{le="<upper inclusive bound>"}
the total sum of all observed values, exposed as <basename>_sum
the count of events that have been observed, exposed as <basename>_count (identical to <basename>_bucket{le="+Inf"} above)
Use the histogram_quantile() function to calculate quantiles from histograms or even aggregations of histograms. A histogram is also suitable to calculate an Apdex score. When operating on buckets, remember that the histogram is cumulative. See histograms and summaries for details of histogram usage and differences to summaries.