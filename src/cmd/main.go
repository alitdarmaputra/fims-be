package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := InitializeServer()

	go func() {
		//Service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("listen and serve failed:", err.Error())
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 30 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Panicln("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("shutting down failed:", err.Error())
	}

	log.Println("server exiting")
}
