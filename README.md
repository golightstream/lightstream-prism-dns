# Lightstream Console DNS
This project enables you to run a local version of the Lightstream Console DNS servers we provide.

## Don't want to run your own?
Visit [psdnscheck.golightstream.com](https://psdnscheck.golightstream.com/) to identify the closest DNS server that we host to you.


## Self-hosting
### **Downloads**:

[![Download for Windows](https://img.shields.io/badge/Windows-32bit-blue)](https://google.com)
[![Download for MacOS Intel](https://img.shields.io/badge/Mac-Intel-orange)](https://google.com)
[![Download for MacOS M1](https://img.shields.io/badge/Mac-M1-orange)](https://google.com)
[![Download for Linux 32bit](https://img.shields.io/badge/Linux-32bit-red)](https://google.com)
[![Download for Linux ARM](https://img.shields.io/badge/Linux-ARM-red)](https://google.com)


### **Getting Started (Windows)**:
1. Download the correct build for your operating system.
2. Launch Command prompt from the start-menu (search `cmd`)
3. Navigate to where the downloaded file is located (likely `cd Downloads`)
4. Run `lightstream-prism-dns-windows-386.exe` to launch the DNS server.
5. Windows may prompt you that it is blocking connections, make sure you choose to allow them.
6. Find the IP address of your computer in the output (typically, this looks like 192.168.xxx.xxx, 10.xxx.xxx.xxx, or 172.16.xxx.xxxx)
7. Enter the DNS server into your console and start broadcasting.

### **Getting started (MacOS / Linux)**:
1. Download the correct build for your operating system and architecture.
2. Launch your terminal application (search for Terminal on macOS)
3. Navigate to where the downloaded file is located (likely `cd Downloads`)
4. Run the executable you have downloaded (e.g `./lightstream-prism-dns-darwin-amd64 for macOS`)
5. You may be prompted to allow incoming connections, make sure to allow this.
6. Find the IP address of your computer in the output (typically, this looks like 192.168.xxx.xxx, 10.xxx.xxx.xxx, or 172.16.xxx.xxxx)
7. Enter the DNS server into your console and start broadcasting.

## Note:
This is a light fork of the [CoreDNS](https://coredns.io) project with small modifications to handle updating and auto-configuration.
