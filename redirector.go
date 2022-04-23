package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var response = `
<html>
<head>
<title>Redirecting...</title>
<script>window.location.protocol = 'https:';</script>
</head>
<body>
Just switch to https up there ↑
</body>
</html>
`

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
		log.Printf("Got a %s request from %s: %s (%s)",
			r.Proto, r.RemoteAddr, r.URL, r.Host)
		totalReqs.WithLabelValues().Inc()
		// this matches urls like chrissx.de.evil.com, but
		// there are no ways to exploit that (except if there
		// are other misdesigns)
		if strings.Contains(r.Host, "chrissx.de") ||
			strings.Contains(r.Host, "chrissx.eu") ||
			strings.Contains(r.Host, "zerm.eu") ||
			strings.Contains(r.Host, "zerm.link") {
			var url = url.URL{}
			url.Host = r.Host
			url.Scheme = "https"
			url.Path = r.URL.Path
			w.Header().Add("Location", url.String())
			w.WriteHeader(307)
			hostReqs.WithLabelValues().Inc()
		}
		fmt.Fprintf(w, response)
	})

	http.ListenAndServe(":80", nil)
}
