package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/fatih/structs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/waleedsamy/letsgo/config"

	log "github.com/sirupsen/logrus"
)

var (
	httpResponsesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "letsgo",
			Subsystem: "http_server",
			Name:      "http_responses_total",
			Help:      "The count of http responses issued, classified by code and method.",
		},
		[]string{"code", "method"},
	)

	httpResponseLatencies = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "letsgo",
			Subsystem: "http_server",
			Name:      "http_response_latencies",
			Help:      "Duration of sample batch send calls to the remote storage.",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"code", "method"},
	)
)

func init() {
	prometheus.MustRegister(httpResponsesTotal)
	prometheus.MustRegister(httpResponseLatencies)
	log.SetOutput(os.Stdout)
}

type pound float32
type database map[string]pound

func (p pound) String() string {
	return fmt.Sprintf("£%.2f", p)
}

func (d database) bar(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	fmt.Fprintf(w, "%s: %s", "bar", d["bar"])
	elapsed := time.Since(start)
	msElapsed := elapsed / time.Millisecond
	httpResponsesTotal.WithLabelValues("200", r.Method).Inc()
	httpResponseLatencies.WithLabelValues("200", r.Method).Observe(float64(msElapsed))
}

func (d database) foo(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	fmt.Fprintf(w, "%s: %s", "foo", d["foo"])
	elapsed := time.Since(start)
	msElapsed := elapsed / time.Millisecond
	httpResponsesTotal.WithLabelValues("200", r.Method).Inc()
	httpResponseLatencies.WithLabelValues("200", r.Method).Observe(float64(msElapsed))
}

func (d database) baz(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	fmt.Fprintf(w, "%s: %s", "baz", d["baz"])
	elapsed := time.Since(start)
	msElapsed := elapsed / time.Millisecond
	httpResponsesTotal.WithLabelValues("200", r.Method).Inc()
	httpResponseLatencies.WithLabelValues("200", r.Method).Observe(float64(msElapsed))
}
func main() {
	configFile := flag.String("config.file", "", "Let configuration file name.")
	flag.Parse()

	if *configFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var db = database{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}
	http.HandleFunc("/foo", db.foo)
	http.HandleFunc("/bar", db.bar)
	http.Handle("/baz", http.HandlerFunc(db.baz))

	http.Handle("/metrics", promhttp.Handler())

	cfg, err := config.GetConfig(*configFile)
	if err != nil {
		log.WithField("config.file", *configFile).Fatal("unable to parse file")
	}

	log.WithFields(log.Fields{"port": cfg.Port}).Info("Let is ready")
	log.WithFields(structs.Map(cfg)).Info("Let is configured by")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}
