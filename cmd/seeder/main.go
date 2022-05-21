package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/ghazlabs/es-starterkit/internal/kuliner-api/core"
	"github.com/ghazlabs/es-starterkit/internal/kuliner-api/driven/storage"
)

const (
	envKeyESHost   = "ES_HOST"
	envKeyDataPath = "DATA_PATH"
)

func main() {
	// initialize es client
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{os.Getenv(envKeyESHost)},
	})
	if err != nil {
		log.Fatalf("unable to initialize es client due: %v", err)
	}
	// initialize mappings
	err = initMappings(esClient)
	if err != nil {
		log.Fatalf("unable to initialize mappings due: %v", err)
	}
	// initialize storage for seed data operation
	strg, err := storage.New(storage.Config{
		ESClient:    esClient,
		ESIndexName: "foods",
	})
	if err != nil {
		log.Fatalf("unable to initialize storage due: %v", err)
	}
	// seed data
	err = seedData(strg, os.Getenv(envKeyDataPath))
	if err != nil {
		log.Fatalf("unable to seed data due: %v", err)
	}
}

func initMappings(esClient *elasticsearch.Client) error {
	// initialize `foods` index mapping as written in es_mappings.md
	indexSettings := `
		{
			"mappings": {
				"properties": {
					"name": {
						"type": "text"
					},
					"description": {
						"type": "text"
					}
				},
				"dynamic": false
			}
		}   
	`
	resp, err := esClient.Indices.Create(
		"foods",
		esClient.Indices.Create.WithBody(strings.NewReader(indexSettings)),
	)
	if err != nil {
		return fmt.Errorf("unable to create index due: %w", err)
	}
	defer resp.Body.Close()

	return nil
}

func seedData(strg *storage.Storage, filePath string) error {
	// load file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("unable to read file due: %w", err)
	}
	// parse data
	var foodInputs []core.FoodInput
	err = json.Unmarshal(data, &foodInputs)
	if err != nil {
		return fmt.Errorf("unable to parse data due: %w", err)
	}
	// insert data
	ctx := context.Background()
	for _, foodInput := range foodInputs {
		err = strg.IndexFood(ctx, foodInput.ToFood())
		if err != nil {
			return fmt.Errorf("unable to input food due: %w", err)
		}
	}
	return nil
}
