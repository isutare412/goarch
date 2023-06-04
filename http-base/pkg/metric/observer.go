package metric

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var globalObserver Observer = &noopObserver{}

type Observer interface {
	ObserveHTTPRequestEvent(method, url string)
	ObserveHTTPResponseEvent(method, url string, statusCode int, start time.Time)
}

func ObserveHTTPRequestEvent(method, url string) {
	globalObserver.ObserveHTTPRequestEvent(method, url)
}

func ObserveHTTPResponseEvent(method, url string, statusCode int, start time.Time) {
	globalObserver.ObserveHTTPResponseEvent(method, url, statusCode, start)
}

func Gatherer() prometheus.Gatherer {
	client, ok := globalObserver.(*client)
	if !ok {
		return nil
	}
	return client.gatherer
}

func Registerer() prometheus.Registerer {
	client, ok := globalObserver.(*client)
	if !ok {
		return nil
	}
	return client.registerer
}

type noopObserver struct{}

func (*noopObserver) ObserveHTTPRequestEvent(string, string)                  {}
func (*noopObserver) ObserveHTTPResponseEvent(string, string, int, time.Time) {}
