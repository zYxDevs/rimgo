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

Available at https://rimgo.codeberg.page/ or https://codeberg.org/rimgo/instances

## Install

See [Install](https://rimgo.codeberg.page/docs/getting-started/install/).

## Configuration

See [Configuration](https://rimgo.codeberg.page/docs/usage/configuration/).

## Contributing
Pull requests are welcome! If you have any questions or bug reports, open an [issue](https://codeberg.org/rimgo/rimgo/issues/new). Please remember to follow our [Code of Conduct](https://rimgo.codeberg.page/docs/code-of-conduct/)!

## License
This software is released under the AGPL-3.0 license. If you make any modifications to the code and distribute it (including use on a network server), you must publicly distribute your changes and release them under the AGPL-3.0.

## Legal notice
rimgo does not allow uploads or host any content, media. All content on any rimgo instances is from Imgur™. Imgur is a trademark of Imgur, Inc. Any issues with content on rimgo should be be reported to Imgur.