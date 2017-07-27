## Objective

Use Kafka to help handle distributed messaging, by using the Event Sourcing pattern that is inherently atomic.

Then, by using a pattern called Command-Query Responsibility Segregation (CQRS),
have a materialized view acting as the gate for data retrieval.
Finally, learn how to make our consumer redundant by using consumer group.
The whole application is delivered in Go.
