package app

import (
	"context"
	"github.com/blazee5/EffectiveMobile_test/internal/handler"
	"github.com/blazee5/EffectiveMobile_test/internal/repository"
	"github.com/blazee5/EffectiveMobile_test/internal/service"
	"github.com/blazee5/EffectiveMobile_test/lib/db/postgres"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func Run(log *zap.SugaredLogger) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := postgres.New(ctx)
	repositories := repository.NewRepository(db)
	services := service.NewService(log, repositories)
	srv := handler.NewServer(log, services)

	go func() {
		if err := srv.Run(srv.InitRoutes()); err != nil {
			log.Fatalf("Error while start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		log.Infof("Error occured on server shutting down: %v", err)
	}

	db.Close()
}
