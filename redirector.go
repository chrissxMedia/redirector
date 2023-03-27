package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/chrissxMedia/cm3.go"
	"github.com/prometheus/client_golang/prometheus"
)

//go:embed response.min.html
var response string

func main() {
	var totalReqs = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "redirector_total_requests",
		Help: "Total number of HTTP requests coming in.",
	}, []string{})

	var hostReqs = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "redirector_host_requests",
		Help: "Number of HTTP requests that contain the Host header.",
	}, []string{})

	cm3.HandleMetrics(totalReqs, hostReqs)

	cm3.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		totalReqs.WithLabelValues().Inc()
		if strings.Contains(r.Host, ".") && !strings.Contains(r.Host, ":") {
			var url = url.URL{Host: r.Host, Scheme: "https", Path: r.URL.Path}
			w.Header().Add("Location", url.String())
			w.WriteHeader(301)
			hostReqs.WithLabelValues().Inc()
		}
		fmt.Fprintf(w, response)
	})

	cm3.ListenAndServeHttp(":80", nil)
}
