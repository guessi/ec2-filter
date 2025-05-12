package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "boom",
		Usage: "filter ec2 instances with tag",
		Flags: []cli.Flag{
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
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			r := c.String("region")
			k := c.String("key")
			v := c.String("value")

			if instances := describeInstances(r, k, v); instances != nil {
				if len(instances.Reservations) == 0 {
					fmt.Println("No instances found")
					return nil
				}
				buildOutput(instances)
			}

			return nil

		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
