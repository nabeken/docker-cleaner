# docker-cleaner [![wercker status](https://app.wercker.com/status/c828dfeb0d38bac87f4eb1e3f73c1387/s "wercker status")](https://app.wercker.com/project/bykey/c828dfeb0d38bac87f4eb1e3f73c1387)

A tool that removes orphaned volumes from Docker host. It's useful for testing and development.

*I DO NOT RECOMMEND TO RUN THIS COMMAND ON YOUR PRODUCTION ENVIRONMENT.*

# Installation (from Github releases)

Download from [releases](https://github.com/nabeken/docker-cleaner/releases).

# Installation (from source)

```sh
$ go get -u github.com/nabeken/docker-cleaner
```

# Usage

```sh
$ docker-cleaner help
NAME:
   docker-cleaner - A tool that remove orphaned volumes and obsoleted images from Docker host.

USAGE:
   docker-cleaner [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
  TANABE Ken-ichi - <nabeken@tknetworks.org>

COMMANDS:
   volume, v	Removes orphaned volumes from Docker host
   image, i	Removes orphaned images from Docker host
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --directory, -d '/var/lib/docker'		specify a docker directory
   --endpoint, -e 'unix:///var/run/docker.sock'	specify a docker endpoint
   --help, -h					show help
   --version, -v				print the version
```

# Author

TANABE Ken-ichi

# LICENSE

See [LICENSE](LICENSE).

# LICENSE for binary distribution

The binary form distribution of `docker-cleaner` contains the following products. See individual licenses:

- [The Go's runtime](http://golang.org/LICENSE)
- [go-dockerclient](https://github.com/fsouza/go-dockerclient/LICENSE)
- [cli.go](https://github.com/codegangsta/cli/LICENSE)
