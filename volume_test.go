package main

import (
	"flag"
	"testing"

	"github.com/codegangsta/cli"
	"github.com/stretchr/testify/assert"
)

func TestJoinDockerDir(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	set.String("directory", "/var/lib/docker", "test")
	c := cli.NewContext(nil, set, set)
	assert.Equal(t, "/var/lib/docker/vfs/dir", joinDockerDir(c, "vfs", "dir"))
}
