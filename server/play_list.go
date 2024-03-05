package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Channel struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Logo string `json:"logo"`
}

type PlayList struct {
	Channels []Channel `json:"channels"`
}

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
			channel.Name = line[strings.LastIndex(line, ",")+1:]

			tags := strings.Split(line, " ")
			for _, tag := range tags {
				if strings.HasPrefix(tag, "tvg-logo") {
					channel.Logo = tag[10 : len(tag)-1]
					break
				}
			}

		} else if !strings.HasPrefix(line, "#") {
			idx := strings.LastIndex(line, "/rtp")
			if idx == -1 {
				idx = strings.LastIndex(line, "/udp")
			}
			if idx == -1 {
				idx = 0
			}
			channel.Url = line[idx:]
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
