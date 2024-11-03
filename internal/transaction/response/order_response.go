package response

type CreateOrderResponse struct{
	ID string `json:"o_id"`
	T_ID string `json:"t_id"`
    Time string `json:"o_time"`
	Url  string `json:"o_url"`
}

type GetOrderResponse CreateOrderResponse

type GetOrdersResponse struct{
	Orders []GetOrderResponse `json:"orders"`
}


