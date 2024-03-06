import './mpegts.js'

var videoElement = document.getElementById('videoElement');
var channelListElement = document.getElementById('channelList');
var logoElement = document.getElementById('logoElement');

var player = null;

async function fetchPlayList(url) {
    const response = await fetch(url);
    const data = await response.json();
    return response.ok ? data : Promise.reject(data);
}

function play(url) {
    if (mpegts.getFeatureList().mseLivePlayback) {
        if (player) {
            player.pause();
            player.unload();
            player.detachMediaElement();
            player.destroy();
            player = null;
        }
        player = mpegts.createPlayer({
            type: 'mse',
            isLive: true,
            url: url
        }, {
            autoCleanupSourceBuffer: true
        });
        player.attachMediaElement(videoElement);
        player.load();
        player.play();
    }
}

fetchPlayList('/api/playlist/').then(data => {
    for (const channel of data.channels) {
        var option = document.createElement('option');
        option.text = channel.name;
        option.channel = channel;
        channelListElement.add(option);
    }
}).catch(err => console.log(err))

channelListElement.addEventListener('change', e => {
    var option = e.target.options[e.target.selectedIndex];
    logoElement.src = option.channel.logo
    play(option.channel.url);
})
