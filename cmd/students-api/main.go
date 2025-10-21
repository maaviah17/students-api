package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	slog.Info("server started", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	go func(){
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to start server ;( : %v", err)
		}
	}()

	<-done

	//graceful shutdown : 
	slog.Info("Shutting down the server")

	ctx,cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil{
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	
	slog.Info("server shutdown successful")
}
