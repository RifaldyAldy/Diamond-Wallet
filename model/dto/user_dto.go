package dto

type UserRequestDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
