package utils

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

func StartServer(server *http.Server) {

	go func() {
		log.Println("Server is starting...")
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	log.Println("Server is shutting down...")
}
