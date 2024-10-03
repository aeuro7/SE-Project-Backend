package requests


type CreateTableRequest struct{
	ID string `json:"t_id"`
}

type UpdateTableRequest struct{
	C_ID string `json:"c_id"`
	Status string `json:"t_status"`
}