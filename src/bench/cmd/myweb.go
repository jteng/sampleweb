package main

import (
	"bench/engine"
	"bench/handler"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	mux := http.NewServeMux()

	handler := &handler.HelloWorld{
		Srvc: engine.NewService(),
	}
	mux.Handle("/", handler)

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Printf("listening on port 8080")
	go func() {
		log.Print(srv.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan
	log.Println("got interrupt signal, shutting down")
	ctx := context.Background()
	srv.Shutdown(ctx)
	log.Println("shut down completed")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
