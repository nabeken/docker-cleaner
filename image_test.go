package main

import (
	"testing"
	"time"

	"github.com/fsouza/go-dockerclient"
	"github.com/stretchr/testify/assert"
)

var testNow = time.Now()

var testImages = images{
	docker.APIImages{
		RepoTags: []string{"test:latest"},
		Created:  testNow.Unix(),
	},
	docker.APIImages{
		RepoTags: []string{"test:2.0"},
		Created:  testNow.Add(-time.Hour * 2).Unix(),
	},
	docker.APIImages{
		RepoTags: []string{"test:1.0"},
		Created:  testNow.Add(-time.Hour * 3).Unix(),
	},
	docker.APIImages{
		RepoTags: []string{"hoge", "hoge:latest"},
		Created:  testNow.Add(-time.Hour * 2).Unix(),
	},
}

func TestFilterByName(t *testing.T) {
	assert.Len(t, testImages.Filter(filterByName("FOOBAR")), 0)
	assert.Len(t, testImages.Filter(filterByName("test")), 3)
	assert.Len(t, testImages.Filter(filterByName("hoge:latest")), 1)
}

func TestFilterByCreatedAt(t *testing.T) {
	assert.Len(t, testImages.Filter(filterByCreatedAt(0)), 4)
	assert.Len(t, testImages.Filter(filterByCreatedAt(int(time.Hour*1/time.Second))), 3)
}

func TestFilter(t *testing.T) {
	assert.Len(t, testImages.
		Filter(filterByName("hoge")).
		Filter(filterByCreatedAt(int(time.Hour*3/time.Second))), 0)

	{
		image := testImages.
			Filter(filterByName("test")).
			Filter(filterByCreatedAt(int(time.Hour * 3 / time.Second)))
		assert.Len(t, image, 1)
		assert.Contains(t, image[0].RepoTags, "test:1.0")
	}
}
