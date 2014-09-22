package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandVolume,
	commandImage,
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

var commandImage = cli.Command{
	Name:      "image",
	ShortName: "i",
	Usage:     "Removes orphaned images from Docker host",
	Action:    doImage,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "dry-run, n",
			Usage: "show what images will be deleted",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "specify a image name",
		},
		cli.IntFlag{
			Name:  "duration, d",
			Usage: "delete images whose Created is passed for a specified duration in seconds",
		},
	},
}
