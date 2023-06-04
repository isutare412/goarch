package metric

import "github.com/prometheus/client_golang/prometheus"

func newHTTPRequestTotal() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "goarch",
		Subsystem: "httpbase",
		Name:      "http_request_total",
		Help:      "Count of total HTTP requests",
	}, []string{
		"method",
		"url",
	})
}

func newHTTPResponseTotal() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "goarch",
		Subsystem: "httpbase",
		Name:      "http_response_total",
		Help:      "Count of total HTTP responses",
	}, []string{
		"method",
		"url",
		"status",
	})
}

func newHTTPResponseLatencySummary() *prometheus.SummaryVec {
	return prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  "goarch",
		Subsystem:  "httpbase",
		Name:       "http_response_latency_seconds",
		Help:       "Latency summary of HTTP responses",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.03, 0.99: 0.01},
	}, []string{
		"method",
		"url",
	})
}

func newHTTPResponseLatencyHistogram() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "goarch",
		Subsystem: "httpbase",
		Name:      "http_response_latency",
		Help:      "Latency hitogram of HTTP responses",
		Buckets:   prometheus.DefBuckets,
	}, []string{
		"method",
		"url",
	})
}
