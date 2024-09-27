package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type HealthCheckItem struct {
	Name    string            `yaml:"name"`
	Url     string            `yaml:"url"`
	Method  string            `yaml:"method"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
}

func ParseYaml(filename string) ([]HealthCheckItem, error) {
	// Create a struct to hold the YAML data
	var items []HealthCheckItem

	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return items, err
	}

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &items)
	if err != nil {
		fmt.Println(err)
		return items, err
	}

	return items, nil
}

func GetEndpointHosts(items []HealthCheckItem) map[string][]HealthCheckItem {
	hosts := make(map[string][]HealthCheckItem)
	for i := range items {
		parsed, _ := url.Parse(items[i].Url)
		host := parsed.Host
		hosts[host] = append(hosts[host], items[i])
	}
	return hosts
}

func SendHealthCheck(hci HealthCheckItem) string {
	method := "GET"
	if hci.Method != "" {
		method = hci.Method
	}

	req, _ := http.NewRequest(method, hci.Url, bytes.NewBuffer([]byte(hci.Body)))

	for key, value := range hci.Headers {
		req.Header.Add(key, value)
	}

	var start time.Time

	trace := &httptrace.ClientTrace{}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	resp, err := http.DefaultTransport.RoundTrip(req)
	totalTime := time.Since(start)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 && totalTime.Milliseconds() < 500 {
		return "UP"
	} else {
		return "DOWN"
	}
}

func main() {
	var filePath = flag.String("f", "test_input.yml", "The path to use for configuration")
	flag.Parse()
	if _, err := os.Stat(*filePath); errors.Is(err, os.ErrNotExist) {
		fmt.Println(*filePath, "does not exist.")
		return
	}
	items, _ := ParseYaml("test_input.yml")
	hosts := GetEndpointHosts(items)
	for {
		for host, items := range hosts {
			var availability = make(map[int]string)
			for i, hci := range items {
				availability[i] = SendHealthCheck(hci)
			}
			upDawg := 0
			for _, str := range availability {
				if str == "UP" {
					upDawg++
				}
			}
			avPercent := 100 * upDawg / len(items)
			fmt.Printf("%v has %v%% availability percentage\n", host, avPercent)
		}
		time.Sleep(15 * time.Second)
	}

}
