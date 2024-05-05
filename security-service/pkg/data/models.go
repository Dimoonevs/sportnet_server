package data

import "time"

type User struct {
	Id          int32
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	DateOfBirth string `json:"date_of_birth"`
	Status      Status `json:"status"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	TimeZone    string `json:"time_zone"`
	Active      bool   `json:"active"`
	Created_at  time.Time
	Updated_at  time.Time
}

type Status int32

const (
	Status_SPORTSMEN           Status = 0
	Status_COACH               Status = 1
	Status_SPORTSMEN_AND_COACH Status = 3
)
