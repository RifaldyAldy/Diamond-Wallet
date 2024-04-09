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

type UserRequestEditDto struct {
	Name        string `json:"name"`
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

type UpdatePinResponse struct {
	UserId string `json:"user_id"`
	Pin    string `json:"pin"`
}

type UpdatePinRequest struct {
	UserId string `json:"user_id"`
	OldPin string `json:"old_pin" binding:"required"`
	NewPin string `json:"new_pin" binding:"required"`
}
type UpdatePinRequestSwag struct {
	OldPin string `json:"old_pin" binding:"required"`
	NewPin string `json:"new_pin" binding:"required"`
}

type RekeningDtoSwag struct {
	Rekening string `json:"rekening"`
}
