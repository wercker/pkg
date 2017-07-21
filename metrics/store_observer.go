package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/wercker/pkg/reflectutil"
)

// labels are the labels that are send to prometheus
var labels = []string{
	"method",
}

// NewStoreObserver creates a new StoreObserver
func NewStoreObserver() *StoreObserver {
	durationOpts := prometheus.HistogramOpts{
		Name: "store_handling_seconds",
		Help: "Histogram of response latency (seconds) of store calls that had been handled by the server",
	}
	duration := prometheus.NewHistogramVec(durationOpts, labels)

	counterOpts := prometheus.CounterOpts{
		Name: "store_handled_total",
		Help: "Total number of store calls completed on the server, regardless of success or failure",
	}
	counter := prometheus.NewCounterVec(counterOpts, labels)

	prometheus.MustRegister(duration)
	prometheus.MustRegister(counter)

	return &StoreObserver{duration: duration, counter: counter}
}

// StoreObserver encapsulates exposing of store specific metrics to Prometheus.
type StoreObserver struct {
	duration *prometheus.HistogramVec
	counter  *prometheus.CounterVec
}

// Preload counters and histograms for each method defined on s.
func (m *StoreObserver) Preload(s interface{}) {
	methods := reflectutil.GetMethods(s)
	for _, method := range methods {
		if method == "Close" || method == "Healthy" {
			continue
		}
		m.counter.WithLabelValues(method)
		m.duration.WithLabelValues(method)
	}
}

// Observe immediately increments the counter for method and returns a func
// which will observe an metric item in duration based on the duration.
func (s *StoreObserver) Observe(method string) func() {
	start := time.Now()

	counter := s.counter.WithLabelValues(method)
	counter.Add(1)

	duration := s.duration.WithLabelValues(method)
	return func() {
		duration.Observe(time.Now().Sub(start).Seconds())
	}
}
