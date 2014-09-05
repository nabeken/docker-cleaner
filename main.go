package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

var (
	dockerDir = flag.String("d", "/var/lib/docker", "Specify docker directory")
	endPoint  = flag.String("h", "unix:///var/run/docker.sock", "Specify a docker endpoint")
	dryrun    = flag.Bool("n", true, "Dry-run. Do not delete volumes by default. Use -n=false to delete volumes actually.")
)

func JoinDockerDir(dirs ...string) string {
	return *dockerDir + "/" + strings.Join(dirs, string(os.PathSeparator))
}

func DeleteVolume(volumeId string) {
	var err error
	err = DoDelete(JoinDockerDir("volumes", volumeId))
	err = DoDelete(JoinDockerDir("vfs", "dir", volumeId))
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to delete volume id", volumeId)
	}
}

func ListOndiskVolumes(volumesDir string) (map[string]bool, error) {
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

func ListOnContainerVolumes(client *docker.Client) (map[string]bool, error) {
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

func DoDelete(path string) error {
	if *dryrun {
		fmt.Println("dryrun:", path)
		return nil
	}
	fmt.Println("remove:", path)
	return os.RemoveAll(path)
}

func main() {
	flag.Parse()

	ondiskVolumes, err := ListOndiskVolumes(JoinDockerDir("volumes"))
	if err != nil {
		log.Fatal(err)
	}

	client, err := docker.NewClient(*endPoint)
	if err != nil {
		log.Fatal(err)
	}
	onContainerVolumes, err := ListOnContainerVolumes(client)
	if err != nil {
		log.Fatal(err)
	}

	for n := range onContainerVolumes {
		if ondiskVolumes[n] {
			delete(ondiskVolumes, n)
		}
	}

	for volume := range ondiskVolumes {
		DeleteVolume(volume)
	}
}
