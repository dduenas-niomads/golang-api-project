package requests

type UserRegistrationRequest struct {
	Name 	 string `json:"name" binding:"required"`
	Email 	 string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}