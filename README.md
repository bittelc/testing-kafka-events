## Objective

Use Kafka to help handle distributed messaging, by using the Event Sourcing pattern that is inherently atomic.

Then, by using a pattern called Command-Query Responsibility Segregation (CQRS),
have a materialized view acting as the gate for data retrieval.
Finally, learn how to make our consumer redundant by using consumer group.
The whole application is delivered in Go.

## Microservices

Microservices allow us to decompose a large, cumbersome application into smaller, more manageable repositories. This allows greater scalability and development speed on each individual microservice.

## Challenges

Atomicity!! How do we deal with distributed data, which is inherent to microservice architecture?

## Kafka

Database :: Kafka

Table:Database :: Topic:Kafka

Row:Table :: Commit Log:Topic

Commit logs are ordered based on offset, where each new message N+1 the offset of the last commit.

## Up and Running

1. Download and install Kafka. If necessary, also download ZooKeeper independently
2. Start Zookeeper server(s)
    * `zookeeper-server-start.sh config/zookeeper.properties`
3. Start Kafka server(s)
    * `kafka-server-start.sh config/server.properties`
4.Create first topic
    * `kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic xbanku-transactions-t1`
5. 



