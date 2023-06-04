package metric

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func Init() {
	cli := client{
		gatherer:   prometheus.DefaultGatherer,
		registerer: prometheus.DefaultRegisterer,

		httpRequestTotal:             newHTTPRequestTotal(),
		httpResponseTotal:            newHTTPResponseTotal(),
		httpResponseLatencySummary:   newHTTPResponseLatencySummary(),
		httpResponseLatencyHistogram: newHTTPResponseLatencyHistogram(),
	}

	cli.registerer.MustRegister(cli.httpRequestTotal)
	cli.registerer.MustRegister(cli.httpResponseTotal)
	cli.registerer.MustRegister(cli.httpResponseLatencySummary)
	cli.registerer.MustRegister(cli.httpResponseLatencyHistogram)

	globalObserver = &cli
}

type client struct {
	gatherer   prometheus.Gatherer
	registerer prometheus.Registerer

	httpRequestTotal             *prometheus.CounterVec
	httpResponseTotal            *prometheus.CounterVec
	httpResponseLatencySummary   *prometheus.SummaryVec
	httpResponseLatencyHistogram *prometheus.HistogramVec
}

func (cli *client) ObserveHTTPRequestEvent(method, url string) {
	cli.httpRequestTotal.WithLabelValues(method, url).Inc()
}

func (cli *client) ObserveHTTPResponseEvent(method, url string, statusCode int, start time.Time) {
	dur := time.Since(start)
	cli.httpResponseTotal.WithLabelValues(method, url, strconv.Itoa(statusCode)).Inc()
	cli.httpResponseLatencySummary.WithLabelValues(method, url).Observe(dur.Seconds())
	cli.httpResponseLatencyHistogram.WithLabelValues(method, url).Observe(dur.Seconds())
}
