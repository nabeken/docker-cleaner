package main

import (
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-cleaner"
	app.Version = Version
	app.Usage = "A tool that remove orphaned volumes and obsoleted images from Docker host."
	app.Author = "TANABE Ken-ichi"
	app.Email = "nabeken@tknetworks.org"
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "directory, d",
			Value: "/var/lib/docker",
			Usage: "specify a docker directory",
		},
		cli.StringFlag{
			Name:  "endpoint, e",
			Value: "unix:///var/run/docker.sock",
			Usage: "specify a docker endpoint",
		},
	}

	app.Run(os.Args)
}

func joinDockerDir(c *cli.Context, dirs ...string) string {
	return filepath.Join(c.GlobalString("directory"), filepath.Join(dirs...))
}

func run(dryrun bool, dryRunF func(), runF func()) {
	if dryrun {
		dryRunF()
	} else {
		runF()
	}
}
