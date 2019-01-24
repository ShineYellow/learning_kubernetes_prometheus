How to scrape all metrics from a federate endpoint?
https://stackoverflow.com/questions/39249048/how-to-scrape-all-metrics-from-a-federate-endpoint

```
Yes, you can do: match[]="{__name__=~".+"}" (note the + instead of * to not match the empty string).
```
Prometheus requires at least one matcher in a label matcher set that doesn't match everything.