# docker-volume-cleanup [![wercker status](https://app.wercker.com/status/4efcb9ae1b9d0f7d7c2335c4342afec7/s/master "wercker status")](https://app.wercker.com/project/bykey/4efcb9ae1b9d0f7d7c2335c4342afec7)

A tool that removes orphaned volumes from Docker host.

# Installation (from source)

```sh
$ go get -u github.com/nabeken/docker-volume-cleanup
```

# Usage

```sh
Usage of docker-volume-cleanup:
  -d="/var/lib/docker": Specify docker directory
  -h="unix:///var/run/docker.sock": Specify a docker endpoint
  -n=true: Dry-run. Do not delete volumes by default. Use -n=false to delete volumes actually.
```

# Author

TANABE Ken-ichi

# LICENSE

See [LICENSE](LICENSE).

# LICENSE for binary distribution

The binary form distribution of `docker-volume-cleanup` contains the following products. See individual licenses:

- [The Go's runtime](http://golang.org/LICENSE)
- [go-dockerclient](https://github.com/fsouza/go-dockerclient/LICENSE)
