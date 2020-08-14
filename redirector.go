package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	var response = "<html>"
	response += "<head>"
	response += "<title>Redirecting...</title>"
	response += "<script>"
	response += "window.location.protocol = 'https:'"
	response += "</script>"
	response += "</head>"
	response += "<body>"
	response += "Just switch to https up there ↑"
	response += "</body>"
	response += "</html>"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got an %s request from %s: %s (%s)",
			r.Proto, r.RemoteAddr, r.URL, r.Host)
		// this matches urls like chrissx.de.evil.com, but
		// there are no ways to exploit that (except if there
		// are other misdesigns)
		if strings.Contains(r.Host, "chrissx.de") ||
			strings.Contains(r.Host, "chrissx.eu") ||
			strings.Contains(r.Host, "zerm.eu") ||
			strings.Contains(r.Host, "zerm.link") {
			var url = r.URL
			url.Scheme = "https"
			w.Header().Add("Location", url.String())
			w.WriteHeader(307)
		}
		fmt.Fprintf(w, response)
	})

	http.ListenAndServe(":80", nil)
}
