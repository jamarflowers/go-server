/*

Jamar Flowers
5/31/2018

DESC:
simple go server serving static html file
will update to serve web apps

*/



package main

import (
	"log"
	"net/http"
)

func main() {
	staticserv := http.FileServer(http.Dir("static"))
	http.Handle("/", staticserv)

	log.Println("file server on...")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

/*
    package main
import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}
*/
/