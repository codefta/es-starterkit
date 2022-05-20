package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/ghazlabs/es-starterkit/internal/kuliner-api/core"
	"github.com/ghazlabs/es-starterkit/internal/kuliner-api/driven/storage"
	"github.com/ghazlabs/es-starterkit/internal/kuliner-api/driver/rest"
)

const (
	envKeyESHost = "ES_HOST"
	listenAddr   = ":8101"
)

func main() {
	// initialize es client
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{os.Getenv(envKeyESHost)},
	})
	if err != nil {
		log.Fatalf("unable to initialize es client due: %v", err)
	}
	// initialize storage
	strg, err := storage.New(storage.Config{
		ESClient:    esClient,
		ESIndexName: "foods",
	})
	if err != nil {
		log.Fatalf("unable to initialize storage due: %v", err)
	}
	// initialize service
	svc, err := core.NewService(core.Config{
		Storage:     strg,
		SearchLimit: 10,
	})
	if err != nil {
		log.Fatalf("unable to initialize service due: %v", err)
	}
	// initialize api
	api, err := rest.NewAPI(rest.Config{Service: svc})
	if err != nil {
		log.Fatalf("unable to initialize api due: %v", err)
	}
	// initialize server
	server := &http.Server{
		Addr:        listenAddr,
		Handler:     api.GetHandler(),
		ReadTimeout: 3 * time.Second,
	}
	// run server
	log.Printf("server is listening on %v", listenAddr)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("unable to run server due: %v", err)
	}
}
