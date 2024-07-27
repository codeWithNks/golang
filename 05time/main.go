package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Learning time in Go\n")

	currentTime := time.Now()
	fmt.Println("Current time is ", currentTime)
	formattedTime := currentTime.Format("02-01-2006 15:04:05 Monday") // Remember 02 is for day, 01 is for month, 2006 is for year, 15 is for hour, 04 is for minute, 05 is for second
	fmt.Println("Formatted time is ", formattedTime)
	// All other useful methods
	fmt.Println("Year is ", currentTime.Year())
	fmt.Println("Month is ", currentTime.Month())
	fmt.Println("Day is ", currentTime.Day())
	fmt.Println("Hour is ", currentTime.Hour())
	fmt.Println("Minute is ", currentTime.Minute())
	fmt.Println("Second is ", currentTime.Second())
	fmt.Println("Weekday is ", currentTime.Weekday())
	fmt.Println("Location is ", currentTime.Location())
	fmt.Println("Timezone is ", currentTime.Location().String())
	// Gettinf data from some random date (27-07-2024, 18:58)
	randomDate := time.Date(2024, time.July, 27, 18, 58, 0, 0, time.UTC)
	fmt.Println("Random date is ", randomDate)
	fmt.Println("Random date is ", randomDate.Format("02-01-2006 15:04:05 Monday"))

}
