package rest

import (
	"errors"
	"net/http"

	"github.com/ghazlabs/es-starterkit/internal/kuliner-api/core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"gopkg.in/validator.v2"
)

type API struct {
	svc core.Service
}

type Config struct {
	Service core.Service `validate:"nonnil"`
}

func NewAPI(cfg Config) (*API, error) {
	err := validator.Validate(cfg)
	if err != nil {
		return nil, err
	}
	api := &API{svc: cfg.Service}
	return api, nil
}

func (a *API) GetHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/foods", func(r chi.Router) {
		r.Get("/", a.serveSearchFoods)
		r.Post("/", a.serveIndexFood)
		r.Delete("/{food_id}", a.serveDeleteFood)
	})

	return r
}

func (a *API) serveSearchFoods(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	foods, err := a.svc.SearchFoods(r.Context(), query)
	if err != nil {
		render.Render(w, r, newErrorResp(err))
		return
	}
	var frs []foodResp
	for _, food := range foods {
		frs = append(frs, newFoodResp(food))
	}
	render.Render(w, r, newSuccessResp(map[string]interface{}{
		"query": query,
		"foods": frs,
	}))
}

func (a *API) serveIndexFood(w http.ResponseWriter, r *http.Request) {
	var rb indexFoodReqBody
	err := render.Bind(r, &rb)
	if err != nil {
		render.Render(w, r, newErrorResp(newBadRequestError(err.Error())))
		return
	}
	food, err := a.svc.IndexFood(r.Context(), core.FoodInput{
		Name:        rb.Name,
		Description: rb.Description,
	})
	if err != nil {
		switch err {
		case core.ErrMissingName:
			err = newBadRequestError("missing `name`")
		case core.ErrMissingDescription:
			err = newBadRequestError("missing `description`")
		}
		render.Render(w, r, newErrorResp(err))
		return
	}
	fr := newFoodResp(*food)
	render.Render(w, r, newSuccessResp(fr))
}

func (a *API) serveDeleteFood(w http.ResponseWriter, r *http.Request) {
	foodID := chi.URLParam(r, "food_id")
	err := a.svc.DeleteFood(r.Context(), foodID)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			err = newNotFoundError()
		}
		render.Render(w, r, newErrorResp(err))
		return
	}
	render.Render(w, r, newSuccessResp(nil))
}
