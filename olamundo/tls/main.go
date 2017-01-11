package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func auth(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !(ok && password == "gopher-belô") {
			w.WriteHeader(401)
			return
		}
		ctx := context.WithValue(r.Context(), "username", username)
		next(w, r.WithContext(ctx))
	}
}

func oláGophers(w http.ResponseWriter, r *http.Request) {
	gopher := r.Context().Value("username")
	fmt.Fprintf(w, "Olá %s", gopher)
}

func main() {
	var addr, certFile, keyFile string
	var tls bool

	flag.StringVar(&addr, "addr", "0.0.0.0:3000", "TCP address to listen on")
	flag.StringVar(&certFile, "cert", "", "path to cert file")
	flag.StringVar(&keyFile, "key", "", "path to key file")
	flag.BoolVar(&tls, "tls", false, "use TLS")
	flag.Parse()

	http.HandleFunc("/", auth(oláGophers))

	if tls {
		log.Fatal(http.ListenAndServeTLS(addr, certFile, keyFile, nil))
	}
	log.Fatal(http.ListenAndServe(addr, nil))
}
