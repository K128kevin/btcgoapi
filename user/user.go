package user

import "time"

type User struct {
	Id			int64
	Email			string
	LastLogin		time.Location
	Registered		time.Location
}

type Users []User
