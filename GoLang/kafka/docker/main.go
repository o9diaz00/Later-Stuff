package main

import (
    "fmt"
    "context"
    "os"
    "os/exec"

    "github.com/segmentio/kafka-go"
)

var c_white string = "\033[0m";
var c_cyan string = "\033[036m";
var c_red string = "\033[31m";
var c_yellow string = "\033[33m";

func connect(topic string, partition int) (*kafka.Conn, error) {
    conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", topic, partition);
    if (err != nil){
        fmt.Printf("%s", err);
    }
    return conn, err;
}

func readMessages(topic string, count string) {
    conn, err := connect(topic, 0);
    if (err != nil) {
        fmt.Printf("%s", err);
    }
    batch := conn.ReadBatch(10e1, 10e3) // get the batch

    msg:= make([]byte, 10e3) // set the max length of each message
    fmt.Printf("%s--------\n%s\n--------\n%s", c_cyan, topic, c_yellow);

    file, err := os.Create("/tmp/messages.txt");
    if (err != nil) {
        fmt.Printf("%s", err);
    }
    defer file.Close();

    for {
        msgSize, err := batch.Read(msg); // read the next message
        if err != nil {
            break
        }
        _,err = file.WriteString(string(msg[:msgSize])+"\n");
        if (err != nil) {
            fmt.Printf("%s", err);
        }
    }
    result, err := exec.Command("tail", count, "/tmp/messages.txt").Output();
    fmt.Println(string(result));
    err = os.Remove("/tmp/messages.txt");
    fmt.Printf("%s", c_white);

    if err := batch.Close(); err != nil {
        fmt.Println("failed to close batch:", err);
    }
}

func listTopics() {
    conn, err := kafka.Dial("tcp", "localhost:29092")
    if (err != nil) {
        panic(err.Error());
    }
    defer conn.Close();

    partitions, err := conn.ReadPartitions();
    if (err != nil) {
        panic(err.Error());
    }

    fmt.Printf("%s--------\nTopics\n--------\n%s", c_cyan, c_yellow);
    for _, p := range partitions {
        fmt.Printf("%#v (P%#v)\n", p.Topic, p.ID);
    }
    fmt.Printf("\n%s", c_white);
}

func main() {
    listTopics();
    readMessages("Months", "-3");
}