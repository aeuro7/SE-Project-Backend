package requests

import "time"

type CreateDiscountRequest struct {
    C_ID string `json:"c_id"`
    O_ID string `json:"o_id"`
    Percent float32 `json:"d_percent"`
    Name string `json:"d_name"`
    Code string `json:"d_code"`
    Description string `json:"d_description"`
    StartDate   time.Time  `json:"d_startDate"`
	ExpDate     time.Time  `json:"d_expDate"`
    Status bool `json:"d_status"`
}