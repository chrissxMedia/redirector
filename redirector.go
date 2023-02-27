package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	prometheus.MustRegister(totalReqs)
	prometheus.MustRegister(hostReqs)
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s request from %s: %s (%s)",
			r.Proto, r.RemoteAddr, r.URL, r.Host)
		totalReqs.WithLabelValues().Inc()
		if strings.Contains(r.Host, ".") && !strings.Contains(r.Host, ":") {
			var url = url.URL{}
			url.Host = r.Host
			url.Scheme = "https"
			url.Path = r.URL.Path
			w.Header().Add("Location", url.String())
			w.WriteHeader(301)
			hostReqs.WithLabelValues().Inc()
		}
		fmt.Fprintf(w, response)
	})

	http.ListenAndServe(":80", nil)
}
