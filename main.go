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

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func registerOldHandlers() {
	db := connection()
	defer db.Close()
	a, _ := findAll(db)

	for i := range a {
		handler := APIHandlerStr{a[i]}
		handle := fmt.Sprintf("%s", a[i])
		http.Handle(handle, &handler)
	}
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)

	factory := HandlerFactory{""}
	http.Handle("/mock", &factory)
	http.HandleFunc("/", RootHandler)

	go registerOldHandlers()
	log.Println("Starting web server at", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal("ERROR:", err)
	}

}
