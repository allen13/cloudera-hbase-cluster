package main

import (
	"flag"
	"fmt"
	"github.com/influxdb/influxdb/client"
	"math/rand"
	"sync"
)

func main() {
	host := flag.String("host", "127.0.0.1:8086", "host:port to connect to")
	points := flag.Int("points", 1, "total data points to write to server")
	threads := flag.Int("threads", 1, "number of threads to use for execution")
	flag.Parse()

	loadTest := NewLoadTest(*host, "server_metrics")
	loadTest.pageQueries(*points, *threads)
}

type LoadTest struct {
	client *client.Client
}

func NewLoadTest(host string, database string) *LoadTest {
	loadTest := new(LoadTest)

	loadTest.client, _ = client.NewClient(&client.ClientConfig{
		Host:     host,
		Username: "root",
		Password: "root",
		Database: database,
	})

	return loadTest
}
func (self *LoadTest) pageQueries(queries int, threadCount int) {
	var waitGroup sync.WaitGroup
	pageSize := queries / threadCount

	for thread := 0; thread < threadCount; thread++ {
		waitGroup.Add(1)
		go self.runPage(thread,pageSize, &waitGroup)
	}

	waitGroup.Wait()
}

func (self *LoadTest) runPage(thread int, pageSize int, waitGroup *sync.WaitGroup) {
	for threadQuery := 1; threadQuery <= pageSize; threadQuery++ {
		pageStart := thread * pageSize
		queryId := pageStart + threadQuery
		self.writeDataPoint(queryId)
	}
	waitGroup.Done()
}

func (self *LoadTest) writeDataPoint(trackingObjectId int) {
	random_public_out := rand.Intn(100)
	err := self.client.WriteSeries(dataPoint(trackingObjectId, random_public_out))
	if err != nil {
		fmt.Println("Write failed.")
		panic(err)
	}
	fmt.Println("wrote trackingObjectId: ", trackingObjectId)
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
