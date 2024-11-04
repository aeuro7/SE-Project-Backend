package response

import "github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"

type LoginResponse struct{
	User entities.User `json:"user"`
	Token string `json:"token"`
}

type RegisterResponse struct{
	ID 		 string `json:"ID"`
	Email    string `json:"email"`
    Name     string `json:"name"`
}