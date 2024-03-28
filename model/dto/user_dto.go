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

type LoginRequestDto struct {
	Username string `json:"username" binding:"required"`
	Pass     string `json:"password" binding:"required"`
}

type LoginResponseDto struct {
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
}
