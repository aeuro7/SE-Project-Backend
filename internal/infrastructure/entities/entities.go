package entities

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
    ID       pgtype.UUID `gorm:"primaryKey;type:uuid"`
    Email    string `gorm:"column:c_email"`
    Password string `gorm:"column:c_password"`
    Name     string `gorm:"column:c_name"`
    Phone    string `gorm:"column:c_phone"`
    
    MusicLines []MusicLine `gorm:"foreignKey:C_ID"`
    IGLines    []IGLine    `gorm:"foreignKey:C_ID"`
    Tables     []Table     `gorm:"foreignKey:C_ID"`
}

type MusicLine struct {
    ID   pgtype.UUID `gorm:"primaryKey;type:uuid"`
    C_ID pgtype.UUID `gorm:"column:c_id"`
    Name string `gorm:"column:m_name"`
}

type IGLine struct {
    ID   pgtype.UUID `gorm:"primaryKey;type:uuid"`
    C_ID pgtype.UUID `gorm:"column:c_id"`
    Name string `gorm:"column:ig_account"`
}

type Table struct {
    ID     string `gorm:"primaryKey;type:string"`
    C_ID   pgtype.UUID `gorm:"column:c_id"`
    Status string   `gorm:"type:char(1);default:'A';check:t_status IN ('A', 'O', 'R');column:t_status"`
}

type Order struct {
    ID         pgtype.UUID `gorm:"primaryKey;type:uuid"`
    Status     bool      `gorm:"column:o_status"`
    Time       time.Time `gorm:"column:o_time"`
    Url        pgtype.UUID    `gorm:"column:o_urlslip"`
    TotalPrice float32   `gorm:"column:o_totalprice"`
    
    OrderLines []OrderLine `gorm:"foreignKey:O_ID"`
}

type OrderLine struct {
    ID       pgtype.UUID `gorm:"primaryKey;type:uuid"`
    O_ID     pgtype.UUID `gorm:"column:o_id"`
    P_ID     pgtype.UUID `gorm:"column:p_id"`
    Quantity int    `gorm:"column:i_quantity"`
}

type Menu struct {
    ID          pgtype.UUID   `gorm:"primaryKey;type:uuid"`
    Price       float32  `gorm:"column:m_price"`
    Description string   `gorm:"column:m_description"`
    Url         []string `gorm:"type:text[]"`
}
