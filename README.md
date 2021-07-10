# Prototype-Polluter

Take a list of domains and probe for working http and https servers.

## Install

```
▶ go get github.com/Piyush-Paliwal/prototype-polluter
```
## Basic Usage

httprobe accepts line-delimited domains on `stdin`:

```
▶ cat recon/example/domains.txt
example.com
example.edu
example.net
▶ cat recon/example/domains.txt | httprobe
http://example.com
http://example.net
http://example.edu
https://example.com
https://example.edu
https://example.net
```
