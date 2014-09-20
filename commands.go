package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandVolume,
}

var commandVolume = cli.Command{
	Name:      "volume",
	ShortName: "v",
	Usage:     "Removes orphaned volumes from Docker host",
	Action:    doVolume,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "dry-run, n",
			Usage: "show what volumes will be deleted",
		},
	},
}
