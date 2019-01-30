irate

irate is like  rate in that it returns the per-second rate at which a counter is increas‐
ing. The algorithm it uses is much simpler though; it only looks at the last two sam‐
ples of the range vector it is passed. This has the advantage that it is much more
responsive to changes and you don’t have to care so much about the relationship
between the vector’s range and the scrape interval, but comes with the corresponding
disadvantage that as it is only looking at two samples, it can only be safely used in
graphs that are fully zoomed in. 8 Figure 16-1 shows a comparison of a 5-minute  rate
against an  irate