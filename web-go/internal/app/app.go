package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"agussyahrilmubarok.github.io/web/internal/config"
	"agussyahrilmubarok.github.io/web/internal/domain"
	"agussyahrilmubarok.github.io/web/internal/repository"
	"agussyahrilmubarok.github.io/web/internal/service"
	"agussyahrilmubarok.github.io/web/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	cfg *config.Config
	db  *gorm.DB
	log *zap.Logger

	userRepository repository.IUserRepository

	authService service.IAuthService
	userService service.IUserService
}

func (app *App) Run() error {
	app.autoMigrate()

	handler := app.setGinRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.cfg.App.Port),
		Handler: handler,
	}

	go func() {
		app.log.Info("Server started", zap.String("port", app.cfg.App.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.log.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		app.log.Error("Server forced to shutdown", zap.Error(err))
		return err
	}

	app.log.Info("Server exited")
	return nil
}

func (app *App) autoMigrate() {
	app.db.AutoMigrate(&domain.User{})
}

func NewApp(cfg *config.Config, db *gorm.DB) *App {
	userRepository := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)

	return &App{
		cfg: cfg,
		db:  db,
		log: logger.Get(),

		userRepository: userRepository,

		authService: authService,
		userService: userService,
	}
}
