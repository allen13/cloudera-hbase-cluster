package main

import (
	"flag"
	"fmt"
	"github.com/influxdb/influxdb/client"
)

func main() {
	host := flag.String("host", "127.0.0.1:8086", "host:port to connect to")
	database := flag.String("database", "server_metrics", "database to create")
	replicationFactor := flag.Int("replication_factor", 1, "how many servers to replicate the data across")
	split := flag.Int("split", 1, "how many pieces a shard will be split into")

	flag.Parse()

	influxdbClient := NewClient(*host)
	err := influxdbClient.CreateDatabase(*database)
	if err != nil {
		fmt.Println(err)
	}

	err = influxdbClient.CreateShardSpace(*database, createShardSpace(*database, *replicationFactor, *split))
	if err != nil {
		fmt.Println(err)
	}

}

func createShardSpace(database string, replicationFactor int, split int) *client.ShardSpace {
	return &client.ShardSpace{
		Name:              "default",
		Database:          database,
		RetentionPolicy:   "inf",
		ShardDuration:     "30d",
		ReplicationFactor: uint32(replicationFactor),
		Split:             uint32(split),
	}
}

func NewClient(host string) *client.Client {

	influxdbClient, err := client.NewClient(&client.ClientConfig{
		Host:     host,
		Username: "root",
		Password: "root",
	})

	if err != nil {
		fmt.Println(err)
	}
	return influxdbClient
}
