# drs

Experimental tool to learn more about DNS resolution.
Please **DO NOT** use it.

## Features

- ✅ DNS lookup an address using a specific DNS server.

```sh
$ drs +short @1.1.1.1 example.com # Cloudflare DNS Server
$ drs +short @8.8.8.8 example.com # Google DNS Server
```

- ⏳ DNS lookup an address using the system configuration file (default).

```sh
$ drs +short example.com # Using system config file
```

## Getting Started

- Install the latest version of Go (at least >= 1.21).
- Build the tool as a binary executable.
```sh
$ make build
```
- Now, a `drs` binary should exist in the `bin` directory of this project.
- Have fun with `drs`!
