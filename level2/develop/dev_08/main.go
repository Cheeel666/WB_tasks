package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

func checkFlag(args []string) bool {
	for _, v := range args {
		if v == "-u" {
			return true
		}
	}
	return false
}

func netcat(args []string) error {
	prot := "tcp"
	udp := checkFlag(args)
	if udp {
		prot = "udp"
	}
	host := args[0]
	port := args[1]

	conn, err := net.Dial(prot, net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal(err)
	}
	osSignals := make(chan os.Signal, 1)
	listenErr := make(chan error, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	go request(conn, listenErr, osSignals)
	go response(conn, listenErr, osSignals)

	select {
	case <-osSignals:
		conn.Close()
	case err = <-listenErr:
		if err != nil {
			log.Fatalf("go-telnet: %s", err)
		}
	}

	return nil
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	path, _ := filepath.Abs(".")
	fmt.Print(path, " > ")

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		switch command[0] {
		case "pwd":
			fmt.Println(path)
		case "cd":
			err := os.Chdir(command[1])
			if err != nil {
				fmt.Println("Incorrect path")
			}
		case "echo":
			for i := 1; i < len(command); i++ {
				fmt.Fprint(os.Stdout, command[i], " ")
			}
			fmt.Println()
		case "nc": //usage: nc 192.168.88.1 80
			if len(command[1:]) < 2 {
				fmt.Println("Incorrect command!")
				fmt.Print(path, " > ")
				continue
			}
			err := netcat(command[1:])
			if err != nil {
				fmt.Println(err)
			}

		default:
			cmd := exec.Command(command[0], command[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
		}

		path, _ = filepath.Abs(".")
		fmt.Print(path, " > ")
	}
}

func request(conn net.Conn, listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			listenErr <- err
		}
		fmt.Fprintf(conn, text+"\n")
	}
}

func response(conn net.Conn, listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(conn)
		text, err := reader.ReadString('\n')
		if err != nil {
			listenErr <- err
		}
		fmt.Print(text)
	}
}

// app := &cli.App{}
// app.UseShortOptionHandling = true
// app.Commands = []*cli.Command{
// 	{
// 		Name:  "cd",
// 		Usage: "change directory",
// 		Action: func(c *cli.Context) error {
// 			fmt.Println("Here: ", c.Args().Slice()[0])
// 			err := os.Chdir(c.Args().Slice()[0])
// 			return err
// 		},
// 	},
// 	{
// 		Name:  "pwd",
// 		Usage: "print working directory",
// 		Action: func(c *cli.Context) error {
// 			path, err := filepath.Abs(".")
// 			fmt.Println(path)
// 			return err
// 		},
// 	},
// 	{
// 		Name:  "kill",
// 		Usage: "kill process",
// 		Action: func(c *cli.Context) error {
// 			// cut := .Create(c.String("fields"), c.String("delimiter"), c.Bool("separated"))
// 			// cut.Run(c.Args().Slice())
// 			return nil
// 		},
// 	},
// 	{
// 		Name:  "ps",
// 		Usage: "report a snapshot of current processes",
// 		// Flags: []cli.Flag{
// 		// 	&cli.StringFlag{Name: "fields", Aliases: []string{"f"}},
// 		// 	&cli.StringFlag{Name: "delimiter", Aliases: []string{"d"}},
// 		// 	&cli.BoolFlag{Name: "separated", Aliases: []string{"s"}},
// 		// },
// 		Action: func(c *cli.Context) error {
// 			// cut := .Create(c.String("fields"), c.String("delimiter"), c.Bool("separated"))
// 			// cut.Run(c.Args().Slice())
// 			return nil
// 		},
// 	},
// 	{
// 		Name:  "cd",
// 		Usage: "change directory",
// 		// Flags: []cli.Flag{
// 		// 	&cli.StringFlag{Name: "fields", Aliases: []string{"f"}},
// 		// 	&cli.StringFlag{Name: "delimiter", Aliases: []string{"d"}},
// 		// 	&cli.BoolFlag{Name: "separated", Aliases: []string{"s"}},
// 		// },
// 		Action: func(c *cli.Context) error {
// 			// cut := .Create(c.String("fields"), c.String("delimiter"), c.Bool("separated"))
// 			// cut.Run(c.Args().Slice())
// 			return nil
// 		},
// 	},
// }

// err := app.Run(os.Args)
// if err != nil {
// 	log.Fatal(err)
// }}
