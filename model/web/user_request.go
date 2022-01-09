package web


type UserCreateRequest struct {
	Name string `validate:"required" json:"name"`
	Company CompanyCreateRequest `json:"company" validate:"required"`

}
type UserUpdateRequest struct {
	ID int `validate:"required" json:"id"`
	Name string `validate:"required" json:"name"`
}

type UserResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Company CompanyResponse `json:"company"`
}