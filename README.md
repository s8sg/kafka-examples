# kafka-examples


## Consumer Group

Simple example to understand kafka consumer group 

### Getting Started 
Build the docker images for producer and consumer 
```sh
docker-compose build 
```
Now run the kafka 
```sh
docker-compose up -d kafka
```
This kafka is configured to auto create the Topics, 
although to set the custom partition number we wil manually create it


Create topic `my_topic` with **3** partitions
```sh
docker exec kafka_kafka_1 /opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh \
   --create --zookeeper localhost:2181 --replication-factor 1 --partitions 3 \
   --topic my_topic
```

#### Run Consumer 
```sh
docker-compose run kafka-consumer
```
This will spin off the 3 consumer  

#### Run Producer
```sh
docker-compose run kafka-producer
```
This will provice an interactive prompt to send messages 

### Configuration 
This example demonstrate how consumer-group can be used to broadcast or equally distributed topics accross multiple consumer 
```yaml
- "BROADCAST=false"
```
This `BROADCAST` flag controls how we assign the consumer group in the example 
#### `BROADCAST=false`
This assigns all the consumer in a same consumer group 
which will allocate one partiion to each consumer.

In this case each consumer will only receives message for the parition its asssigned with 
```sh
message received by 2 at topic/partition/offset my_topic/0/2: 1 = pantu

message received by 1 at topic/partition/offset my_topic/1/2: 2 = test2

message received by 2 at topic/partition/offset my_topic/0/3: 1 = jaja

message received by 3 at topic/partition/offset my_topic/2/2: 3 = test3 

message received by 2 at topic/partition/offset my_topic/0/4: 4 = test4

message received by 3 at topic/partition/offset my_topic/2/3: 6 = test5

message received by 1 at topic/partition/offset my_topic/1/3: 5 = empty

message received by 1 at topic/partition/offset my_topic/1/4: 7 = test6

message received by 2 at topic/partition/offset my_topic/0/5: 8 = test7
```

#### `BROADCAST=false`
This assignes all the consumer in seperate consumer group
which will assign all the partiion to each of the consumer.

In this case each consumer will receive all messages
```sh
message received by 3 at topic/partition/offset my_topic/0/3: 1 = jaja

message received by 1 at topic/partition/offset my_topic/1/2: 2 = test2

message received by 3 at topic/partition/offset my_topic/0/4: 4 = test4

message received by 2 at topic/partition/offset my_topic/1/2: 2 = test2

message received by 3 at topic/partition/offset my_topic/0/5: 8 = test7

message received by 2 at topic/partition/offset my_topic/1/3: 5 = empty

message received by 3 at topic/partition/offset my_topic/2/2: 3 = test3 

message received by 3 at topic/partition/offset my_topic/2/3: 6 = test5

message received by 2 at topic/partition/offset my_topic/1/4: 7 = test6

message received by 1 at topic/partition/offset my_topic/1/3: 5 = empty

message received by 3 at topic/partition/offset my_topic/1/2: 2 = test2

message received by 1 at topic/partition/offset my_topic/1/4: 7 = test6

message received by 2 at topic/partition/offset my_topic/0/3: 1 = jaja

message received by 1 at topic/partition/offset my_topic/0/3: 1 = jaja

message received by 3 at topic/partition/offset my_topic/1/3: 5 = empty

message received by 1 at topic/partition/offset my_topic/0/4: 4 = test4

message received by 3 at topic/partition/offset my_topic/1/4: 7 = test6

message received by 1 at topic/partition/offset my_topic/0/5: 8 = test7

message received by 2 at topic/partition/offset my_topic/0/4: 4 = test4

message received by 1 at topic/partition/offset my_topic/2/2: 3 = test3 

message received by 2 at topic/partition/offset my_topic/0/5: 8 = test7

message received by 1 at topic/partition/offset my_topic/2/3: 6 = test5

message received by 2 at topic/partition/offset my_topic/2/2: 3 = test3 

message received by 2 at topic/partition/offset my_topic/2/3: 6 = test5
```


