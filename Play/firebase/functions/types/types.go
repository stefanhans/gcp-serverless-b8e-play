package types

import (
	"fmt"
	"google.golang.org/genproto/googleapis/type/latlng"
	"time"
)

type Booking struct {
	DocId         string    `firestore:"DocId"`
	User          string    `firestore:"User"`
	Vehicle       string    `firestore:"Vehicle"`
	VehicleType   string    `firestore:"VehicleType"`
	VehicleStatus string    `firestore:"VehicleStatus"`
	ParkingLot    string    `firestore:"ParkingLot"`
	From          time.Time `firestore:"From"`
	To            time.Time `firestore:"To"`
	Status        string    `firestore:"Status"`
	StatusTime    time.Time `firestore:"StatusTime"`
}

func (booking *Booking) FromToQuarters() int64 {

	return booking.From.Unix() / 60 / 15
}

func (booking *Booking) ToToQuarters() int64 {

	return booking.To.Unix() / 60 / 15
}

//func timeToQuarters(time time.Time) int64 {
//
//	return time.Unix()/60/15
//}
//
//func quartersToTime(quarters int64) time.Time {
//
//	return time.Unix(quarters*15*60, 0)
//}

type Vehicle struct {
	DocId       string         `firestore:"DocId"`
	Name        string         `firestore:"Name"`
	Type        string         `firestore:"Type"`
	Status      string         `firestore:"Status"`
	ParkingLot  string         `firestore:"ParkingLot"`
	GeoPoint    *latlng.LatLng `firestore:"GeoPoint"`
	Description string         `firestore:"Description"`
}

type User struct {
	DocId       string `firestore:"DocId"`
	Name        string `firestore:"Name"`
	Type        string `firestore:"Type"`
	Status      string `firestore:"Status"`
	Description string `firestore:"Description"`
}

//type Message interface {
//	CreateResponse (status string,
//		error string) string
//}
//
//
//type MessageData struct {
//	Status string
//	Error  string
//	Data   string
//	Message string
//}

func CreateResponse(status string,
	error string, data []byte) string {

	return fmt.Sprintf("{ "+
		"\n    %q: %q, "+
		"\n    %q: %q, "+
		"\n    %q: %s "+
		"\n}",
		"Status", status,
		"Error", error,
		"Data", string(data))
}

type MasterData struct {
	Users      []User    `firestore:"users"`
	Vehicles   []Vehicle `firestore:"vehicles"`
	Bookings   []Booking `firestore:"bookings"`
	From       time.Time `firestore:"from"`
	To         time.Time `firestore:"to"`
	Status     string    `firestore:"status"`
	StatusTime time.Time `firestore:"status-time"`
}
