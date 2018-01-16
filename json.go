package main

import (
	"fmt"
	"net/http"
	"goji.io"
	"goji.io/pat"
    "github.com/cactus/go-statsd-client/statsd"
)

func hello(w http.ResponseWriter, r *http.Request) {
	    client, _ := statsd.NewClient("127.0.0.1:8125", "helloItsMe")
		   defer client.Close()

    // Send a stat
    client.Inc("stat1", 42, 1.0)
        name := pat.Param(r, "name")
        fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
        mux := goji.NewMux()
        mux.HandleFunc(pat.Get("/hello/:name"), hello)
        http.ListenAndServe("0.0.0.0:8080", mux)
}
