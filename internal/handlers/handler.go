package handlers

import (
	"github.com/alioygur/gores"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	router *mux.Router
}

type ResponseData struct {
	Code         int         `json:"code,omitempty"`
	Text         string      `json:"text,omitempty"`
	ClientErrors interface{} `json:"errors,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Meta         interface{} `json:"meta,omitempty"`
	Links        interface{} `json:"links,omitempty"`
	Error        error       `json:"-"`
}

func NewAppHandler(muxRouter *mux.Router) Handler {
	apiV1Router := muxRouter.PathPrefix("/api/v1").Subrouter()

	handler := Handler{
		router: apiV1Router,
	}

	handler.router.Use(handler.corsJsonMiddleware)

	return handler

}

func (h *Handler) getMuxId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	stringID := vars["id"]

	intID, err := strconv.Atoi(stringID)
	return intID, err
}

func (h *Handler) corsJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("X-Version", "dev")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS,DELETE")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if r.Method != http.MethodOptions {
			next.ServeHTTP(w, r)
		}
	})
}

func (h *Handler) httpResponseJSON(w http.ResponseWriter, code int, response ResponseData) {
	if code == http.StatusInternalServerError {
		// log error
	}

	_ = gores.JSON(w, code, response)
}
