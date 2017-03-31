package main;

import (
	"net/http"
	"fmt"
	);

func d (w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "dooooog");
}

func m (w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Felix");
}

func main() {
	http.HandleFunc("/dog/", d );
	http.HandleFunc("/me/", m);
	http.HandleFunc("/", m);
	http.ListenAndServe(":8080", nil);
}