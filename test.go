package main

import (
	"context"
	"fmt"
	"github.com/ThinkontrolSY/flux-builder/filter"
	pipe "github.com/ThinkontrolSY/flux-builder/transform_pipe"
	"github.com/influxdata/influxdb-client-go/v2"
)

type FluxQuery struct {
	bucket     string
	start      string
	stop       *string
	filters    []*filter.FluxFilter
	transforms []pipe.TransformPipe
}

func main() {
	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	client := influxdb2.NewClient("http://9king.club:8086", "yuEKW9Kp96EGIwVq_8bIs0eWV1sAHFFUFUV7cVdXIkQUjvKA8517bY3ujRZBv2G7IKVGTB2Kd3wQclSbPf981g==")
	// always close client at the end
	defer client.Close()
	// get non-blocking write client
	//writeAPI := client.WriteAPI("9king.club", "test1")
	//
	//// write line protocol
	//writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	//// Flush writes
	//writeAPI.Flush()

	// Get query client
	queryAPI := client.QueryAPI("9king.club")

	query := `from(bucket:"test1")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`

	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	// Iterate over query response
	for result.Next() {
		// Notice when group key has changed
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// Access data
		fmt.Printf("value: %v\n", result.Record().Value())
	}
	// check for an error
	if result.Err() != nil {
		fmt.Printf("query parsing error: %\n", result.Err().Error())
	}

}
