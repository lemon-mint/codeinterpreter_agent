package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ln, err := net.Listen("tcp", os.Getenv("ADDR")+":"+port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/writefile", WriteFileHandler)
	mux.HandleFunc("/runcommand", RunCommandHandler)
	go http.Serve(ln, mux)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
}
