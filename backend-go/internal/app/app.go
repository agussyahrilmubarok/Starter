package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"agussyahrilmubarok.github.io/backend/internal/config"
	"agussyahrilmubarok.github.io/backend/internal/repository"
	"agussyahrilmubarok.github.io/backend/internal/service"
	"agussyahrilmubarok.github.io/backend/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	config *config.Config
	db     *gorm.DB

	userRepository repository.IUserRepository

	authService service.IAuthService
	userService service.IUserService
}

func (app *App) Run() error {
	handler := app.setGinRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.App.Port),
		Handler: handler,
	}

	go func() {
		logger.Info("Server started", zap.String("port", app.config.App.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", zap.Error(err))
		return err
	}

	logger.Info("Server exited")
	return nil
}

func NewApp(
	config *config.Config,
	db *gorm.DB,
) *App {
	userRepository := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService()

	return &App{
		config: config,
		db:     db,

		userRepository: userRepository,

		authService: authService,
		userService: userService,
	}
}
