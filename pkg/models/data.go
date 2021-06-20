package models

import "time"

type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Id    int    `json:"id,omitempty"`
}

type Interview struct {
	ID           int       `json:"id,omitempty"`
	StartTime    time.Time `json:"starttime,omitempty"`
	EndTime      time.Time `json:"endtime,omitempty"`
	Participants []User    `json:"participants,omitempty"`
}

type Response struct {
	Success bool        `json:"success,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
