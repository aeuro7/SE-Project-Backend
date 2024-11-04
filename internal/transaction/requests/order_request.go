package requests


type CreateOrderRequest struct{
	T_ID string `json:"t_id"`
	Url  string  `json:"o_url"`
}

type CreateOrderWithOrderLinesRequest struct{
	T_ID string `json:"t_id"`
	Url  string  `json:"o_url"`

	OrderLines []CreateOrderLineRequest `json:"orderlines"`
}