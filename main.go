package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

//go:generate go get -u github.com/programmfabrik/esc
//go:generate esc -private -local-prefix-cwd -pkg=main -o=static.go static/ blacklist.txt

var serveUrl, port string
var useLocal bool

func init() {
	flag.BoolVar(&useLocal, "local", false, "Use assets from local filesystem")
	flag.StringVar(&serveUrl, "url", "http://localhost:8080", "The server url this server runs on. (Required for the frontend)")
	flag.StringVar(&port, "port", "8080", "Port to run the server on")

	flag.Parse()
}

func main() {
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "" || req.URL.Path == "/" || req.URL.Path == "/d/" || req.URL.Path == "/d" {
			HandleIndex(rw, req)
			return
		}
		if strings.HasPrefix(req.URL.Path, "/d/") {
			req.URL.Path = strings.TrimPrefix(req.URL.Path, "/d")
			HandleUnShort(rw, req, false)
			return
		}
		HandleUnShort(rw, req, true)
	}

	http.Handle("/static/", http.FileServer(_escFS(useLocal)))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
