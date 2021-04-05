package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/common/log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// metrics that only increase
var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "aaa_counter", //name of metrics (aaa just to appears on top o
	Help: "aaa_desc",    // metric desc
})

// metrics that only increase
var counterWithLabels = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "aaa_counter_vec", //name of metrics (aaa just to appears on top o
	Help: "aaa_desc_vec",    // metric desc
}, []string{"name"})

//metrics that can also decrease

var gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "aaa_gauge", //name of metrics (aaa just to appears on top o
	Help: "aaa_gauge", // metric desc
})

var histogram = promauto.NewHistogram(prometheus.HistogramOpts{
	Name: "aaa_histogram",
	Help: "aaa_histogram",
	Buckets: prometheus.DefBuckets,
})

func main() {

	measure(counter.Inc)
	measure(counterWithLabels.With(map[string]string{"name": "p1"}).Inc)
	measure(counterWithLabels.With(map[string]string{"name": "p2"}).Inc)
	measure(counterWithLabels.With(map[string]string{"name": "p2"}).Inc)

	measure(func() {
		histogram.Observe(rand.Float64() / 1.5)
	})

	measure(func() {
		gauge.Set(rand.Float64() *30)
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func measure(f func()) {
	go func() {
		for {
			time.Sleep(3 * time.Second)
			f()
		}

	}()
}
