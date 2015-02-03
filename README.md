# servefiles

A trivial static file HTTP server, written in Go

## Build/install

    $ go get github.com/boynton/servefiles

## Usage

    $ $GOPATH/bin/servefiles
    Usage: servefiles [-e [host]:port] directory
    $ $GOPATH/bin/servefiles .
    Serving files from . at http://localhost:8080/
    ^C
    $ $GOPATH/bin/servefiles -e :8000 .
    Serving files from . at http://0.0.0.0:8000/
    $ $GOPATH/bin/servefiles -e myhostname:8000 http_root
    Serving files from http_root at http://myhostname:8000/

