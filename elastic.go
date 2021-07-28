package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var (
	es *elasticsearch.Client
)

func init() {
	esAddr := getEnv("esAddr", "http://main.dorabc.com:9200")
	cfg := elasticsearch.Config{
		Addresses: []string{esAddr},
	}
	es, _ = elasticsearch.NewClient(cfg)
}
func saveToES(data *bData) {
	body, err := json.Marshal(data)
	// fmt.Println(string(body))
	checkError(err, "json.Marshal")
	req := esapi.IndexRequest{
		Index:   "battery_data",
		Body:    bytes.NewReader(body),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		checkError(err, "req.Do")
	} else {
		defer res.Body.Close()
		if res.IsError() {
			log.Printf("[%s] Error indexing document clientIp=%v", res.Status(), data.IpAddr)
		} else {
			// Deserialize the response into a map.
			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				fmt.Printf("Error parsing the response body: %s\n", err)
			} else {
				// Print the response status and indexed document version.
				fmt.Printf("[%s] %s; version=%d\n", res.Status(), r["result"], int(r["_version"].(float64)))
			}
		}
	}
}
