package response



type CreateDiscountResponse struct {
	ID string `json:"d_id"`
    C_ID string `json:"c_id"`
    O_ID string `json:"o_id"`
    Percent float32 `json:"d_percent"`
    Name string `json:"d_name"`
    Code string `json:"d_code"`
    Description string `json:"d_description"`
    StartDate string `json:"d_startDate"`
    ExpDate string `json:"d_expDate"`
    Status bool `json:"d_status"`
}

type GetDiscountResponse CreateDiscountResponse

type GetDiscountsResponse struct {
    Discounts []GetDiscountResponse `json:"discount"`
}

type UpdateDiscount CreateDiscountResponse