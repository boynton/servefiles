package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: servefiles [-p <port>] directory")
	}
	var port = flag.Int("p", 8080, "port number")
	flag.Parse()
	args := flag.Args()
	if len(args) == 1 {
		panic(http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir(args[0]))))
	}
	flag.Usage()
}
