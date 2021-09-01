package main

import (
	"fmt"
	"log"
	"os"

	sortCustom "utils/task3"

	"github.com/urfave/cli/v2"
)

func sortUsage() {
	log.Printf("Usage: ./cmd sort [OPTION]... [FILE]... \n")
}

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "sort",
			Usage: "complete a task on the list",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "key", Aliases: []string{"k"}},
				&cli.BoolFlag{Name: "number", Aliases: []string{"n"}},
				&cli.BoolFlag{Name: "reverse", Aliases: []string{"r"}},
				&cli.BoolFlag{Name: "unique", Aliases: []string{"u"}},
			},
			Action: func(c *cli.Context) error {
				if len(c.Args().Slice()) < 1 {
					sortUsage()
					os.Exit(1)
				}
				sortObject := sortCustom.Create(c.String("key"), c.Bool("number"), c.Bool("reverse"), c.Bool("unique"), c.Args().Slice())

				err := sortObject.Run()
				if err != nil {
					fmt.Println("Error in sort:", err)
				}
				// fmt.Println("before output")
				return sortObject.Output()
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
