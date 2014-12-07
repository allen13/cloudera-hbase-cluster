package main

import (
	"flag"
	"fmt"
	"github.com/influxdb/influxdb/client"
	"math/rand"
	"runtime"
	"sync"
	"net/http"
	"net"
	"time"
	"crypto/tls"
)

func main() {
	host := flag.String("host", "127.0.0.1:8086", "host:port to connect to")
	points := flag.Int("points", 1, "total data points to write to server")
	connections := flag.Int("connections", 1, "number of threads to use for execution")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	loadTest := &LoadTest{
		host:        *host,
		database:    "server_metrics",
		points:      *points,
		connections: *connections,
		seriesStore:      make(chan *client.Series)}

	loadTest.Start()

}

type LoadTest struct {
	host        string
	database    string
	points      int
	connections int
	seriesStore      chan *client.Series
	waitGroup   sync.WaitGroup
}

func (self *LoadTest) Start() {
	self.waitGroup.Add(self.points)

	self.startDataPointWriters()
	self.submitDataPoints()

	self.waitGroup.Wait()
}

func (self *LoadTest) startDataPointWriters() {
	for i := 0; i < self.connections; i++ {
		go self.createDataPointWriter()
	}
}

func (self *LoadTest) submitDataPoints() {
	for i := 1; i <= self.points; i++ {
		self.seriesStore <- dataPoint(i)
	}
}

func (self *LoadTest) createDataPointWriter() {

	influxdbClient := self.newInfluxdbClient()

	for {
		series := <- self.seriesStore

		err := influxdbClient.WriteSeries([]*client.Series{series})

		if err != nil {
			fmt.Println(err)
		}

		self.waitGroup.Done()
	}
}

func (self *LoadTest) newInfluxdbClient() *client.Client {
	newClient, err := client.New(&client.ClientConfig{
		Host:     self.host,
		Database: self.database,
		Username: "root",
		Password: "root",
		HttpClient: NewHttpClient(),
	})
	if err != nil {
		panic(fmt.Sprintf("Error connecting to server \"%s\": %s", self.host, err))
	}
	return newClient
}

func NewHttpClient() *http.Client {
	timeout := 10 * time.Second
	return &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			ResponseHeaderTimeout: timeout,
			Dial: func(network, address string) (net.Conn, error) {
				return net.DialTimeout(network, address, timeout)
			},
		},
	}
}

func dataPoint(trackingObjectId int) *client.Series {
	series := "server_metrics"
	random_public_out := rand.Intn(100)

	return &client.Series{
		Name: series,
		Columns: []string{
			"trackingObjectId",
			"cpu0",
			"cpu1",
			"cpu2",
			"cpu3",
			"cpu4",
			"cpu5",
			"cpu6",
			"cpu7",
			"cpu8",
			"cpu9",
			"cpu10",
			"cpu11",
			"cpu12",
			"cpu13",
			"cpu14",
			"cpu15",
			"vbd_xvdb_write",
			"vbd_xvdb_read",
			"vbd_xvda_write",
			"vbd_xvda_read",
			"vbd_xvdc_write",
			"vbd_xvdc_read",
			"vbd_xvdd_write",
			"vbd_xvdd_read",
			"memory",
			"public_in",
			"public_out",
			"private_in",
			"private_out",
			"memory_internal_free",
			"memory_usage"},
		Points: [][]interface{}{
			{
				trackingObjectId,
				87.280266,
				19.74551,
				17.898554,
				41.58789,
				49.124672,
				14.738017,
				6.2651515,
				84.44474,
				72.30867,
				64.2923,
				46.528107,
				33.20553,
				80.86744,
				82.33386,
				75.422905,
				42.261715,
				95656032,
				297410498,
				906254121,
				686616193,
				200624334,
				799981580,
				85821045,
				288228070,
				2147483648,
				84635587,
				random_public_out,
				77778281,
				56218055,
				1431024020,
				716459628},
		},
	}
}


