package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/J4yTr1n1ty/DocuSeal-Discord-Redirector/pkg/api"
	"github.com/J4yTr1n1ty/DocuSeal-Discord-Redirector/pkg/config"
)

func main() {
	config.LoadConfig()

	r := api.InitRouter()

	srv := &http.Server{
		Addr:    ":" + config.Config.Port,
		Handler: r,
	}

	go func() {
		log.Printf("Starting server on port %s", config.Config.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
