package main

import (
	"fmt"
	"log"
	"os"
	cut "task6/pkg"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "customcut",
			Usage: "complete a task on the list",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "fields", Aliases: []string{"f"}},
				&cli.StringFlag{Name: "delimiter", Aliases: []string{"d"}},
				&cli.BoolFlag{Name: "separated", Aliases: []string{"s"}},
			},
			Action: func(c *cli.Context) error {
				cut := cut.Create(c.String("fields"), c.String("delimiter"), c.Bool("separated"))
				fmt.Println(cut)
				cut.Run()

				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
