package task10

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Flags implements telnet flags
type Flags struct {
	host    string
	port    string
	timeout time.Duration
}

// InitFlags creates telnet struct
func InitFlags(host, port string, timeout int) *Flags {
	return &Flags{
		host:    host,
		port:    port,
		timeout: time.Duration(timeout),
	}
}

// Run runs telnet
func (t *Flags) Run() error {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(t.host, t.port), t.timeout)
	if err != nil {
		time.Sleep(t.timeout)
		return err
	}
	osSignals := make(chan os.Signal, 1)
	listenErr := make(chan error, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
	go req(conn, listenErr, osSignals)
	go resp(conn, listenErr, osSignals)

	select {
	case <-osSignals:
		conn.Close()
	case err = <-listenErr:
		if err != nil {
			return err
		}
	}
	return nil
}

func req(conn net.Conn, listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				osSignals <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			listenErr <- err
		}

		fmt.Fprintf(conn, text+"\n")
	}
}

func resp(conn net.Conn, listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(conn)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				osSignals <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			listenErr <- err
		}

		fmt.Print(text)
	}
}

// go-telnet -t 3 1.1.1.1 123
