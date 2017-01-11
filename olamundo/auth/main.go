package main

import (
	"flag"
	"net/http"
)

func auth(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, password, ok := r.BasicAuth()
		if !(ok && password == "gopher-belô") {
			w.WriteHeader(401)
			return
		}
		next(w, r)
	}
}

func oláGophers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Gopher"))
}

func main() {
	var addr string

	flag.StringVar(&addr, "addr", "0.0.0.0:3000", "TCP address to listen on")
	flag.Parse()

	http.HandleFunc("/", auth(oláGophers))

	http.ListenAndServe(addr, nil)
}
