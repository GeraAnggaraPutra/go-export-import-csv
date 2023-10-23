package model

import "time"

type User struct {
	ID    int64  `json:"id"`
	GUID  string `json:"guid"`
	Email string `json:"email"`
	// Password  string `json:"omitempty"`
	RoleGUID  string     `json:"role_guid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ExportUser struct {
	NO        int
	Email     string
	RoleName  string
	CreatedAt time.Time
}
