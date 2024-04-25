package models

import "time"

type User struct {
	UserID      string    `json:"id_user"`
	RoleID      int       `json:"id_role"`
	Username    string    `json:"username"`
	FirstName   string    `json:"first_name" validate:"required,min=1,max=100"`
	LastName    string    `json:"last_name" validate:"required,min=1,max=100"`
	Password    string    `json:"password" validate:"required,min=8"`
	Email       string    `json:"email" validate:"email,required"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	Balance     int       `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
