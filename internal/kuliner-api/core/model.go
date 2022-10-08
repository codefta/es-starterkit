package core

import "strings"

const defaultSearchLimit = 10

type Food struct {
	ID          string
	Name        string
	Description string
}

type FoodInput struct {
	Name        string
	Description string
}

func (i FoodInput) Validate() error {
	if len(i.Name) == 0 {
		return ErrMissingName
	}
	if len(i.Description) == 0 {
		return ErrMissingDescription
	}
	return nil
}

func (i FoodInput) ToFood(id ...string) Food {
	var idResult string
	if len(id) != 0 {
		idResult = id[0]
	} else {
		idResult = strings.ReplaceAll(strings.ToLower(strings.TrimSpace(i.Name)), " ", "-")
	}

	return Food{
		ID:          idResult,
		Name:        i.Name,
		Description: i.Description,
	}
}
