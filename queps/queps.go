package queps

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type QPSMeter struct {
	Host        string
	Port        string
	Verbose     bool
	Interval    int
	LabelNames  []string
	LabelValues []string
	MetricPath  string

	requestCount  int64
	requestMetric *prometheus.CounterVec
}

// Start starts measuring the QPS
func (qps *QPSMeter) Start() error {
	// Create counter
	qps.requestMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "qps_total_requests",
			Help: "Total number of HTTP requests",
		},
		qps.LabelNames,
	)

	// Register counter
	if err := prometheus.Register(qps.requestMetric); err != nil {
		return fmt.Errorf("qps error registering metric: %w", err)
	}

	// Initialize metric route
	http.Handle(qps.MetricPath, promhttp.Handler())

	// Handle health check
	// This is done separatly from the main route so healthchecks don't affect QPS
	http.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) })

	// Handle the rest of the requests
	http.HandleFunc("/", qps.MainRoute)

	// Start printing to stdout
	go qps.Printer()

	// Start the service
	address := fmt.Sprintf("%s:%s", qps.Host, qps.Port)
	return http.ListenAndServe(address, nil)
}

// MainRoute handles QPS measuring
func (qps *QPSMeter) MainRoute(w http.ResponseWriter, r *http.Request) {
	// Log request if verbose is on
	if qps.Verbose {
		host, port, _ := net.SplitHostPort(r.RemoteAddr)
		log.Printf("Host: %s, Path: %s, Port: %s, IP: %s\n", r.Host, r.URL.Path, port, host)
	}

	// Increase internal qps meter
	atomic.AddInt64(&qps.requestCount, 1)

	// Increase prometheus counter
	qps.requestMetric.WithLabelValues(qps.LabelValues...).Inc()

	w.WriteHeader(200)
}

// Printer prints the QPS to stdout
func (qps *QPSMeter) Printer() {
	ticker := time.NewTicker(time.Duration(qps.Interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		count := atomic.SwapInt64(&qps.requestCount, 0)
		log.Printf("Requests per second: %.2f\n", float64(count)/float64(qps.Interval))
	}
}
