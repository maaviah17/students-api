package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/maaviah17/students-api/internal/config"
)

func main(){
	
	//initial setups :
		//load config
		//db setup
		//setup router
		//setup server

	// loading config
	cfg := config.MustLoad()


	//database setup

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Students API"))
	})

	//setup server
	server := http.Server{
		Addr: cfg.Address,
		Handler: router,
	}

	fmt.Printf("Server started :) %s", cfg.HTTPServer.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	go func(){
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to start server ;( : %v", err)
		}
	}()

	<-done

	slog.Info("Shutting down the server")

}
