package model

type UserResponse struct {
	ID    int64  `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}
