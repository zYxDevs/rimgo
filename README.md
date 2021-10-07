# rimgu: image host alternative frontend

## rimgu is an alternative frontend / proxy for imgur

It's read-only and works without browser JavaScript. Images and albums can be viewed without wasting resources and anyonymity from downloading and running tracking scripts. No sign-up nags.

It's lightweight and easy to configure.

Inspired by and (soon) integratable with:

* [searx](https://github.com/searx/searx)
* [teddit](https://codeberg.org/teddit/teddit)
* [Privacy Redirect](https://github.com/SimonBrazell/privacy-redirect)
* [nitter](https://github.com/zedeus/nitter)
* [bibliogram](https://sr.ht/~cadence/bibliogram/)

This is currently very early stage software. Some things left to implement (contributions welcome!):

[ ] Localization/internationalization
[ ] Pretty CSS styling
[ ] Automatically fetch / rotate / renew client ID
[ ] Support for other popular image sites than only imgur
[ ] Prometheus metrics

Things that are *currently* considered out of scope:

* Uploading/commenting/voting/etc - rimgu is read-only for now.
* Caching, authentication, serving HTTPS, rate limiting etc - Just use a load balancer like haproxy/envoy/nginx/traefik/caddy.
* Anything requiring client-side JS or client-side directly interacting with upstream servers

## Building

### Locally

Dependencies:

* node.js >= v16 (earlier most likely works fine)

```
$ npm install && npm run build
```

### Docker
```
$ docker build -t rimgu:latest .
```

## Running

### Locally

```
$ node dist/index.js
```

### Docker
```
$ docker run -p 8080:8080 -e -it RIMGU_ADDRESS=0.0.0.0 -e RIMGU_PORT=8080 rimgu:latest
```

## Configuration

Rimgu is configured via environment variables. See available variable in [src/config.ts](./src/config.ts).

### API and client ID

Media and galleries can be scraped without authorization through the public web interface.
Some imgur functionality (comments, full albums) requires a provisioned client ID to authenticate requests.

You can get a client ID by opening https://imgur.com in a web browser and looking for requests to `https://api.imgur.com/...?client_id=1234567deadbeef` under "Network" in the developer console.

*To run without API/key*: `RIMGU_USE_API=false`

*To run with API/key*: `RIMGU_USE_API=true RIMGU_IMGUR_CLIENT_ID=1234567deadbeef`

This key may be used to track you and could be banned when overused. Keep this in mind before exposing a public instance with a client key associated with your personal imgur account. Consider any ToS you may have signed before associating a personal imgur account with a public instance.


## Contributing

PRs are welcome! Before submitting a PR, please make sure linting passes successfully:

```
$ npm run lint
```

This software is released under the AGPL 3.0 license. In short, this means that if you make any modifications to the code and then publish the result (e.g. by hosting the result on a webserver), you must publicly distribute your changes and declare that they also use AGPL 3.0.

You are also requested to not remove attribution and donation information from forks and publications.