package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
)

func auth(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth() // HLxxx
		if !(ok && password == "gopher-belô") {
			w.WriteHeader(401)
			return
		}
		ctx := context.WithValue(r.Context(), "username", username) // HLxxx
		next(w, r.WithContext(ctx))                                 // HLxxx
	}
}

func oláGophers(w http.ResponseWriter, r *http.Request) {
	gopher := r.Context().Value("username")
	fmt.Fprintf(w, "Olá %s", gopher)
}

func main() {
	var addr string

	flag.StringVar(&addr, "addr", "0.0.0.0:3000", "TCP address to listen on")
	flag.Parse()

	http.HandleFunc("/", auth(oláGophers))

	http.ListenAndServe(addr, nil)
}
