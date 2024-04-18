package dto
type CreateUserDto struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}



type CreateSuccessResponse struct {
	Id int64 `json:"id"`
}