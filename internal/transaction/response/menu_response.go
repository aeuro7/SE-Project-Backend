package response

import "github.com/jackc/pgx/v5/pgtype"

type CreateMenuResponse struct {
	ID          pgtype.UUID `json:"m_id"`
	Price       float64     `json:"m_price"`
	Description string      `json:"m_description"`
	Url         string    `json:"m_url"`
}

type GetMenuResponse CreateMenuResponse

type GetAllMenuResponse struct {
		Menu []GetMenuResponse `json:"Menu"`
}

type UpdateMenuResponse CreateMenuResponse