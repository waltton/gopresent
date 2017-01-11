package main

import (
	"flag"
	"net/http"
)

func oláMundo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
}

func main() {
	var addr string

	flag.StringVar(&addr, "addr", "0.0.0.0:3000", "TCP address to listen on")
	flag.Parse()

	http.HandleFunc("/", oláMundo)
	http.ListenAndServe(addr, nil)
}
