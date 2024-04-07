package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/blazee5/EffectiveMobile_test/internal/domain"
	"github.com/blazee5/EffectiveMobile_test/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CarRepository struct {
	db *pgxpool.Pool
}

func NewCarRepository(db *pgxpool.Pool) *CarRepository {
	return &CarRepository{db: db}
}

func (repo *CarRepository) Get(ctx context.Context, input domain.GetCarsRequest) (domain.CarList, error) {
	query := sq.Select("id, reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic").From("cars")

	if input.RegNum != "" {
		query = query.Where(squirrel.Eq{"reg_num": input.RegNum})
	}
	if input.Mark != "" {
		query = query.Where(squirrel.Eq{"mark": input.Mark})
	}
	if input.Model != "" {
		query = query.Where(squirrel.Eq{"model": input.Model})
	}
	if input.Year != 0 {
		query = query.Where(squirrel.Eq{"year": input.Year})
	}
	if input.Name != "" {
		query = query.Where(squirrel.Eq{"owner_name": input.Name})
	}
	if input.Surname != "" {
		query = query.Where(squirrel.Eq{"owner_surname": input.Surname})
	}
	if input.Patronymic != "" {
		query = query.Where(squirrel.Eq{"owner_patronymic": input.Patronymic})
	}
	if input.Limit == 0 {
		input.Limit = 5
	}

	query = query.Limit(uint64(input.Limit)).Offset(uint64(input.Offset)).PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return domain.CarList{}, err
	}

	rows, err := repo.db.Query(ctx, sql, args...)
	if err != nil {
		return domain.CarList{}, err
	}
	defer rows.Close()

	cars, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Car])

	if err != nil {
		return domain.CarList{}, err
	}

	countQuery := sq.Select("COUNT(*)").From("cars")

	countSql, countArgs, err := countQuery.ToSql()
	if err != nil {
		return domain.CarList{}, err
	}

	var total int
	err = repo.db.QueryRow(ctx, countSql, countArgs...).Scan(&total)
	if err != nil {
		return domain.CarList{}, err
	}

	return domain.CarList{
		Meta: domain.Meta{
			Total:  total,
			Limit:  input.Limit,
			Offset: input.Offset,
		},
		Cars: cars,
	}, nil
}

func (repo *CarRepository) Create(ctx context.Context, input []domain.Car) ([]models.Car, error) {
	cars := make([]models.Car, len(input))

	for i, car := range input {
		var createdCar models.Car

		err := repo.db.QueryRow(ctx, `INSERT INTO cars (reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic`, car.RegNum, car.Mark, car.Model, car.Year,
			car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic).Scan(&createdCar.ID, &createdCar.RegNum, &createdCar.Mark, &createdCar.Model, &createdCar.Year, &createdCar.Owner.Name, &createdCar.Owner.Surname, &createdCar.Owner.Patronymic)

		if err != nil {
			return nil, err
		}

		cars[i] = createdCar
	}

	return cars, nil
}

func (repo *CarRepository) Update(ctx context.Context, id int, input domain.UpdateCarRequest) (models.Car, error) {
	var car models.Car

	err := repo.db.QueryRow(ctx, `UPDATE cars SET
        reg_num = COALESCE(NULLIF($1, ''), reg_num),
        mark = COALESCE(NULLIF($2, ''), mark),
        model = COALESCE(NULLIF($3, ''), model),
        year = COALESCE(NULLIF($4, 0), year),
        owner_name = COALESCE(NULLIF($5, ''), owner_name),
        owner_surname = COALESCE(NULLIF($6, ''), owner_surname),
        owner_patronymic = COALESCE(NULLIF($7, ''), owner_patronymic)
        WHERE id = $8
        RETURNING id, reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic`,
		input.RegNum, input.Mark, input.Model, input.Year, input.Owner.Name, input.Owner.Surname, input.Owner.Patronymic, id).Scan(&car.ID, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic)
	if err != nil {
		return models.Car{}, err
	}

	return car, nil
}

func (repo *CarRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.Exec(ctx, `DELETE FROM cars WHERE id = $1`, id)

	return err
}
