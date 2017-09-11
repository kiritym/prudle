package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

func clearDBBeforeTerminates() {
	var err = os.Remove("http.db")
	if isError(err) {
		return
	}
}
func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Println(addr)

	factory := HandlerFactory{""}
	http.Handle("/mock", &factory)
	http.HandleFunc("/", RootHandler)
	//clear the DB before program ends
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		clearDBBeforeTerminates()
		os.Exit(0)
	}()

	log.Println("Starting web server at", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
