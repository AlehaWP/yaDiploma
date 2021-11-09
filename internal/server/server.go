package server

import (
	"context"
	"net/http"
	"time"

	"github.com/AlehaWP/yaDiploma.git/internal/config"
	"github.com/AlehaWP/yaDiploma.git/internal/handlers"
	"github.com/AlehaWP/yaDiploma.git/internal/middlewares"
	"github.com/AlehaWP/yaDiploma.git/internal/models"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	http.Server
	models.ServerDB
}

//Start server with router.
// func (s *Server) Start(ctx context.Context, repo models.Repository, opt models.Options) {
func (s *Server) Start(ctx context.Context) {
	r := chi.NewRouter()
	// handlers.NewHandlers(repo, opt)
	// middlewares.NewCookie(repo)
	r.Use(middlewares.ZipHandlerRead, middlewares.ZipHandlerWrite)
	r.Get("/*", handlers.HandlerStartPage)
	r.Post("/api/user/register", handlers.HandlerRegistration(s.NewDBUserRepo()))
	r.Post("/api/user/login", handlers.HandlerLogin(s.NewDBUserRepo()))
	r.Route("/api", func(r chi.Router) {
		r.Use(middlewares.CheckAuthorization(s.NewDBUserRepo()))
		r.Post("/user/orders", handlers.HandlersNewOrder(s.NewDBOrdersRepo()))
		r.Get("/user/orders", handlers.HandlersGetUserOrders(s.NewDBOrdersRepo()))
		r.Get("/user/balance", handlers.HandlerGetUserBalance(s.NewDBBalanceRepo()))
		r.Get("/user/balance/withdrawals", handlers.HandlerGetUserWithdrawals(s.NewDBBalanceRepo()))
		r.Post("/user/balance/withdraw", handlers.HandlerGetUserWithdraw(s.NewDBBalanceRepo()))
	})

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

	s.Addr = config.Cfg.ServAddr()
	logger.Info("Старт сервера по адресу", config.Cfg.ServAddr())
	s.Handler = r
	go s.ListenAndServe()

	logger.Info("Сервер запущен")

	<-ctx.Done()
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	s.Shutdown(ctx)
}
