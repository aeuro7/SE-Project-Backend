package response

type UserCreateResponse struct{
	ID int `json:"ID"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

type GetUserResponse struct{
	ID int `json:"ID"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

type GetUsersResponse struct{
	Users []GetUserResponse
}