package main

import (
	"net/http"
	"strings"
)

func reactFileServer(fs http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		staticIndex := strings.Index(req.URL.Path, "/static/");
		if staticIndex == -1 {
			fsHandler := http.StripPrefix(req.URL.Path, fs)
			fsHandler.ServeHTTP(w, req)
		} else {
			fs.ServeHTTP(w, req)
		}
	}
	return http.HandlerFunc(fn)
}

func init() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", reactFileServer(fs))
}
