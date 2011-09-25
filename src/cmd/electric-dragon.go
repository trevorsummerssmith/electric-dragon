package main

import (
	"electric-dragon/web"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	port = flag.String("p", "8080", "web listen address")
	ip = flag.String("l", "127.0.0.1", "address to listen on")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	// Get listening socket
	var wsock net.Listener
	var err os.Error
	var laddr = *ip + ":" + *port
	wsock, err = net.Listen("tcp", laddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on %s...", laddr)
	web.Serve(wsock)
}