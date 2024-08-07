package user

type registerRequest struct {
	Username string `json:"username" binding:"required,gte=3,lte=36"`
	Password string `json:"password" binding:"required,gte=6,lte=128"`
}

type userResponse struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Money    float64 `json:"money"`
}
