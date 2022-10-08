package rest

import (
	"errors"
	"net/http"

	"github.com/ghazlabs/es-starterkit/internal/kuliner-api/core"
	"github.com/go-chi/render"
)

type respBody struct {
	StatusCode int         `json:"-"`
	OK         bool        `json:"ok"`
	Data       interface{} `json:"data,omitempty"`
	Err        string      `json:"err,omitempty"`
	Message    string      `json:"msg,omitempty"`
}

func (rb *respBody) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, rb.StatusCode)
	return nil
}

func newSuccessResp(data interface{}) *respBody {
	return &respBody{
		StatusCode: http.StatusOK,
		OK:         true,
		Data:       data,
	}
}

func newErrorResp(err error) *respBody {
	var restErr *apiError
	if !errors.As(err, &restErr) {
		restErr = newInternalServerError(err.Error())
	}
	return &respBody{
		StatusCode: restErr.StatusCode,
		OK:         false,
		Err:        restErr.Err,
		Message:    restErr.Message,
	}
}

type indexFoodReqBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (rb *indexFoodReqBody) Bind(r *http.Request) error {
	return nil
}

type updateFoodReqBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (rb *updateFoodReqBody) Bind(r *http.Request) error {
	return nil
}

type foodResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func newFoodResp(food core.Food) foodResp {
	return foodResp{
		ID:          food.ID,
		Name:        food.Name,
		Description: food.Description,
	}
}
