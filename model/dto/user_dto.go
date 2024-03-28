package dto

type LoginRequestDto struct {
	Username string `json:"username" binding:"required"`
	Pass     string `json:"password" binding:"required"`
}

type LoginResponseDto struct {
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
}
