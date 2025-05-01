package main

import (
	"fmt"
	"time"
)

/*
Given a time in 12-hour AM/PM format, convert it to military (24-hour) time.

Note:
- 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.
*/

func main() {
	layout := "03:04:05 PM"      //layout 12
	militaryLayout := "15:04:05" //layout 24

	customTime := "07:05:45 PM"

	normalTime, err := time.Parse(layout, customTime)
	if err != nil {
		fmt.Println(err)
	}
	militaryHour := normalTime.Format(militaryLayout)
	fmt.Println(militaryHour)
}
