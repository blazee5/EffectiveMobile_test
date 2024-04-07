package domain

type Owner struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type OwnerQuery struct {
	Name       string `form:"name"`
	Surname    string `form:"surname"`
	Patronymic string `form:"patronymic"`
}
