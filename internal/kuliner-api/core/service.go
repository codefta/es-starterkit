package core

import (
	"context"
	"fmt"
)

type Service interface {
	IndexFood(ctx context.Context, input FoodInput) (*Food, error)
	DeleteFood(ctx context.Context, id string) error
	SearchFoods(ctx context.Context, query string) ([]Food, error)
}

func NewService(s Storage) (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("missing storage")
	}
	return &service{storage: s}, nil
}

type service struct {
	storage Storage
}

func (s *service) IndexFood(ctx context.Context, input FoodInput) (*Food, error) {
	err := input.Validate()
	if err != nil {
		return nil, err
	}
	err = s.storage.IndexFood(ctx, input.ToFood())
	if err != nil {
		return nil, fmt.Errorf("unable to index food into storage due: %w", err)
	}
	return nil, nil
}

func (s *service) DeleteFood(ctx context.Context, id string) error {
	if len(id) == 0 {
		return ErrMissingID
	}
	err := s.storage.DeleteFood(ctx, id)
	if err != nil {
		return fmt.Errorf("unable to delete food from storage due: %w", err)
	}
	return nil
}

func (s *service) SearchFoods(ctx context.Context, query string) ([]Food, error) {
	foods, err := s.storage.SearchFood(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to search foods from storage due: %w", err)
	}
	return foods, nil
}
