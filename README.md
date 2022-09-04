<img src="https://codeberg.org/video-prize-ranch/rimgo/raw/branch/main/static/img/rimgo.svg" width="96" height="96" />

# rimgo
An alternative frontend for Imgur. Originally based on [rimgu](https://codeberg.org/3np/rimgu).

<a href="https://www.gnu.org/licenses/agpl-3.0.en.html">
  <img alt="License: AGPLv3" src="https://shields.io/badge/License-AGPL%20v3-blue.svg">
</a>
<a href="https://matrix.to/#/#rimgo:nitro.chat">
  <img alt="Matrix" src="https://img.shields.io/badge/chat-matrix-blue">
</a>

## Features
- Lightweight
- No JavaScript
- No ads or tracking
- No sign up or app install prompts
- Bandwidth efficient - automatically uses newer image formats (if enabled)

## Comparison
Comparing rimgo to Imgur.

### Speed
Tested using [Google PageSpeed Insights](https://pagespeed.web.dev/).

| | [rimgo](https://pagespeed.web.dev/report?url=https%3A%2F%2Fi.bcow.xyz%2Fgallery%2FgYiQLWy) | [Imgur](https://pagespeed.web.dev/report?url=https%3A%2F%2Fimgur.com%2Fgallery%2FgYiQLWy) |
| ------------------- | ------- | --------- |
| Performance         | 91      | 28        |
| Request count       | 29      | 340       |
| Resource Size       | 218 KiB | 2,542 KiB |
| Time to Interactive | 1.6s    | 23.8s     |

## Instances
Open an issue to have your instance listed here!

### Clearnet
To help distribute load, consider using instances other than the official one.

| URL                                                        	  | Country      | Provider            | Privacy               | Notes |
| :------------------------------------------------------------ | :----------- | :------------------ | :-------------------- | :---- |
| [i.bcow.xyz](https://i.bcow.xyz) (official)                	  | 🇨🇦️ CA, 🇳🇱 NL | Fly.io	            | ⚠️ Data collected     |       |
| [rimgo.pussthecat.org](https://rimgo.pussthecat.org)       	  | 🇩🇪 DE        | Hetzner             | ⚠️ Data collected     |       |
| [rimgo.totaldarkness.net](https://rimgo.totaldarkness.net) 	  | 🇨🇦 CA        | Vultr               | ✅ Data not collected |       |
| [rimgo.bus-hit.me](https://rimgo.bus-hit.me)               	  | 🇨🇦 CA        | Oracle              | ⚠️ Data collected     |       |
| [rimgo.esmailelbob.xyz](https://rimgo.esmailelbob.xyz)     	  | 🇨🇦 CA        | OVH                 | ⚠️ Data collected     |       |
| [i.actionsack.com](https://i.actionsack.com)               	  | 🇺🇸 US        | Cloudflare          | ❓️ No details         |       |
| [rimgo.privacydev.net](https://rimgo.privacydev.net)       	  | 🇺🇸 US        | Cyber Wurx          | ❓️ No details         |       |
| [imgur.artemislena.eu](https://imgur.artemislena.eu)       	  | 🇩🇪 DE        | Telekom Deutschland | ✅ Data not collected | Self-hosted, provider is ISP |
| [rimgo.vern.cc](https://rimgo.vern.cc)                        | 🇨🇦️ CA        | OVHCloud            | ✅ Data not collected | [Edited theme](https://git.vern.cc/root/modifications/src/branch/master/rimgo) |
| [rimgo.encrypted-data.xyz](https://rimgo.encrypted-data.xyz/)	| 🇫🇷️ FR        | Cloudflare	        | ❓️ No details         |       |
| [rimgo.mha.fi](https://rimgo.mha.fi/)                       	| 🇫🇮 FI        | Hetzner             | ❓️ No details         |       |

### Tor

| URL | Privacy | Notes                    |
| :-- | :------ | :----------------------- |
| [rimgo.lqs5fjmajyp7rvp4qvyubwofzi6d4imua7vs237rkc4m5qogitqwrgyd.onion](http://rimgo.lqs5fjmajyp7rvp4qvyubwofzi6d4imua7vs237rkc4m5qogitqwrgyd.onion) | ⚠️ Data collected | Onion of rimgo.esmailelbob.xyz |
| [rimgo.vernccvbvyi5qhfzyqengccj7lkove6bjot2xhh5kajhwvidqafczrad.onion](http://rimgo.vernccvbvyi5qhfzyqengccj7lkove6bjot2xhh5kajhwvidqafczrad.onion) | ✅ Data not collected | Onion of rimgo.vern.cc         |
| [rimgo.micohauwkjbyw5meacrb4ipicwvwg4xtzl7y7viv53kig2mdcsvwkyyd.onion](http://rimgo.micohauwkjbyw5meacrb4ipicwvwg4xtzl7y7viv53kig2mdcsvwkyyd.onion) | ❓️ No details   | Onion of rimgo.mha.fi          |
| [imgur.lpoaj7z2zkajuhgnlltpeqh3zyq7wk2iyeggqaduhgxhyajtdt2j7wad.onion](http://imgur.lpoaj7z2zkajuhgnlltpeqh3zyq7wk2iyeggqaduhgxhyajtdt2j7wad.onion) | ✅ Data not collected | Onion of imgur.artemislena.eu  |

## Automatically redirect links

### LibRedirect
Use [LibRedirect](https://github.com/libredirect/libredirect) to automatically redirect Imgur links to rimgo!
- [Firefox](https://addons.mozilla.org/firefox/addon/libredirect/)
- [Chromium-based browsers (Brave, Google Chrome)](https://github.com/libredirect/libredirect#install-in-chromium-brave-and-chrome)
- [Edge](https://microsoftedge.microsoft.com/addons/detail/libredirect/aodffkeankebfonljgbcfbbaljopcpdb)

### GreaseMonkey script
There is a script to redirect Imgur links to rimgo.
[https://codeberg.org/zortazert/GreaseMonkey-Redirect/src/branch/main/imgur-to-rimgo.user.js](https://codeberg.org/zortazert/GreaseMonkey-Redirect/src/branch/main/imgur-to-rimgo.user.js)

## Install
rimgo can run on any platform Go compiles on.

### Docker (recommended)
Install [Docker](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/), then clone this repository.
```
git clone https://codeberg.org/video-prize-ranch/rimgo
cd rimgo
```

Edit the `docker-compose.yml` file using your favorite text editor.
```
nvim docker-compose.yml
```

You can now run rimgo.
```
sudo docker-compose up -d
```

#### Automatic updates
[Watchtower](https://containrrr.dev/watchtower/) can automatically update your Docker containers.

Create a new `docker-compose.yml` file or add the watchtower section to your existing `docker-compose.yml` file.
```yml
version: "3"
services:
  watchtower:
    image: containrrr/watchtower
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

### Build from source

#### Requirements
* Go v1.16 or later

Clone the repository and `cd` into it.
```
git clone https://codeberg.org/video-prize-ranch/rimgo
cd rimgo
```

Build rimgo.
```
go build
```

You can now run rimgo.
```
./rimgo
```

To include version information use:
```
go build -ldflags "-X codeberg.org/video-prize-ranch/rimgo/pages.VersionInfo=$(date '+%Y-%m-%d')-$(git rev-list --abbrev-commit -1 HEAD)"
```

(optional) You can use a .env file to set environment variables for configuration.
```
cp .env.example .env
nvim .env
```

## Configuration

rimgo can be configured using environment variables. The path to the .env file can be changed the -c flag.

### Environment variables

| Name                  | Default         |
|-----------------------|-----------------|
| PORT                  | 3000            |
| ADDRESS               | 0.0.0.0         |
| IMGUR_CLIENT_ID       | 546c25a59c58ad7 |
| FORCE_WEBP            | 0               |
| PRIVACY_POLICY        |                 |
| PRIVACY_MESSAGE       |                 |
| PRIVACY_COUNTRY       |                 |
| PRIVACY_PROVIDER      |                 |
| PRIVACY_CLOUDFLARE    |                 |
| PRIVACY_NOT_COLLECTED |                 |
| PRIVACY_IP            |                 |
| PRIVACY_URL           |                 |
| PRIVACY_DEVICE        |                 |
| PRIVACY_DIAGNOSTICS   |                 |

## Contributing

Pull requests are welcome!

This software is released under the AGPL 3.0 license. In short, this means that if you make any modifications to the code and then publish the result (e.g. by hosting the result on a web server), you must publicly distribute your changes and declare that they also use AGPL 3.0.
