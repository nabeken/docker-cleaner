package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinDockerDir(t *testing.T) {
	assert.Equal(t, "/var/lib/docker/vfs/dir", JoinDockerDir("vfs", "dir"))
}
