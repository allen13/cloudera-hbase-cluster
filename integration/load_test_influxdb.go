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

type LoadTest struct {
	host        string
	database    string
	points      int
	connections int
	writes      chan *LoadWrite
	sync.WaitGroup
}

type LoadWrite struct {
	Series []*client.Series
}

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
		writes:      make(chan *LoadWrite)}

	loadTest.Start()

}

func (self *LoadTest) Start() {
	self.Add(self.points)
	self.startPostWorkers()
	self.writeDataPoints()
	self.Wait()
}

func (self *LoadTest) startPostWorkers() {
	for i := 0; i < self.connections; i++ {
		go self.handleWrites()
	}
}

func (self *LoadTest) handleWrites() {

	influxdbClient := self.newClient()

	for {
		write := <-self.writes

		err := influxdbClient.WriteSeries(write.Series)
		self.Done()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (self *LoadTest) newClient() *client.Client {
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

func (self *LoadTest) writeDataPoints() {
	for i := 1; i <= self.points; i++ {
		self.writeDataPoint(i)
	}
}

func (self *LoadTest) writeDataPoint(trackingObjectId int) {
	random_public_out := rand.Intn(100)
	self.writes <- &LoadWrite{Series: dataPoint(trackingObjectId, random_public_out)}
}

func dataPoint(trackingObjectId int, public_out int) []*client.Series {
	series := "server_metrics"

	return []*client.Series{&client.Series{
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
			{trackingObjectId, 87.280266, 19.74551, 17.898554, 41.58789, 49.124672, 14.738017, 6.2651515, 84.44474, 72.30867, 64.2923, 46.528107, 33.20553, 80.86744, 82.33386, 75.422905, 42.261715, 95656032, 297410498, 906254121, 686616193, 200624334, 799981580, 85821045, 288228070, 2147483648, 84635587, public_out, 77778281, 56218055, 1431024020, 716459628},
		},
	}}
}
