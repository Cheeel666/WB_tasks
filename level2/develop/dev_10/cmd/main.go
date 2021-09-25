package main

import (
	"fmt"
	"log"
	"os"
	telnet "task10/pkg"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "telnet",
			Usage: "complete a task on the list",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "timeout", Aliases: []string{"t"}},
			},
			Action: func(c *cli.Context) error {
				telnet := telnet.InitFlags(c.Args().Slice()[0], c.Args().Slice()[1], c.Int("t"))
				// fmt.Println(telnet)
				err := telnet.Run()
				if err != nil {
					fmt.Println(err)
				}
				return err
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
