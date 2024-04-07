package handler

import (
	"context"
	_ "github.com/blazee5/EffectiveMobile_test/docs"
	"github.com/blazee5/EffectiveMobile_test/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

type Server struct {
	log        *zap.SugaredLogger
	httpServer *http.Server
	services   *service.Service
}

func NewServer(log *zap.SugaredLogger, services *service.Service) *Server {
	return &Server{log: log, services: services}
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) InitRoutes() *gin.Engine {
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
	}))

	api := r.Group("/api/v1")

	car := api.Group("/cars")
	{
		car.POST("", s.CreateCars)
		car.GET("", s.GetCars)
		car.PUT("/:id", s.UpdateCar)
		car.DELETE("/:id", s.DeleteCar)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
