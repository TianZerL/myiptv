import './mpegts.js'

var videoElement = document.getElementById('videoElement');
var channelListElement = document.getElementById('channelList');
var groupListElement = document.getElementById('groupList');
var logoElement = document.getElementById('logoElement');

var groups = { 'all': [] };
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

function setChannelList(channelOptions) {
    channelListElement.options.length = channelOptions.length + 1
    for (let i = 0; i < channelOptions.length; i++) {
        channelListElement.options[i + 1] = channelOptions[i];
    }
    channelListElement.options.selectedIndex = 0;
}

fetchPlayList('/api/playlist/').then(data => {
    for (const channel of data.channels) {
        if (groups[channel.group] == undefined) {
            let groupOption = document.createElement('option');
            groupOption.text = channel.group;
            groupListElement.add(groupOption);
            groups[channel.group] = [];
        }
        let channelOption = document.createElement('option');
        channelOption.text = channel.name;
        channelOption.channel = channel;
        groups[channel.group].push(channelOption);
        groups['all'].push(channelOption);
    }
    setChannelList(groups['all'])
}).catch(err => console.log(err))

channelListElement.addEventListener('change', e => {
    const option = e.target.options[e.target.selectedIndex];
    logoElement.src = option.channel.logo
    play(option.channel.url);
})

groupListElement.addEventListener('change', e => {
    const option = e.target.options[e.target.selectedIndex];
    setChannelList(groups[option.text])
})
