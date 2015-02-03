package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: servefiles [-e [host]:port] directory")
	}
	endpoint := flag.String("e", "localhost:8080", "endpoint ([host]:port")
	flag.Parse()
	args := flag.Args()
	if len(args) == 1 && len(*endpoint) > 0 {
		s := *endpoint
		if s[0] == ':' {
			s = "0.0.0.0" + s
		}
		fmt.Printf("Serving files from %s at http://%s/\n", args[0], s)
		err := http.ListenAndServe(*endpoint, http.FileServer(http.Dir(args[0])))
		if err != nil {
			fmt.Printf("*** %v\n", err)
			os.Exit(1)
		}
	}
	flag.Usage()
}
