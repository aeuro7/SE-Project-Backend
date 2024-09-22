package requests

type CreateUserRequest struct{
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}