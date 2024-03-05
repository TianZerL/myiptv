package main

import "os"

type Config struct {
	Port         string
	ProxyUrl     string
	WebHomePath  string
	PlayListPath string
}

func NewConfig() *Config {
	config := Config{
		Port:         os.Getenv("SERVER_PORT"),
		ProxyUrl:     os.Getenv("PROXY_URL"),
		WebHomePath:  os.Getenv("WEB_HOME_PATH"),
		PlayListPath: os.Getenv("PLAY_LIST_PATH"),
	}

	if config.Port == "" {
		config.Port = "4000"
	}
	if config.ProxyUrl == "" {
		config.ProxyUrl = "http://192.168.1.20:4000"
	}
	if config.WebHomePath == "" {
		config.WebHomePath = "./web"
	}
	if config.PlayListPath == "" {
		config.PlayListPath = config.WebHomePath + "/playlist.m3u"
	}
	return &config
}
