package transport

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateNewItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	text := r.FormValue("text")
	return CreateItemRequest{Text: text}, nil
}

func NewHTTPHandler(e Endpoints) http.Handler {
	router := mux.NewRouter()
	v1SubRouter := router.PathPrefix("/api/v1").Subrouter()
	endpoint := "/items"
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	v1SubRouter.Methods("POST").Path(endpoint).Handler(kithttp.NewServer(
		e.CreateItemEndpoint,
		decodeCreateNewItemRequest,
		encodeResponse,
		options...))

	return router
}
