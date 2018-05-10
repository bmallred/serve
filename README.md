# A simple file server over HTTP

Just a simple file server over HTTP useful for when you would like to quickly share the contents of a directory with others or to simply test some HTML/Javascript/CSS out in the browser.

This program was converted to **literate programming** with the help of [lmt](https://github.com/driusan/lmt) by **Dave MacFarlane**.

## How to get it

Go ahead and pull the latest from Github

``` shell
go get -u github.com/bmallred/serve
```

To generate the latest code from this file run

``` shell
lmt README.md
```

and then build normally with Go

``` shell
go build
```

## Basic structure

The main code block will have the following basic structure.

``` go main.go
package main

import (
	<<<main.go imports>>>
)

func main() {
	<<<main implementation>>>
}
```

The logic to be implemented is pretty straight forward.

 1. Find the directory to serve from
 2. Find the address and port to bind to
 3. Listen and serve
 
The base directory for now will always be the current working directory to keep it simple. However someone may want to bind to a different address and/or port combination. The default port will be `8080` on all publicly accessible addresses.
 
``` go "main implementation"
<<<base directory>>>
<<<address and port>>>
<<<listen and serve>>>
```

Currently only a few packages need to be imported from the Go standard library.

``` go "main.go imports"
"log"
"net/http"
"os"
"path/filepath"
```

## Determine the base directory

The files served will always be from the current working directory. So if executed at `/tmp` then all directories and files will be accessible from the address being served.

``` go "base directory"
d, err := filepath.Abs(".")
if err != nil {
	log.Fatal(err)
}
```

## Bind to an address and port

By default the program will bind to `localhost:8080`. However, if this is not desired simply pass the address and port combination as the first argument to the command. For example:

``` shell
serve 0.0.0.0:8081
```

This will bind to the address `0.0.0.0` on port `8081`.

``` go "address and port"
addr := "localhost:8080"
if len(os.Args) > 1 {
	addr := os.Args[1]
	if addr == "" {
		addr = ":8000"
	}
}
```

## The HTTP server

A basic implementation of the file server would result in the following snippet.

``` go "listen and serve"
log.Println("Server running at " + addr)
log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir(d))))
```

## Logging of requests

It would be nice to know what files are being accessed from the server. In order to accomplish this we can inject, or wrap, our logic for serving files with another function to log each incoming request.

``` go main.go +=

// Logger middleware for HTTP handling.
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] [%s] %s", r.RemoteAddr, r.Method, r.URL.String())
		h.ServeHTTP(w, r)
	})
}
```

And now we want to inject the request logging around the file server.

``` go "listen and serve"
log.Println("Server running at " + addr)
log.Fatal(http.ListenAndServe(addr, Logger(http.FileServer(http.Dir(d)))))
```
