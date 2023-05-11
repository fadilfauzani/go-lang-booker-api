package dtos

import "time"

type BookingDto struct {
	Id              int       `json:"id"`
	GorId           int       `json:"gorid"`
	UserId          int       `json:"userid"`
	BookedTimeStart time.Time `json:"bookedTimeStart"`
	BookedTimeEnd   time.Time `json:"bookedTimeEnd"`
	BookedAt        time.Time `json:"bookedAt"`
	IsPaid          bool      `json:"isPaid"`
	IsVisited       bool      `json:"isVisited"`
}
