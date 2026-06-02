package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"agussyahrilmubarok.github.io/web/internal/infrastructure/config"
	"agussyahrilmubarok.github.io/web/pkg/logger"
	"go.uber.org/zap"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type WEBServer struct {
	cfg    *config.Config
	router *gin.Engine
	server *http.Server
}

func NewWEBServer(cfg *config.Config) *WEBServer {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	store := cookie.NewStore([]byte(cfg.App.Session.Secret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   cfg.App.Session.MaxAge,
		HttpOnly: true,
		// Secure: true, // (HTTPS)
	})
	router.Use(sessions.Sessions(cfg.App.Name, store))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	return &WEBServer{
		cfg:    cfg,
		router: router,
		server: srv,
	}
}

func (s *WEBServer) Router() *gin.Engine {
	return s.router
}

func (s *WEBServer) Run() error {
	log := logger.Get()

	go func() {
		log.Info("server running...", zap.String("app_name", s.cfg.App.Name), zap.String("app_addr", s.server.Addr))
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down server....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Info("server exited gracefully")
	return nil
}
