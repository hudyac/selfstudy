package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/golang/glog"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}
	glog.Info("set header")

	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	glog.Info(version)

	if r.URL.Path == "/healthz" {
		glog.Info("quest healthz")
		w.WriteHeader(http.StatusOK)
	}

	clientIP := r.RemoteAddr
	glog.Infof("Client IP: %s, HTTP status code: %d", clientIP, http.StatusOK)
}

func main() {
	flag.Set("v", "4")
	http.HandleFunc("/", handlerFunc)
	glog.Info("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		glog.Fatal(err)
	}
}
