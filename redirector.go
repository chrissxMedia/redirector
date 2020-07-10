package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got a request over %s from %s: %s.",
			r.Proto, r.RemoteAddr, r.URL)
		fmt.Fprintf(w, "<html><head><title>Redirecting...</title><script>window.location.protocol = 'https:'</script></head><body>Just replace that \"http://\" with a \"https://\" up there ^</body></html>")
	})

	http.ListenAndServe(":80", nil)
}
