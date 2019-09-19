package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	serverMode bool
	port       string
	host       string
	username   string

	defaultLocation = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func init() {
	flag.BoolVar(&serverMode, "s", false, "run as a server")
	flag.StringVar(&host, "h", "localhost:8080", "host of server")
	flag.StringVar(&username, "u", "anonymous", "username of client")
	flag.Parse()
}

const dateFormat = "2006/01/02 15:04:05"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if serverMode {
		s := NewServer(host)
		if err := s.Run(ctx); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	c, err := NewClient(host, username)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	if err := c.Run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
