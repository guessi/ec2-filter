package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Usage = "filter ec2 instances with tag"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "region",
			Aliases:  []string{"r"},
			Value:    "us-east-1",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "key",
			Aliases:  []string{"k"},
			Value:    "",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "value",
			Aliases:  []string{"v"},
			Value:    "",
			Required: false,
		},
	}

	app.Action = func(c *cli.Context) error {
		r := c.String("region")
		k := c.String("key")
		v := c.String("value")

		if instances := describeInstances(r, k, v); instances != nil {
			buildOutput(instances)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
