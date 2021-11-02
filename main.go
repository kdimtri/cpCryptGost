package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage:\nNo usage yet))")
	flag.PrintDefaults()
}

type InPut struct {
	Data string `json:"data"`
}
type OutPut struct {
	Data      string `json:"data"`
	Signature string `json:"signature"`
	Pubkey    string `json:"pubkey"`
}

func main() {
	var addr = flag.String("a", "localhost:8080", "default server host:port")
	var help = flag.Bool("h", false, "Show help message")
	var run = flag.Bool("s", false, "run http server")
	var data string
	flag.Parse()
	if *help {
		usage()
		os.Exit(0)
	}
	if *run {
		runHttpServer(*addr)
	} else {
		args := flag.Args()
		if len(args) < 1 {
			log.Println("No arguments")
			usage()
			os.Exit(1)
		}

		if len(args) > 1 {
			data = strings.Join(args[:], " ")
		} else {
			data = args[0]
		}
		request(data, *addr)
	}
	os.Exit(0)
}
