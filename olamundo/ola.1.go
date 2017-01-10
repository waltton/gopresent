package main

import "net/http"

func oláMundo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
}

func main() {
	http.HandleFunc("/", oláMundo)
	http.ListenAndServe("0.0.0.0:3000", nil)
}
