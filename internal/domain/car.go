package domain

import "github.com/blazee5/EffectiveMobile_test/internal/models"

type CreateCarsRequest struct {
	RegNums []string `json:"regNums"`
}

type GetCarsRequest struct {
	RegNum string `form:"regNum"`
	Mark   string `form:"mark"`
	Model  string `form:"model"`
	Year   int    `form:"year"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
	OwnerQuery
}

type UpdateCarRequest struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  Owner  `json:"owner"`
}

type Car struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  Owner  `json:"owner"`
}

type Meta struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type CarList struct {
	Meta `json:"meta"`
	Cars []models.Car `json:"cars"`
}
