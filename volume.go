package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
)

func listOndiskVolumes(volumesDir string) (map[string]bool, error) {
	dirs, err := ioutil.ReadDir(volumesDir)
	if err != nil {
		return nil, err
	}

	ondiskVolumes := map[string]bool{}
	for _, dir := range dirs {
		if dir.IsDir() && len(dir.Name()) == 64 {
			ondiskVolumes[dir.Name()] = true
		}
	}
	return ondiskVolumes, nil
}

func listOnContainerVolumes(client *docker.Client) (map[string]bool, error) {
	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		return nil, err
	}

	onContainerVolumes := map[string]bool{}
	for _, container := range containers {
		ci, err := client.InspectContainer(container.ID)
		if err != nil {
			return nil, err
		}
		for _, volumeDir := range ci.Volumes {
			onContainerVolumes[filepath.Base(volumeDir)] = true
		}
	}
	return onContainerVolumes, nil
}

func doVolume(c *cli.Context) {
	ondiskVolumes, err := listOndiskVolumes(joinDockerDir(c, "volumes"))
	if err != nil {
		log.Fatal(err)
	}

	client, err := docker.NewClient(c.GlobalString("endpoint"))
	if err != nil {
		log.Fatal(err)
	}
	onContainerVolumes, err := listOnContainerVolumes(client)
	if err != nil {
		log.Fatal(err)
	}

	for n := range onContainerVolumes {
		if ondiskVolumes[n] {
			delete(ondiskVolumes, n)
		}
	}

	for volumeId := range ondiskVolumes {
		dirs := []string{
			joinDockerDir(c, "volumes", volumeId),
			joinDockerDir(c, "vfs", "dir", volumeId),
		}
		for _, dir := range dirs {
			var err error
			run(!c.Bool("force"),
				func() {
					fmt.Println("dryrun: removed:", dir)
				},
				func() {
					err = os.RemoveAll(dir)
					fmt.Println("removed:", dir)
				},
			)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: failed to delete volume id:%s", err, volumeId)
			}
		}
	}
}
