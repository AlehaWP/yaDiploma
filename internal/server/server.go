package server

import (
	"context"
	"net/http"
	"time"

	"github.com/AlehaWP/yaDiploma.git/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	http.Server
}

//Start server with router.
// func (s *Server) Start(ctx context.Context, repo models.Repository, opt models.Options) {
func (s *Server) Start() {
	r := chi.NewRouter()
	// handlers.NewHandlers(repo, opt)
	// middlewares.NewCookie(repo)
	r.Use(middlewares.CheckAuthorization, middlewares.ZipHandlerRead, middlewares.ZipHandlerWrite)
	// //r.Use(middlewares.ZipHandlerRead, middlewares.ZipHandlerWrite)

	// r.Get("/user/urls", handlers.HandlerUserPostURLs)
	// r.Get("/ping", handlers.HandlerCheckDBConnect)
	// r.Route("/{id}", func(r chi.Router) {
	// 	r.Use(middlewares.URLCtx)
	// 	r.Get("/", handlers.HandlerURLGet)
	// })
	// r.Post("/", handlers.HandlerURLPost)
	// r.Post("/api/shorten", handlers.HandlerAPIURLPost)
	// r.Post("/api/shorten/batch", handlers.HandlerAPIURLsPost)
	// r.Delete("/api/user/urls", handlers.HandlerDeleteUserUrls)

	// s.Addr = opt.ServAddr()
	s.Handler = r
	go s.ListenAndServe()

	// <-ctx.Done()
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	s.Shutdown(ctx)
}
