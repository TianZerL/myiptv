package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Message struct {
	Msg string `json:"msg"`
}

func ProxyServer(target string) *httputil.ReverseProxy {
	url, err := url.Parse(target)
	if err != nil {
		log.Fatalln(err)
	}
	return httputil.NewSingleHostReverseProxy(url)
}

func main() {
	config := NewConfig()

	http.HandleFunc("/play/list/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		playList := ParseM3U(config.PlayListPath)
		if playList != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(playList)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Message{"Failed to parse play list file"})
		}
	})
	http.Handle("/play/", http.StripPrefix("/play/", http.FileServer(http.Dir(config.WebHomePath))))
	http.Handle("/", ProxyServer(config.ProxyUrl))

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
