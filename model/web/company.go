package web

type CompanyCreateRequest struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type CompanyResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
}


