# servefiles

A trivial static file HTTP server, written in Go

## Build/install

    $ GOPATH=$PWD go get github.com/boynton/servefiles
    $ GOPATH=$PWD go install github.com/boynton/servefiles

## Usage

    $ ./bin/servefiles
    Usage: servefiles [-p <port>] directory
    $ ./bin/servefiles .
    Serving files from . at http://localhost:8080/
    ^C
    $ ./bin/servefiles -p 8000 .
    Serving files from . at http://localhost:8000/

