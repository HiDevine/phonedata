package main

import (
	"flag"
	"log"
	"net/http"
)

/**
 * Author:  WangDepeng
 * Date :   2020-05-27 
 * Time :   09:57
 * Package: 
 * Mail :   wangdepeng@cmss.chinamobile.com
 * Project: phonedata
 * Description: 

 */
// A basic HTTP server.
// By default, it serves the current working directory on port 8080.

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}
