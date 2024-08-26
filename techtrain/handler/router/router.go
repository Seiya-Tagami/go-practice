package router

import (
	"database/sql"
	"net/http"

	"techtrain-go-practice/handler"
	"techtrain-go-practice/middleware"
	"techtrain-go-practice/middleware/utils"
	"techtrain-go-practice/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	svc := service.NewTODOService(todoDB)

	// register routes
	mux := http.NewServeMux()
	mux.Handle("/healthz", utils.ChainMiddlewares(
		handler.NewHealthzHandler(),
		middleware.Recovery,
		middleware.SetOS,
		middleware.Logger,
	))
	mux.Handle("/do-panic", utils.ChainMiddlewares(
		handler.NewPanicHandler(),
		middleware.Recovery,
		middleware.SetOS,
		middleware.Logger,
	))
	mux.Handle("/todos", utils.ChainMiddlewares(
		handler.NewTODOHandler(svc),
		middleware.Recovery,
		middleware.SetOS,
		middleware.Logger,
	))

	return mux
}
