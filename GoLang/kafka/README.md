# Kafka CLI ([REMOVED])

---
This is a work in progress / proof of concept of using GoLang to interact with kafka clusters.

<b>Note:</b> This version will leave out any company information and/or allusions to it.  With that being said, the code also has to be modified some to fulfill said criteria

---
## Getting Started

---
The compressed file should contain the GO code, as well as teh information for the Docker containers used.  Ideally, you'll just have to spin up the docker container, log into the provided Kafka UI, and add some data.  Then run the GO code (as well as edit it as you want to see changes).

---
## Prerequisites

---
You may or may not need to match the software versions, but for your safe knowledge, these are what I am currently running:
```
$ docker --version
Docker version 27.5.1, build 9f9e405

docker-compose --version
docker-compose version 1.29.2, build unknown
```
---
## Setting Up

---
**<u>Spinning Up Docker Container</u>**
- From the terminal, navigate to the corresponding "kafka/docker" directory
- Bring up the container with the command `docker-compose up -d [--build]`
  - note: build may or may not be necessary

**<u>Adding Data to Kafka</u>**
- Log into the kafka UI: https://localhost:8080
  - Since the port 8080 is used on the container, the above URL should log you into it
- Navigate to "Topics" in side bar and Select "+ Add a Topic" in the top right to create a topic
- Produce messages for each topic(s)
  - Topics -> $TOPIC NAME -> "Produce Message" in top right corner

## Running Tests

---
After the containers have been spun up and data added to kafka (Topics and Messages), navigate to the GO code (main.go).

Navigate to the bottom within the function "main" and change the arguments of "readMessages" to match your configurations:

- The first argument should match a topic that exists in your Kafka that you want to read from
- The second argument should match the number of messages you want to read from said topic (preface the number w/ "-".  For example, if you want to read 3 messages, input "-3").

Execute the GO code by running the command `go run main.go`

## Next Steps

---
A few immediate next step implementations can be made.

Firstly, we can adjust this kafka setup so that it does indeed require authentication and adjust the GOLang code accordingly (as that will more closely mimic our production kafka environment)

Secondly, we can add consumer groups into the fold, and adjust the GOLang code to utilize the consumer groups when reading rather than read directly from the partition.  (This I am not so clear on, if it will be a necessity.  As I never had an actual kafka cluster of ours to test on, I was unable to see whether or not we could read from the partition directly and effectively, or if we'd need to utilize the consumer group[s])

---
## Misc

---
**Segment IO Kafka Development Package Documentation:** https://pkg.go.dev/github.com/segmentio/kafka-go#section-documentation
**Segment IO Kafka Github:** https://github.com/segmentio/kafka-go