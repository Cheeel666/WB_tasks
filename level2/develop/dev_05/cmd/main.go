package main

import (
	"log"
	"os"
	grepCustom "task5/pkg"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "customgrep",
			Usage: "complete a task on the list",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "after", Aliases: []string{"A"}},
				&cli.IntFlag{Name: "before", Aliases: []string{"B"}},
				&cli.IntFlag{Name: "context", Aliases: []string{"C"}},
				&cli.BoolFlag{Name: "count", Aliases: []string{"c"}},
				&cli.BoolFlag{Name: "ignoreCase", Aliases: []string{"i"}},
				&cli.BoolFlag{Name: "invert", Aliases: []string{"v"}},
				&cli.BoolFlag{Name: "fixed", Aliases: []string{"F"}},
				&cli.BoolFlag{Name: "lineNum", Aliases: []string{"n"}},
			},
			Action: func(c *cli.Context) error {
				grep := grepCustom.Create(c.Int("after"), c.Int("before"), c.Int("context"),
					c.Bool("count"), c.Bool("ignoreCase"), c.Bool("invert"), c.Bool("fixed"),
					c.Bool("lineNum"), c.Args().Slice()[0], c.Args().Slice()[1:])
				grep.Run()

				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
