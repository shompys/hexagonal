package http

import (
	"net/http"

	"github.com/shompys/hexagonal/internal/user/infrastructure/adapters/http/handler"
)

type Router struct {
	mux     *http.ServeMux
	handler *handler.HandlerUser
}

func NewRouter(h *handler.HandlerUser) *Router {
	return &Router{
		handler: h,
		mux:     http.NewServeMux(),
	}
}

func (r *Router) RegisterRoutes() {
	r.mux.HandleFunc("GET /users", r.handler.GetUsers)
	r.mux.HandleFunc("GET /users/{id}", r.handler.GetUser)
	r.mux.HandleFunc("POST /users", r.handler.CreateUser)
	r.mux.HandleFunc("PATCH /users/{id}", r.handler.UpdateUser)
	r.mux.HandleFunc("DELETE /user/{id}", r.handler.DeleteUser)
}
func (r *Router) Handler() *http.ServeMux {
	return r.mux
}
