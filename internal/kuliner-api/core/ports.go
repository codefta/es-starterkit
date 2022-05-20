package core

import "context"

type Storage interface {
	IndexFood(ctx context.Context, food Food) error
	DeleteFood(ctx context.Context, id string) error
	SearchFood(ctx context.Context, query string) ([]Food, error)
}
