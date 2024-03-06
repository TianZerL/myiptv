# myiptv
A simple server to watch video stream from udpxy in browser

# How it work
The go server will reverse proxy udpxy and feed your m3u play list to web page, then we use [mpegts.js](https://github.com/xqq/mpegts.js) to play the video stream.

# Screenshot
![Screenshot](/images/screenshot.png)

# Installation
## Docker
### images
your can find images in [DockerHub](https://hub.docker.com/r/tianzer/myiptv)

### docker cli
```shell
docker run -d \
  --name=myiptv \
  --user 1000:1000 \
  -p 4000:4000 \
  --mount type=bind,source=/path/to/playlist,target=/app/web/playlist.m3u \
  -e PROXY_URL=http://your.udpxy.ip:port \
  --restart unless-stopped \
  tianzer/myiptv
```

### docker-compose
```docker-compose
---
services:
  myiptv:
    image: tianzer/myiptv
    container_name: myiptv
    volumes:
      - type: bind
        source: /path/to/playlist
        target: /app/web/playlist.m3u
        read_only: true
    environment:
      - PROXY_URL=http://your.udpxy.ip:port
    ports:
      - 4000:4000
    user: "1000:1000"
    restart: unless-stopped
```

Then open http://your.server.ip:4000 in your browser

### Environment Variables (-e)
| Environment variable | Function                     | default                    |
| -------------------- | ---------------------------- | -------------------------- |
| SERVER_PORT          | The port for web page        | 4000                       |
| PROXY_URL            | URL for your udpxy server    | N/A                        |
| WEB_HOME_PATH        | Set custom web static folder | /app/web                   |
| PLAY_LIST_PATH       | Set play list file path      | WEB_HOME_PATH/playlist.m3u |

## manually
Download [mpegts.js](https://github.com/xqq/mpegts.js) and copy it to `web` folder, then copy your play list file to `web` folder an rename to playlist.m3u.

then run
```shell
cd server
go build -o ../
cd ..
./myiptv
```
Then open http://your.server.ip:4000 in your browser
