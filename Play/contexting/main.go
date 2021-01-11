package main

import (
	"fmt"
	"time"
)

func timeToQuarters(time time.Time) int64 {

	return time.Unix() / 60 / 15
}

func quartersToTime(quarters int64) time.Time {

	return time.Unix(quarters*15*60, 0)
}

func main() {

	//now := time.Now()
	//
	//year, month, day := now.Date()
	//hour, min, sec := now.Clock()

	//fmt.Printf("now.Date(): %v %v %v\n", year, month, day)
	//fmt.Printf("now.Clock(): %v %v %v\n", hour, min, sec)

	//days := now.YearDay()
	//
	//fmt.Printf("days: %v\n", days) // day of the year
	//fmt.Printf("hours: %v + %v\n", days*24, hour) // hours of the year
	//fmt.Printf("15 min: %v + %v + %v\n", days*384, hour*4, min/15) // 15 min of the year
	//
	//quarterHoursOfTheYear := days*384 + hour*4 + min/15
	//
	//fmt.Printf("quarterHoursOfTheYear: %v\n", quarterHoursOfTheYear) // 15 min of the year
	//fmt.Printf("quarterHoursOfTheYear * 15: %v\n", quarterHoursOfTheYear * 15) // min of the year

	//unix := now.Unix()
	//fmt.Printf("unix seconds: %v\n", unix)
	//fmt.Printf("unix minutes: %v (rest: %v sec)\n", unix/60, unix%60)
	//fmt.Printf("unix hours: %v (rest: %v minutes)\n", unix/60/60, (unix/60)%60)
	//fmt.Printf("unix days: %v (rest: %v hours)\n", unix/60/60/24, (unix/60/60)%(24))
	//fmt.Printf("unix 15 minutes: %v (rest: %v minutes)\n", unix/60/15, (unix/60)%15)

	//fmt.Printf("unix ~ month: %v\n", unix/60/60/24/30)
	//fmt.Printf("unix ~ year: %v\n", unix/60/60/24/365)
	//fmt.Printf("unix ~ year: %v\n", unix/60/60/24/30/12)
	//fmt.Printf("unix mod(15 min): %v\n", (61 % 10 ))

	now := time.Now()
	quarters := timeToQuarters(now)
	quartersTime := quartersToTime(quarters)

	fmt.Printf("time.Now(): %v\n", now)
	fmt.Printf("timeToQuarters(now): %v\n", quarters)
	fmt.Printf("quartersToTime(now): %v\n", quartersTime)

	bookingFrom := time.Date(2019, 12, 30, 19, 59, 59, 99, time.Local)
	quarters = timeToQuarters(bookingFrom)
	quartersTime = quartersToTime(quarters)

	fmt.Printf("bookingFrom: %v\n", bookingFrom)
	fmt.Printf("timeToQuarters(bookingFrom): %v\n", quarters)
	fmt.Printf("quartersToTime(bookingFrom): %v\n", quartersTime)

}
