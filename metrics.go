package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func initMetricPusher() {
	log.Println("[metrics] Initializing the metricsPusher")
	go metricsPusher(METRICBUNDLERURL, metricsPusherDelay)
}

func metricsPusher(metricsBundlerURL string, metricsPusherDelay int64) {
	log.Printf("[metrics pusher] The metrics pusher was startet, it will push metrics every %d seconds to the metrics-bundler at %s", metricsPusherDelay, metricsBundlerURL)

	// if the metrics map does not exist yet, create it
	if len(metrics) == 0 {
		metrics = make(map[string]float64)
	}

	// define the endpoint URL to where the metrics should be pushed
	metricsEndpoint := fmt.Sprintf("http://%s/metrics", metricsBundlerURL)

	// infinite loop posting the metrics every n seconds to the metrics-bundler
	for {
		log.Println("[metrics pusher] Pushing the metrics...")

		// get all keys and values from all items in the metrics map
		for key, value := range metrics {

			// make a POST request to the metrics-bundler
			_, err := http.PostForm(metricsEndpoint,
				url.Values{
					"key": {
						fmt.Sprintf("%v", key),
					},
					"value": {
						fmt.Sprintf("%v", value),
					},
				})

			// handle potential errors by panicking
			if err != nil {
				panic(err)
			}

		}

		log.Printf("[metrics pusher] Done pushing the metrics. Waiting %d seconds...", metricsPusherDelay)

		// sleep for 5 seconds
		time.Sleep(time.Second * time.Duration(metricsPusherDelay))
	}
}
