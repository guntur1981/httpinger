package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const port = ":3000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ping", ping)

	log.Println("server listening on port", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("a new client is connected")
	defer log.Println("a client is disconnected")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	targetUrl := r.URL.Query().Get("url")
	log.Println("pinging to", targetUrl)

	_, err := url.ParseRequestURI(targetUrl)
	if err != nil {
		fmt.Fprintf(w, "data: %v\n\n", err)
		w.(http.Flusher).Flush()
		return
	}

	for {
		select {
		case <-time.Tick(2 * time.Second):
			// send head request just to receive the headers (no need to get the content)
			t := time.Now()
			resp, err := http.Head(targetUrl)
			if err != nil {
				fmt.Fprintf(w, "data: error making request status=%d\n\n", resp.StatusCode)
				w.(http.Flusher).Flush()
				return
			}
			dur := time.Since(t)

			fmt.Fprintf(w, "data: Reply from %s status=%d time=%dms\n\n", targetUrl, resp.StatusCode, dur.Milliseconds())
			w.(http.Flusher).Flush()

		case <-r.Context().Done():
			// received browser disconnection
			return
		}
	}
}
