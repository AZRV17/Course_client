package app

import (
	"errors"
	"github.com/AZRV17/Course_client/internal/config"
	handlers "github.com/AZRV17/Course_client/internal/delivery/http"
	"github.com/AZRV17/Course_client/internal/repository"
	"github.com/AZRV17/Course_client/internal/server"
	"github.com/AZRV17/Course_client/internal/service"
	"github.com/AZRV17/Course_client/pkg/coin"
	"github.com/AZRV17/Course_client/pkg/storage/inmem"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

func Run() {
	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	storage := inmem.NewInMemoryCache()

	repo := repository.NewRepository(storage)
	serv := service.NewService(repo)
	handler := handlers.NewHandler(*serv)

	client := coin.NewClient(cfg.Client.Url, cfg.Client.Timeout)

	go func() {
		for {
			err := client.FetchData(repo)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Data updated successfully")
			time.Sleep(time.Minute * 10)
		}
	}()

	r := chi.NewRouter()

	handler.Init(r)

	httpSrv := server.NewHttpServer(cfg, r)

	stopped := make(chan struct{})

	go httpSrv.Shutdown(stopped)

	log.Printf("Starting HTTP server on %s\n", cfg.HTTP.Host+":"+cfg.HTTP.Port)

	go func() {
		if err := httpSrv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server ListenAndServe Error: %v", err)
		}
	}()

	<-stopped

	log.Println("HTTP server stopped")
}
