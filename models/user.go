package models

import "time"

type User struct {
	Id         int
	Name       string
	Created_at time.Time
	Status     bool
}

func (user *User) AddUser(id int, name string, created_at time.Time, status bool) {
	user.Id = id
	user.Name = name
	user.Created_at = created_at
	user.Status = status
}
