package main

import "testing"

func Test_initMetricPusher(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMetricPusher()
		})
	}
}

func Test_metricsPusher(t *testing.T) {
	type args struct {
		metricsBundlerURL  string
		metricsPusherDelay int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metricsPusher(tt.args.metricsBundlerURL, tt.args.metricsPusherDelay)
		})
	}
}
