package application

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
	cfg *config.Config
	db  *gorm.DB
	log *zap.Logger

	userRepository repository.IUserRepository

	jwtService  service.IJWTService
	authService service.IAuthService
	userService service.IUserService
}

func (app *App) Run() error {
	handler := app.newGinRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.cfg.App.Port),
		Handler: handler,
	}

	go func() {
		app.log.Info("starting http server", zap.String("port", app.cfg.App.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.log.Fatal("failed to start http server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.log.Info("shutdown signal received...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		app.log.Error("failed to gracefully shutdown http server", zap.Error(err))
		return err
	}

	app.log.Info("http server stopped successfully")
	return nil
}

func New(
	cfg *config.Config,
	db *gorm.DB,
) *App {
	userRepository := repository.NewUserRepository(db)

	jwtService := service.NewJWTService(&cfg.JWT)
	authService := service.NewAuthService(userRepository, jwtService)
	userService := service.NewUserService(userRepository)

	return &App{
		cfg: cfg,
		db:  db,
		log: logger.Get(),

		userRepository: userRepository,

		jwtService:  jwtService,
		authService: authService,
		userService: userService,
	}
}
