package metric

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var (
	requestLabels = []string{"path"}

	reqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_count_total",
			Help: "Total number of HTTP requests made.",
		}, requestLabels,
	)
)

func init() {
	prometheus.MustRegister(collectors.NewBuildInfoCollector())
	prometheus.MustRegister(reqCount)

	go recordUptime()
}

func recordUptime() {
	for range time.Tick(time.Second) {
		reqCount.With(prometheus.Labels{"path": "/hello"}).Inc()
		fmt.Println()
	}
}

// https://blog.csdn.net/a1510841693/article/details/124102654
func Init() {
	server := http.NewServeMux()
	server.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", server)
}
