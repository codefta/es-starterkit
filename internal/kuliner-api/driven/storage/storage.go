package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/ghazlabs/es-starterkit/internal/kuliner-api/core"
	"gopkg.in/validator.v2"
)

type Storage struct {
	esClient        *elasticsearch.Client
	esUpdateRequest *esapi.UpdateRequest
	esIndexName     string
}

type Config struct {
	ESClient    *elasticsearch.Client `validate:"nonnil"`
	ESIndexName string                `validate:"nonzero"`
}

func New(cfg Config) (*Storage, error) {
	err := validator.Validate(cfg)
	if err != nil {
		return nil, err
	}
	strg := &Storage{esClient: cfg.ESClient, esIndexName: cfg.ESIndexName}
	return strg, nil
}

func (s *Storage) IndexFood(ctx context.Context, food core.Food) error {
	foodDoc := newFoodDoc(food)
	resp, err := s.esClient.Index(
		s.esIndexName,
		strings.NewReader(foodDoc.String()),
		s.esClient.Index.WithContext(ctx),
		s.esClient.Index.WithDocumentID(food.ID),
	)
	if err != nil {
		return fmt.Errorf("unable to index document due: %w", err)
	}
	defer resp.Body.Close()

	if resp.IsError() {
		data, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("unable to index document due: %s", data)
	}
	return nil
}

func (s *Storage) UpdateFood(ctx context.Context, id string, food core.Food) error {
	foodDoc := newFoodDoc(food)
	req := esapi.UpdateRequest{
		Index:      s.esIndexName,
		DocumentID: id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, foodDoc.String()))),
	}

	resp, err := req.Do(ctx, s.esClient)
	if err != nil {
		return fmt.Errorf("unable to update document due: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return core.ErrNotFound
	}

	if resp.IsError() {
		data, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("unable to update document due: %s", data)
	}
	return nil
}

func (s *Storage) DeleteFood(ctx context.Context, id string) error {
	resp, err := s.esClient.Delete(
		s.esIndexName,
		id,
		s.esClient.Delete.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("unable to execute delete request due: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return core.ErrNotFound
	}
	if resp.IsError() {
		data, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("unable to execute delete due: %s", data)
	}

	return nil
}

func (s *Storage) SearchFood(ctx context.Context, query string, size int) ([]core.Food, error) {
	resp, err := s.esClient.Search(
		s.esClient.Search.WithContext(ctx),
		s.esClient.Search.WithQuery(query),
		s.esClient.Search.WithIndex(s.esIndexName),
		s.esClient.Search.WithSize(size),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to execute search request due: %w", err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	if resp.IsError() {
		return nil, fmt.Errorf("unable to search due: %s", data)
	}
	var r searchFoodsResult
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, fmt.Errorf("unable to parse search result due: %w", err)
	}
	return r.GetFoods(), nil
}
