package main

import (
	"net/http"
	"path/filepath"
	"fmt"
)

var (
	webapp = http.Dir("webapp/build/unbundled/")
)

func init() {
	http.Handle("/static/", http.FileServer(webapp))
	http.Handle("/manifest.json", http.FileServer(webapp))
	http.Handle("/sw-precache-config.js", http.FileServer(webapp))
	http.Handle("/service-worker.js", http.FileServer(webapp))
	http.Handle("/build/sw-toolbox.map.json",http.FileServer(webapp))
	http.HandleFunc("/api/", handleApi)
	http.HandleFunc("/", handlerFunc)

}
func handleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(string(webapp), "index.html"))
}