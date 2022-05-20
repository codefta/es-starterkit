package core

import "strings"

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

func (i FoodInput) ToFood() Food {
	return Food{
		ID:          strings.ReplaceAll(strings.ToLower(strings.TrimSpace(i.Name)), " ", "-"),
		Name:        i.Name,
		Description: i.Description,
	}
}
