package entities

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Email    string `gorm:"not null"`
}
