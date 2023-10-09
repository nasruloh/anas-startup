package user

import "time"

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	Password_hash  string
	AvatarFileName string
	Role           string
	CreatedAt      	time.Time
	UpdatedAt		time.Time
}