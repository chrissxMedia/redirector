package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got an %s request from %s: %s",
			r.Proto, r.RemoteAddr, r.URL)
		url := r.URL
		url.Scheme = "https"
		fmt.Fprintf(w, "<html>")
		fmt.Fprintf(w, "<head>")
		fmt.Fprintf(w, "<title>Redirecting...</title>")
		fmt.Fprintf(w, "<script>")
		fmt.Fprintf(w, "window.location.protocol = 'https:'")
		fmt.Fprintf(w, "</script>")
		fmt.Fprintf(w, "</head>")
		fmt.Fprintf(w, "<body>")
		fmt.Fprintf(w, "Just switch to https up there ↑")
		fmt.Fprintf(w, "<br/>")
		fmt.Fprintf(w, "Or click <a href='%s'>here</a>", url)
		fmt.Fprintf(w, "</body>")
		fmt.Fprintf(w, "</html>")
	})

	http.ListenAndServe(":80", nil)
}
