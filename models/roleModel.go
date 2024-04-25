package models

import "time"

type Role struct {
	RoleID     int       `json:"id_role"`
	RoleName   string    `json:"role_name"`
	RoleDesc   string    `json:"role_desc"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
