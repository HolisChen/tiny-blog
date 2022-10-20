package request

type UserLoginRequest struct {
	LoginId  string `json:"loginId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterRequest struct {
	Name       string `json:"name" binding:"required"`
	LoginId    string `json:"loginId" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required"`
	Mail       string `json:"mail" binding:"required"`
}
