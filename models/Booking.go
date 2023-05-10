package models

import "time"

type Booking struct {
	Id              int       `json:"id"`
	GorId           int       `json:"gorid"`
	UserId          int       `json:"userid"`
	BookedTimeStart time.Time `json:"bookedTimeStart"`
	BookedTimeEnd   time.Time `json:"bookedTimeEnd"`
	BookedAt time.Time `json:"bookedat"`
	IsPaid   bool      `json:"isPaid"`
}
