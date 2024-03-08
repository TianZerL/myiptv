package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type Channel struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Logo  string `json:"logo"`
	Group string `json:"group"`
}

type PlayList struct {
	Channels []Channel `json:"channels"`
}

var reLogo = regexp.MustCompile("tvg-logo=\"([^\"]*)\"")
var reGroup = regexp.MustCompile("group-title=\"([^\"]*)\"")

func ParseM3U(path string) *PlayList {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var channel *Channel
	var channels []Channel
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#EXTINF") {
			channel = new(Channel)
			channel.Name = strings.TrimSpace(line[strings.LastIndex(line, ",")+1:])
			channel.Logo = strings.TrimSpace(reLogo.FindStringSubmatch(line)[1])
			channel.Group = strings.TrimSpace(reGroup.FindStringSubmatch(line)[1])
		} else if !strings.HasPrefix(line, "#") {
			idx := strings.LastIndex(line, "/rtp/")
			if idx == -1 {
				idx = strings.LastIndex(line, "/udp/")
			}
			if idx == -1 {
				idx = 0
			}
			channel.Url = strings.TrimSpace(line[idx:])
			channels = append(channels, *channel)
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Println(err)
		return nil
	}

	return &PlayList{channels}
}
