package main

import (
	"fmt"
	"log"
	"os"

	sortCustom "utils/dev_03"
	grepCustom "utils/pkg/dev05"

	// cutCustom "utils/dev_06/pkg"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "customsort",
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
		// {
		// 	Name:  "customcut",
		// 	Usage: "complete a task on the list",
		// 	Flags: []cli.Flag{
		// 		&cli.StringFlag{Name: "fields", Aliases: []string{"f"}},
		// 		&cli.StringFlag{Name: "delimiter", Aliases: []string{"d"}},
		// 		&cli.BoolFlag{Name: "separated", Aliases: []string{"s"}},
		// 	},
		// 	Action: func(c *cli.Context) error {
		// 		cut := cutCustom.Create(c.String("fields"), c.String("delimiter"), c.Bool("separated"))
		// 		fmt.Println(cut)
		// 		cut.Run()

		// 		return nil
		// 	},
		// },
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
