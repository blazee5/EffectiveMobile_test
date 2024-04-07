package repository

import (
	"context"
	"github.com/blazee5/EffectiveMobile_test/internal/domain"
	"github.com/blazee5/EffectiveMobile_test/internal/models"
	"github.com/blazee5/EffectiveMobile_test/internal/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Car
}

type Car interface {
	Get(ctx context.Context, input domain.GetCarsRequest) (domain.CarList, error)
	Create(ctx context.Context, input []domain.Car) ([]models.Car, error)
	Update(ctx context.Context, id int, input domain.UpdateCarRequest) (models.Car, error)
	Delete(ctx context.Context, id int) error
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Car: postgres.NewCarRepository(db),
	}
}
