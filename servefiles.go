package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func defaultEndpoint(port int) (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}
		for _, a := range addrs {
			addr := fmt.Sprint(a)
			if strings.HasPrefix(addr, "192.") || strings.HasPrefix(addr, "10.") || strings.HasPrefix(addr, "172.") {
				s := strings.Split(addr, "/")
				return fmt.Sprintf("%s:%d", s[0], port), nil
			}
		}
	}
	return fmt.Sprintf("127.0.0.1:%d", port), nil
}

func main() {
	pPort := flag.Int("p", 8080, "port to listen on")
	flag.Parse()
	endpoint, err := defaultEndpoint(*pPort)
	if err != nil {
		fmt.Printf("*** %v\n", err)
		os.Exit(1)
	}
	flag.Usage = func() {
		fmt.Println("Usage: servefiles [-p port] directory")
	}
	args := flag.Args()
	if len(args) > 0 {
		fmt.Printf("Serving files from %s at http://%s/\n", args[0], endpoint)
		err = http.ListenAndServe(endpoint, http.FileServer(http.Dir(args[0])))
		if err != nil {
			fmt.Printf("*** %v\n", err)
			os.Exit(1)
		}
	} else {
		flag.Usage()
	}
}
