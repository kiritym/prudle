package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	host = flag.String("b", "0.0.0.0", "Host of server")
	port = flag.Int("p", 8888, "Port of server")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Println(addr)

	factory := HandlerFactory{""}
	http.Handle("/mock", &factory)
	http.HandleFunc("/", RootHandler)

	log.Println("Starting web server at", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
