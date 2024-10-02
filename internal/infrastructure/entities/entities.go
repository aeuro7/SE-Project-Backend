package entities

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
    ID       pgtype.UUID `gorm:"primaryKey;type:uuid" json:"c_id"`
    Email    string `gorm:"column:c_email" json:"c_email"`
    Password string `gorm:"column:c_password" json:"c_password"`
    Name     string `gorm:"column:c_name" json:"c_name"`
    Phone    string `gorm:"column:c_phone" json:"c_phone"`
    
    MusicLines []MusicLine `gorm:"foreignKey:C_ID" json:"music_lines"`
    IGLines    []IGLine    `gorm:"foreignKey:C_ID" json:"ig_lines"`
    Tables     []Table     `gorm:"foreignKey:C_ID" json:"tables"`
}

type MusicLine struct {
    ID   pgtype.UUID `gorm:"primaryKey;type:uuid" json:"m_id"`
    C_ID pgtype.UUID `gorm:"column:c_id" json:"c_id"`
    Name string `gorm:"column:m_name" json:"m_name"`
}

type IGLine struct {
    ID   pgtype.UUID `gorm:"primaryKey;type:uuid" json:"i_id"`
    C_ID pgtype.UUID `gorm:"column:c_id" json:"c_id"`
    Name string `gorm:"column:ig_account" json:"ig"`
}

type Table struct {
    ID     pgtype.UUID `gorm:"primaryKey;type:uuid" json:"t_id"`
    C_ID   pgtype.UUID `gorm:"column:c_id" json:"c_id"`
    Status bool   `gorm:"column:t_status" json:"t_status"`
}

type Order struct {
    ID         pgtype.UUID `gorm:"primaryKey;type:uuid" json:"o_id"`
    Status     bool      `gorm:"column:o_status" json:"o_status"`
    Time       time.Time `gorm:"column:o_time" json:"o_time"`
    Url        pgtype.UUID    `gorm:"column:o_urlslip" json:"o_urlslip"`
    TotalPrice float32   `gorm:"column:o_totalprice" json:"o_totalprice"`
    
    OrderLines []OrderLine `gorm:"foreignKey:O_ID" json:"order_lines"`
}

type OrderLine struct {
    ID       pgtype.UUID `gorm:"primaryKey;type:uuid" json:"i_id"`
    O_ID     pgtype.UUID `gorm:"column:o_id" json:"o_id"`
    P_ID     pgtype.UUID `gorm:"column:p_id" json:"p_id"`
    Quantity int    `gorm:"column:i_quantity" json:"i_quantity"`
}

type Menu struct {
    ID          pgtype.UUID   `gorm:"primaryKey;type:uuid" json:"m_id"`
    Price       float32  `gorm:"column:m_price" json:"m_price"`
    Description string   `gorm:"column:m_description" json:"m_description"`
    Url         []string `gorm:"type:text[]" json:"m_url_list"`
}
