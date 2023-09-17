<img alt="" src="https://codeberg.org/rimgo/rimgo/raw/branch/main/static/img/rimgo.svg" width="96" height="96" />

# rimgo
An alternative frontend for Imgur. Originally based on [rimgu](https://codeberg.org/3np/rimgu).

<a href="https://www.gnu.org/licenses/agpl-3.0.en.html">
  <img alt="License: AGPLv3" src="https://shields.io/badge/License-AGPL%20v3-blue.svg" height="20px">
</a>
<a href="https://matrix.to/#/#rimgo:nitro.chat">
  <img alt="Matrix" src="https://img.shields.io/badge/chat-matrix-blue" height="20px">
</a>

## Table of Contents
- [Features](#features)
- [Comparison](#comparison)
  - [Speed](#speed)
  - [Privacy](#privacy)
- [Usage](#usage)
- [Instances](#instances)
  - [Clearnet](#clearnet)
  - [Tor](#tor)
- [Contributing](#contributing)
- [License](#license)

### Documentation

Our new documentation is now available at [https://rimgo.codeberg.page/docs/](https://rimgo.codeberg.page/docs/)!

- [Install](https://rimgo.codeberg.page/docs/getting-started/install/)
- [Configuration](https://rimgo.codeberg.page/docs/usage/configuration/)
- [Redirection](https://rimgo.codeberg.page/docs/usage/configuration/)
- [Instance privacy](https://rimgo.codeberg.page/docs/usage/instance-privacy/)

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

### Privacy
Imgur collects information about your device and uses tracking cookies for advertising, this is mentioned in their [privacy policy](https://imgur.com/privacy/). [Blacklight](https://themarkup.org/blacklight) found 31 trackers and 87 third-party cookies.

See what cookies and trackers Imgur uses and where your data gets sent: https://themarkup.org/blacklight?url=imgur.com

## Usage
Replace imgur.com or i.imgur.com with the instance domain. For i.stack.imgur.com, replace i.stack.imgur.com with the instance domain and add stack/ before the media ID. You can use a browser extension to do this [automatically](#automatically-redirect-links).

Imgur: `https://imgur.com/gallery/j2sOQkJ` -> `https://rimgo.bcow.xyz/gallery/j2sOQkJ`
Stack Overflow: `https://i.stack.imgur.com/KnO3v.jpg?s=64&g=1` -> `https://rimgo.bcow.xyz/stack/KnO3v.jpg?s=64&g=1`

To automatically redirect Imgur links, see [Redirection](https://rimgo.codeberg.page/docs/usage/redirection/).

## Instances
Open an issue to have your instance listed here! See the rules for the instance list [here](https://rimgo.codeberg.page/docs/usage/instance-list-rules/).

> For more details on instance privacy, see https://rimgo.codeberg.page/docs/usage/instance-privacy/

### Clearnet

| URL                                                        	      | Country      | Provider                 | Privacy               | Notes |
| :---------------------------------------------------------------- | :----------- | :----------------------- | :-------------------- | :---- |
| [rimgo.pussthecat.org](https://rimgo.pussthecat.org)       	      | 🇩🇪 DE        | Hetzner                  | ⚠️ Data collected      |       |
| [rimgo.totaldarkness.net](https://rimgo.totaldarkness.net) 	      | 🇨🇦 CA        | Vultr                    | ✅ Data not collected |       |
| [rimgo.bus-hit.me](https://rimgo.bus-hit.me)               	      | 🇨🇦 CA        | Oracle                   | ⚠️ Data collected      |       |
| [imgur.artemislena.eu](https://imgur.artemislena.eu)       	      | 🇩🇪 DE        | Vodafone Deutschland     | ✅ Data not collected | Self-hosted, provider is ISP |
| [rimgo.vern.cc](https://rimgo.vern.cc)                            | 🇺🇸 US        | OVHCloud                 | ✅ Data not collected | [Edited theme](https://git.vern.cc/root/modifications/src/branch/master/rimgo) |
| [rim.odyssey346.dev](https://rim.odyssey346.dev/)                 | 🇫🇷️ FR        | Trolling Solutions (OVH) | ✅ Data not collected |       |
| [i.habedieeh.re](https://i.habedieeh.re/)                         | 🇨🇦️ CA        | Oracle Cloud             | ✅ Data not collected |       |
| [rimgo.hostux.net](https://rimgo.hostux.net/)	                    | 🇫🇷️ FR        | Gandi	                   | ⚠️ Data collected      |       |
| [ri.zzls.xyz](https://ri.zzls.xyz/)                               | 🇨🇱 CL        | TELEFÓNICA CHILE         | ✅ Data not collected | Self-hosted, provider is ISP |
| [rimgo.lunar.icu](https://rimgo.marcopisco.com/)                  | 🇩🇪 DE        | Cloudflare               | ✅ Data not collected |       |
| [imgur.010032.xyz](https://imgur.010032.xyz/)                     | 🇰🇷 KR        | Oracle Cloud             | ✅ Data not collected |       |
| [rimgo.kling.gg](https://rimgo.kling.gg/)                         | 🇳🇱 NL        | RamNode                  | ✅ Data not collected |       |
| [i.01r.xyz](https://i.01r.xyz/)                                   | 🇺🇸 US        | Cloudflare               | ✅ Data not collected |       |
| [rimgo.projectsegfau.lt](https://rimgo.projectsegfau.lt/)         | 🌐 Global     | See below                | ✅ Data not collected |       |
| [rimgo.eu.projectsegfau.lt](https://rimgo.eu.projectsegfau.lt/)   | 🇫🇷 FR        | Orange S.A.              | ✅ Data not collected |       |
| [rimgo.us.projectsegfau.lt](https://rimgo.us.projectsegfau.lt/)   | 🇺🇸 US        | Racknerd                 | ✅ Data not collected |       |
| [rimgo.in.projectsegfau.lt](https://rimgo.in.projectsegfau.lt/)   | 🇮🇳 IN        | Airtel                   | ✅ Data not collected |       |
| [rimgo.whateveritworks.org](https://rimgo.whateveritworks.org/)   | 🇩🇪 DE        | Cloudflare               | ✅ Data not collected |       |
| [rimgo.nohost.network](https://rimgo.nohost.network/)             | 🇲🇽 MX        | Telmex                   | ✅ Data not collected |       |
| [rimgo.catsarch.com](https://rimgo.catsarch.com/)                 | 🇺🇸 US        | Comcast                  | ✅ Data not collected | Self-hosted, provider is ISP |
| [rimgo.frontendfriendly.xyz](https://rimgo.frontendfriendly.xyz/) | 🇩🇪 DE        | Hetzner                  | ⚠️ Data collected      |       |
| [rimgo.quantenzitrone.eu](https://rimgo.quantenzitrone.eu/)       | 🇨🇿 CZ        | vpsFree.cz               | ✅ Data not collected |       |

### Tor

| URL | Privacy | Notes                    |
| :-- | :------ | :----------------------- |
| [rimgo.vernccvbvyi5qhfzyqengccj7lkove6bjot2xhh5kajhwvidqafczrad.onion](http://rimgo.vernccvbvyi5qhfzyqengccj7lkove6bjot2xhh5kajhwvidqafczrad.onion) | ✅ Data not collected | Onion of rimgo.vern.cc         |
| [imgur.lpoaj7z2zkajuhgnlltpeqh3zyq7wk2iyeggqaduhgxhyajtdt2j7wad.onion](http://imgur.lpoaj7z2zkajuhgnlltpeqh3zyq7wk2iyeggqaduhgxhyajtdt2j7wad.onion) | ✅ Data not collected | Onion of imgur.artemislena.eu  |
| [rim.odysfvr23q5wgt7i456o5t3trw2cw5dgn56vbjfbq2m7xsc5vqbqpcyd.onion](http://rim.odysfvr23q5wgt7i456o5t3trw2cw5dgn56vbjfbq2m7xsc5vqbqpcyd.onion)     | ⚠️ Data collected |  |
| [tdp6uqjtmok723suum5ms3jbquht6d7dssug4cgcxhfniatb25gcipad.onion](http://tdp6uqjtmok723suum5ms3jbquht6d7dssug4cgcxhfniatb25gcipad.onion)             | ✅ Data not collected | Onion of rimgo.privacytools.io |
| [i.habeehrhadazsw3izbrbilqajalfyqqln54mrja3iwpqxgcuxnus7eid.onion](http://i.habeehrhadazsw3izbrbilqajalfyqqln54mrja3iwpqxgcuxnus7eid.onion/)        | ✅ Data not collected | Onion of i.habedieeh.re |
| [rimgo.zzlsghu6mvvwyy75mvga6gaf4znbp3erk5xwfzedb4gg6qqh2j6rlvid.onion](http://rimgo.zzlsghu6mvvwyy75mvga6gaf4znbp3erk5xwfzedb4gg6qqh2j6rlvid.onion/) | ✅ Data not collected | Onion of ri.zzls.xyz |
| [tdn7zoxctmsopey77mp4eg2gazaudyhgbuyytf4zpk5u7lknlxlgbnid.onion/](http://tdn7zoxctmsopey77mp4eg2gazaudyhgbuyytf4zpk5u7lknlxlgbnid.onion/) | ✅ Data not collected | Onion of rimgo.kling.gg |
| [rimgo.pjsfkvpxlinjamtawaksbnnaqs2fc2mtvmozrzckxh7f3kis6yea25ad.onion](http://rimgo.pjsfkvpxlinjamtawaksbnnaqs2fc2mtvmozrzckxh7f3kis6yea25ad.onion/) | ✅ Data not collected | Onion of rimgo.eu.projectsegfau.lt |
| [rimgo.skunky7dhv7nohsoalpwe3sxfz3fbkad7r3wk632riye25vqm3meqead.onion](http://rimgo.skunky7dhv7nohsoalpwe3sxfz3fbkad7r3wk632riye25vqm3meqead.onion/) | ✅ Data not collected | Self-hosted, provider is ISP, in Donetsk |

### I2P

| URL | Privacy | Notes                    |
| :-- | :------ | :----------------------- |
| [rimgo.i2p](http://rimgo.i2p) | ✅ Data not collected | i.habedieeh.re on I2P |
| [rimgov7l2tqyrm5txrtvhtnfyrzkc5d7ipafofavchbnnyog4r3q.b32.i2p](http://rimgov7l2tqyrm5txrtvhtnfyrzkc5d7ipafofavchbnnyog4r3q.b32.i2p) | ✅ Data not collected | Same as rimgo.i2p |
| [rimgo.zzls.i2p](http://rimgo.zzls.i2p) | ✅ Data not collected | ri.zzls.xyz on I2P |
| [p57356k2xwhxrg2lxrjajcftkrptv4zejeeblzfgkcvpzuetkz2a.b32.i2p](http://p57356k2xwhxrg2lxrjajcftkrptv4zejeeblzfgkcvpzuetkz2a.b32.i2p) | ✅ Data not collected | Same as rimgo.zzls.i2p |
| [ovzamsts5czfx3jasbbhbccyyl2z7qmdngtlqxdh4oi7abhdz3ia.b32.i2p](http://ovzamsts5czfx3jasbbhbccyyl2z7qmdngtlqxdh4oi7abhdz3ia.b32.i2p) | ✅ Data not collected | rimgo.kling.gg on I2P |

## Install

See [Install](https://rimgo.codeberg.page/docs/getting-started/install/).

## Configuration

See [Configuration](https://rimgo.codeberg.page/docs/usage/configuration/).

## Contributing
Pull requests are welcome! If you have any questions or bug reports, open an [issue](https://codeberg.org/rimgo/rimgo/issues/new).

## License
This software is released under the AGPL-3.0 license. If you make any modifications to the code and distribute it (including use on a network server), you must publicly distribute your changes and release them under the AGPL-3.0.
