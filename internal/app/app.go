package app

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test_app/config"
	mongo_repo "test_app/internal/adapter/db/mongo-repo"
	"test_app/internal/handler/http/auth"
	"test_app/internal/usecase"
	"test_app/middleware"
	"test_app/package/database/mongodb"
	jwt_service "test_app/package/jwt"
	"test_app/package/logger/logger/sl"
	validate "test_app/package/validator"
	"time"
)

func Run(cfg *config.Config) {
	log := sl.SetupLogger(cfg.Env)
	log.Info("starting app", slog.String("env", cfg.Env))

	mongo, err := mongodb.New(cfg)
	if err != nil {
		log.Error("failed to init database", sl.Err(err))
		os.Exit(1)
	}
	db := mongo.Client.Database(cfg.Mongo.DBName)
	defer func() {
		if err := mongo.Client.Disconnect(context.TODO()); err != nil {
			log.Error("failed to close", err)
			os.Exit(1)
		}
	}()

	tokenService, err := jwt_service.NewTokenService(cfg.Token.PrivateKey)
	if err != nil {
		log.Error("failed to init token service", err)
		os.Exit(1)
	}
	refreshTokenUseCase := usecase.NewAuthUseCase(mongo_repo.NewRefreshTokenRepo(db), tokenService, cfg.Token.AccessTokenTTL)
	authMiddleware := middleware.NewAuthMiddleware(refreshTokenUseCase, cfg.Token.PrivateKey)

	app := echo.New()
	app.Validator = &validate.Validator{Valid: validator.New()}
	//app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//}))
	auth.NewAuthRouter(app, log, refreshTokenUseCase, authMiddleware)

	log.Info("starting server", slog.String("host", cfg.HTTPServer.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      app,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error("failed to start server")
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))
	}

	log.Info("server stopped")
}
