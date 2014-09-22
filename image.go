package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
)

type images []docker.APIImages

type filter func(image docker.APIImages) bool

func (i images) Filter(f filter) images {
	ret := images{}
	for _, image := range i {
		if f(image) {
			ret = append(ret, image)
		}
	}
	return ret
}

func listImages(client *docker.Client) (images, error) {
	images := images{}
	apiImages, err := client.ListImages(false)
	if err != nil {
		return nil, err
	}
	for i := range apiImages {
		images = append(images, apiImages[i])
	}
	return apiImages, nil
}

func filterByName(name string) func(image docker.APIImages) bool {
	return func(image docker.APIImages) bool {
		for i := range image.RepoTags {
			if strings.HasPrefix(image.RepoTags[i], name) {
				return true
			}
		}
		return false
	}
}

func filterByCreatedAt(duration int) func(image docker.APIImages) bool {
	return func(image docker.APIImages) bool {
		d := time.Second * time.Duration(duration)
		return time.Since(time.Unix(image.Created, 0)) > d
	}
}

func doImage(c *cli.Context) {
	client, err := docker.NewClient(c.GlobalString("endpoint"))
	if err != nil {
		log.Fatal(err)
	}

	images, err := listImages(client)
	if err != nil {
		log.Fatal(err)
	}

	ret := images.
		Filter(filterByName(c.String("name"))).
		Filter(filterByCreatedAt(c.Int("duration")))
	for i := range ret {
		var err error
		run(!c.Bool("force"),
			func() {
				fmt.Println("dryrun: removed:", ret[i].ID, ret[i].RepoTags)
			},
			func() {
				err = client.RemoveImage(ret[i].ID)
				fmt.Println("removed:", ret[i].ID, ret[i].RepoTags)
			},
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: failed to delete a image %v %s", err, ret[i].ID, ret[i].RepoTags)
		}
	}
}
