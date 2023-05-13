# Welcome to AIP2P

## About Project

Distributed Artificial Intelligens protocol implementation base on [libp2p](https://libp2p.io). Peer(s) share CPU and GPU resources with each other and makes Distributed Neural Network (DNN).

Project is written using Go programming language.

[![Go Reference](https://pkg.go.dev/badge/webimizer.dev/aip2p.svg)](https://pkg.go.dev/webimizer.dev/aip2p)

## Supported platforms

![Platforms](/images/platforms.png "Supported platforms")

## Mac OS Intel (64-bit) version

![Aip2p](/images/J68poh.png "Aip2p")

[Download MacOS application for Intel (x64) v0.0.14](https://aip2p.app/downloads/aip2p-0.0.14_amd64.dmg)

## Linux (64-bit) version (tested with KDE desktop environment only)

![Aip2p](/images/linux_kde.png "Aip2p")

[Download Linux application (x64) Debian package v0.0.14](https://aip2p.app/downloads/aip2p-0.0.14_amd64.deb)

Installation steps for Debian-based Linux distro:
1. [Download package](https://aip2p.app/downloads/aip2p-0.0.14_amd64.deb)
```sh
wget https://aip2p.app/downloads/aip2p-0.0.14_amd64.deb
```
2. Install package:
```sh
sudo dpkg -i aip2p-0.0.14_amd64.deb
```

[Download Linux application (x64) tar v0.0.14](https://aip2p.app/downloads/aip2p-0.0.14_amd64.tar.xz)

Installation steps for other (etc. Red Hat) Linux distributions:
1. [Download package](https://aip2p.app/downloads/aip2p-0.0.14_amd64.tar.xz)
```sh
wget https://aip2p.app/downloads/aip2p-0.0.14_amd64.tar.xz
```
2. Extract tar archive with command:
```sh
mkdir aip2p && tar -C ./aip2p -xf aip2p.tar.xz
```
3. Install application with command:
```sh
cd aip2p && sudo make install && cd .. && rm -rf ./aip2p && rm aip2p.tar.xz
```
Notice: this command also removes installation files which are no longer required.

## Windows (64-bit) version

![Aip2pWin64](/images/Win64.png "Aip2p Win64")

[Download Windows setup (x64) v0.0.14](https://aip2p.app/downloads/aip2p-0.0.14_amd64.msi)

## Install from source code (for advanced users)
1. Clone [this repository](https://webimizer.dev/aip2p)
2. Download Go from the [download page](https://go.dev/dl/) and follow instructions
3. Install fyne package:
```sh
go install fyne.io/fyne/v2/cmd/fyne@latest
```
4. Install AIP2P application to your computer with command:
```sh
fyne install
```