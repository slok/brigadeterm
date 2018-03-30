# Brigadeterm [![Build Status](https://travis-ci.org/slok/brigadeterm.svg?branch=master)](https://travis-ci.org/slok/brigadeterm)

Brigadeterm is a text based dashboard for [Brigade][brigade-url] pipeline system.

![projects](screenshots/brigadeterm-latest-builds.gif)
![pipelines](screenshots/jobs2.png)
![builds](screenshots/builds.png)

## Run

First you need to download the binary from [releases][releases-url].

Brigadeterm uses kubectl configuration, so you need access to the cluster using kubectl.

```bash
brigadeterm --namespace {BRIGADE_NAMESPACE}
```

If you have problems with the rendering on your terminal try setting the `TERM` env var. For example:

```bash
TERM=xterm brigadeterm --namespace {BRIGADE_NAMESPACE}
```


## Build

To build just type:

```shell
make build-binary
```

Or use go directly:

```bash
go get -u github.com/slok/brigadeterm/cmd/brigadeterm/...
```

## Screenshots

[Here](screenshots) you have some screenshots.

[brigade-url]: https://brigade.sh
[releases-url]: https://github.com/slok/brigadeterm/releases/latest
