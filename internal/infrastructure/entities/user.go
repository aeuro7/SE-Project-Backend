package entities

type User struct {
    ID    int    `gorm:"primaryKey" json:"id"`
    Fname string `gorm:"column:fname" json:"fname"`
    Lname string `gorm:"column:lname" json:"lname"`
}
