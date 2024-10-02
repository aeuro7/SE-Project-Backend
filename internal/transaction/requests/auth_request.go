package requests

type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}


type RegisterRequest struct{
    Email    string `json:"email"`
    Password string `son:"password"`
    Name     string `json:"name"`
    Phone    string `json:"phone"`
}