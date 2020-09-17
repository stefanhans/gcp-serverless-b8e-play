package types

import "time"

type Member struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type ItemType int

const (
	CAR ItemType = iota
	BIKE
	VAN
)

func (itemType ItemType) String() string {
	names := [...]string{
		"Car",
		"Bike",
		"Van"}

	// `itemType` is out of range of ItemType
	if itemType < CAR || itemType > VAN {
		return "Unknown"
	}
	// return the name of an ItemType
	return names[itemType]
}

type Item struct {
	ID   int32    `json:"id"`
	Name string   `json:"name"`
	Type ItemType `json:"type"`
}

type BookingStatus int

const (
	REQUESTED BookingStatus = iota
	CONFIRMED
	REJECTED
)

func (bookingStatus BookingStatus) String() string {
	names := [...]string{
		"Requested",
		"Confirmed",
		"Rejected"}

	// `bookingStatus` is out of range of BookingStatus
	if bookingStatus < REQUESTED || bookingStatus > REJECTED {
		return "Unknown"
	}
	// return the name of a BookingStatus
	return names[bookingStatus]
}

type Booking struct {
	ID         int32         `json:"id"`
	User       Member        `json:"user"`
	Share      Item          `json:"share"`
	From       time.Time     `json:"from"`
	To         time.Time     `json:"to"`
	Status     BookingStatus `json:"status"`
	StatusTime time.Time     `json:"status-time"`
}
