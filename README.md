# Lightstream Console DNS
This project enables you to run a local version of the Lightstream Console DNS servers we provide.

## Don't want to run your own?
Visit [psdnscheck.golightstream.com](https://psdnscheck.golightstream.com/) to identify the closest DNS server that we host to you.


## Self-hosting
### **Downloads**:

[![Download for Windows](https://img.shields.io/badge/Windows-64bit-blue)](https://github.com/golightstream/lightstream-prism-dns/releases/download/v2.1.0/lightstream-prism-dns-windows-amd64.exe)
[![Download for MacOS Intel](https://img.shields.io/badge/Mac-Intel-orange)](https://github.com/golightstream/lightstream-prism-dns/releases/download/v2.1.0/lightstream-prism-dns-darwin-amd64)
[![Download for MacOS M1](https://img.shields.io/badge/Mac-M1-orange)](https://github.com/golightstream/lightstream-prism-dns/releases/download/v2.1.0/lightstream-prism-dns-darwin-arm64)
[![Download for Linux ARM](https://img.shields.io/badge/Linux-64bit-red)](https://github.com/golightstream/lightstream-prism-dns/releases/download/v2.1.0/lightstream-prism-dns-linux-amd64)
[![Download for Linux 32bit](https://img.shields.io/badge/Linux-32bit-red)](https://github.com/golightstream/lightstream-prism-dns/releases/download/v2.1.0/lightstream-prism-dns-linux-386)


### **Getting Started (Windows)**:
1. Download the correct build for your operating system.
2. Launch Command prompt from the start-menu (search for `cmd`)
3. Navigate to where the downloaded file is located e.g `cd Downloads`
4. Run `lightstream-prism-dns-windows-amd64.exe` to launch the DNS server.
5. Windows may prompt you that it is blocking connections, make sure you choose to allow them.
6. Find the IP address of your computer in the output (typically, this looks like 192.168.xxx.xxx, 10.xxx.xxx.xxx, or 172.16.xxx.xxxx)
7. Enter the IP address of your computer into your console's DNS settings and start broadcasting!

<div style="max-width: 500px; padding: 30px 0px;">

![Example](https://surl.im/i/m7z9glyax5cwrcrgcwvy62ln3jwd7gvq.png)
</div>

### **Getting started (MacOS / Linux)**:
1. Download the correct build for your operating system and architecture.
2. Launch your terminal application (search for Terminal on macOS)
3. Navigate to where the downloaded file is located e.g `cd Downloads`
4. Make the downloaded binary runnable:
    1. **Linux and Mac**: make the binary executable by running: `chmod +x ./lightstream-prism-dns-*`
    2. **Mac**: allow the binary to run: `sudo xattr -rd com.apple.quarantine ./lightstream-prism-dns-*`
4. Run the executable you have downloaded: `./lightstream-prism-dns-*`
5. You may be prompted to allow incoming connections, make sure to allow this.
6. Find the IP address of your computer in the output (typically, this looks like 192.168.xxx.xxx, 10.xxx.xxx.xxx, or 172.16.xxx.xxxx)
7. Enter the IP address of your computer into your console's DNS settings and start broadcasting!


### Building from source
It's possible to build this project from the source in this repository if you desire:
1. Clone the repository to your local machine: `git clone https://github.com/golightstream/lightstream-prism-dns.git`
2. Install and configure Go 1.20+: [go.dev](https://go.dev/dl/)
3. run `make` to build and compile
4. a binary `coredns` will be output, run `./coredns` to execute it.

## Note:
This is a light fork of the [CoreDNS](https://coredns.io) project with small modifications to handle updating and auto-configuration.

