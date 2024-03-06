package main

import (
	"flag"
	"os"
)

type Config struct {
	Port         string
	ProxyUrl     string
	WebHomePath  string
	PlayListPath string
}

func NewConfig() *Config {
	config := Config{}

	flag.StringVar(&config.Port, "SERVER_PORT", os.Getenv("SERVER_PORT"), "The port for web page")
	flag.StringVar(&config.ProxyUrl, "PROXY_URL", os.Getenv("PROXY_URL"), "URL for your udpxy server")
	flag.StringVar(&config.WebHomePath, "WEB_HOME_PATH", os.Getenv("WEB_HOME_PATH"), "Set custom web static folder")
	flag.StringVar(&config.PlayListPath, "PLAY_LIST_PATH", os.Getenv("PLAY_LIST_PATH"), "Set play list file path")
	flag.Parse()

	if config.Port == "" {
		config.Port = "4000"
	}
	if config.WebHomePath == "" {
		config.WebHomePath = "./web"
	}
	if config.PlayListPath == "" {
		config.PlayListPath = config.WebHomePath + "/playlist.m3u"
	}
	return &config
}
