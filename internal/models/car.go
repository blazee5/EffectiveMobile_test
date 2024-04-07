package models

type Car struct {
	ID     int    `json:"id" db:"id"`
	RegNum string `json:"regNum" db:"regNum"`
	Mark   string `json:"mark" db:"mark"`
	Model  string `json:"model" db:"model"`
	Year   int    `json:"year" db:"year"`
	Owner  `json:"owner"`
}
