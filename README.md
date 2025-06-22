[![Readme Card](https://github-readme-stats.vercel.app/api/pin/?username=cyclone-github&repo=container_truncator&theme=gruvbox)](https://github.com/cyclone-github/container_truncator/)

[![Go Report Card](https://goreportcard.com/badge/github.com/cyclone-github/container_truncator)](https://goreportcard.com/report/github.com/cyclone-github/container_truncator)
[![GitHub issues](https://img.shields.io/github/issues/cyclone-github/container_truncator.svg)](https://github.com/cyclone-github/container_truncator/issues)
[![License](https://img.shields.io/github/license/cyclone-github/container_truncator.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/cyclone-github/container_truncator.svg)](https://github.com/cyclone-github/container_truncator/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/cyclone-github/container_truncator.svg)](https://pkg.go.dev/github.com/cyclone-github/container_truncator)

### Cyclone's Truecrypt / Veracrypt Container Truncater
- Tool to process Truecrypt / Veracrypt container files to be ran with hashcat
- Tested on debian linux & Windows 11
- Run tool in directory where Truecrypt *.tc or Veracrypt *.vc container file is located
- Once file is selected from menu, tool will truncate container and save a new file to "truncate_filename"
- "truncate_filename" can now be ran with hashcat using the appropriate mode
  - example: ./veracrypt_container_truncator.bin
  - 1 ) truecrypt.tc
  - outputs new file "truncate_truecrypt.tc"
  - hashcat -m 6211 -a 0 truncate_truecrypt.tc cyclone_hk_v2.txt -r cyclone_250.rule
- Tool will not overwrite any existing files

### Compile from source:
- If you want the latest features, compiling from source is the best option since the release version may run several revisions behind the source code.
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/container_truncator.git`  # clone repo
  - `cd container_truncator`                                               # enter project directory
  - `go mod init container_truncator`                                      # initialize Go module (skips if go.mod exists)
  - `go mod tidy`                                              # download dependencies
  - `go build -ldflags="-s -w" .`                              # compile binary in current directory
  - `go install -ldflags="-s -w" .`                            # compile binary and install to $GOPATH
- Compile from source code how-to:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt