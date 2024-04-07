package models

type Owner struct {
	Name       string `json:"name" db:"owner_name"`
	Surname    string `json:"surname" db:"owner_surname"`
	Patronymic string `json:"patronymic" db:"owner_patronymic"`
}
