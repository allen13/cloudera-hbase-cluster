package main

import (
	"flag"
	"github.com/influxdb/influxdb/client"
)

type LoadTest struct {
	client *client.Client
}

func main() {
	host := flag.String("host", "127.0.0.1:8086", "host:port to connect to")
	database := flag.String("database", "server_metrics", "database to create")
	flag.Parse()

	loadTest := NewLoadTest(*host)
	loadTest.createDB(*database)

}

func NewLoadTest(host string) *LoadTest {
	loadTest := new(LoadTest)

	loadTest.client, _ = client.NewClient(&client.ClientConfig{
		Host:     host,
		Username: "root",
		Password: "root",
	})

	return loadTest
}

func (self *LoadTest) createDB(database string) {
	self.client.CreateDatabase(database)
}
