package storage

import (
	"encoding/json"

	"github.com/ghazlabs/es-starterkit/internal/kuliner-api/core"
)

type foodDoc struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func newFoodDoc(food core.Food) foodDoc {
	return foodDoc{
		Name:        food.Name,
		Description: food.Description,
	}
}

func (d foodDoc) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

type searchFoodsResult struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			ID     string  `json:"_id"`
			Source foodDoc `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (r searchFoodsResult) GetFoods() []core.Food {
	var foods []core.Food
	for _, hit := range r.Hits.Hits {
		foods = append(foods, core.Food{
			ID:          hit.ID,
			Name:        hit.Source.Name,
			Description: hit.Source.Description,
		})
	}
	return foods
}
