package booking

import "fmt"
import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    const layout string = "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)

    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    const layout string = "January 2, 2006 15:04:05"
    t, _ := time.Parse(layout, date)

	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    const layout string = "Monday, January 2, 2006 15:04:05"
    t, _ := time.Parse(layout, date)

	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	const layout string = "1/2/2006 15:04:05"
    t, _ := time.Parse(layout, date)

    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday().String(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
