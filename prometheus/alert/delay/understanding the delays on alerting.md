# Prometheus: understanding the delays on alerting

## Scraping, Evaluation and Alerting

Prometheus scrape metrics from monitored targets at regular intervals, defined by the scrape_interval (defaults to 1m).
The scrape interval can be configured globally, and then overriden per job.


Prometheus has another loop, whose clock is independent from the scraping one, that evaluates alerting rules at a regular interval, defined by `evaluation_interval` (defaults to 1m). 
At each evaluation cycle, Prometheus runs the expression defined in each alerting rule and updates the alert state.


## An alert can have the following states:

- inactive: the state of an alert that is neither firing nor pending
- pending: the state of an alert that has been active for less than the configured threshold duration
- firing: the state of an alert that has been active for longer than the configured threshold duration

An alert transitions from a state to another only during the evaluation cycle. How the transition occurs, depends if the alert’s `FOR` clause has been set or not. From the doc:

> The optional FOR clause causes Prometheus to wait for a certain duration between first encountering a new expression output vector element (like an instance with a high HTTP error rate) and counting an alert as firing for this element.

- Alerts `without` the FOR clause (or set to 0) will immediately transition to firing.
- Alerts `with` the FOR clause will transition first to pending and then firing, thus it will take at least two evaluation cycles before an alert with the FOR clause is fired.

Prometheus: understanding the delays on alerting
https://pracucci.com/prometheus-understanding-the-delays-on-alerting.html