package service

import (
	"context"
	"fmt"
	"github.com/blazee5/EffectiveMobile_test/internal/domain"
	"github.com/blazee5/EffectiveMobile_test/internal/models"
	"github.com/blazee5/EffectiveMobile_test/internal/repository"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type CarService struct {
	log  *zap.SugaredLogger
	repo *repository.Repository
}

func NewCarService(log *zap.SugaredLogger, repo *repository.Repository) *CarService {
	return &CarService{log: log, repo: repo}
}

func (s *CarService) GetCars(ctx context.Context, input domain.GetCarsRequest) (domain.CarList, error) {
	return s.repo.Get(ctx, input)
}

func (s *CarService) CreateCars(ctx context.Context, input domain.CreateCarsRequest) ([]models.Car, error) {
	apiHost := os.Getenv("API_HOST")

	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}

	var mutex sync.Mutex
	cars := make([]domain.Car, 0)
	results := make(chan domain.Car, len(input.RegNums))

	wg := sync.WaitGroup{}
	wg.Add(len(input.RegNums))

	for _, regNum := range input.RegNums {
		go func(regNum string) {
			defer wg.Done()

			res, err := client.Get(fmt.Sprintf("%s/info?regNum=%s", apiHost, regNum))

			if err != nil {
				s.log.Infof("error while get car info: %v", err)
				return
			}

			defer res.Body.Close()

			var car []domain.Car //TODO: убрать []

			body, err := io.ReadAll(res.Body)

			if err != nil {
				s.log.Infof("error while read car body: %v", err)
				return
			}

			if err := json.Unmarshal(body, &car); err != nil {
				s.log.Infof("error while decode car info: %v", err)
				return
			}

			results <- car[0] //TODO: убрать [0]
		}(regNum)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for car := range results {
		mutex.Lock()
		cars = append(cars, car)
		mutex.Unlock()
	}

	return s.repo.Create(ctx, cars)
}

func (s *CarService) UpdateCar(ctx context.Context, id int, input domain.UpdateCarRequest) (models.Car, error) {
	return s.repo.Update(ctx, id, input)
}

func (s *CarService) DeleteCar(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
