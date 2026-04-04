//go:build server
// +build server

package main

import (
	"flag"
	"log"

	am "github.com/adaptivemsg/adaptivemsg-go"
	_ "github.com/adaptivemsg/hello-server-go/api/hello"
)

func main() {
	addr := flag.String("addr", "tcp://127.0.0.1:5555", "listen address (examples: tcp://127.0.0.1:5555, uds://@adaptivemsg-hello, uds:///tmp/adaptivemsg-hello.sock)")
	flag.Parse()

	log.Printf("hello server listening on %s", *addr)
	if err := am.NewServer().Serve(*addr); err != nil {
		log.Fatal(err)
	}
}
