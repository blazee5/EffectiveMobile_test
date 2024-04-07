package service

import (
	"context"
	"github.com/blazee5/EffectiveMobile_test/internal/domain"
	"github.com/blazee5/EffectiveMobile_test/internal/models"
	"github.com/blazee5/EffectiveMobile_test/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	Car
}

type Car interface {
	GetCars(ctx context.Context, input domain.GetCarsRequest) (domain.CarList, error)
	CreateCars(ctx context.Context, input domain.CreateCarsRequest) ([]models.Car, error)
	UpdateCar(ctx context.Context, id int, input domain.UpdateCarRequest) (models.Car, error)
	DeleteCar(ctx context.Context, id int) error
}

func NewService(log *zap.SugaredLogger, repo *repository.Repository) *Service {
	return &Service{Car: NewCarService(log, repo)}
}
