package justamin

import (
	"fmt"
	"math"
	"time"
)

func Duration(moment time.Time) (humanized string) {
	now := time.Now()
	duration := moment.Sub(now)
	fmt.Printf("%+v - %+v : %+v\n", moment, now, duration)

	return Humanize(duration)
}

func Humanize(duration time.Duration) (humanized string) {
	if duration < 1000000000 {
		return "now"
	}

	seconds := int64(math.Round(duration.Seconds()))

	if seconds == 1 {
		return fmt.Sprint("a second ago")
	} else if seconds < 60 {
		return fmt.Sprintf("%d seconds ago", seconds)
	}

	minutes := int64(math.Round(duration.Minutes()))

	if minutes == 1 {
		return fmt.Sprint("a minute ago")
	} else if minutes < 60 {
		return fmt.Sprintf("%d minutes ago", minutes)
	}

	hours := int64(math.Round(duration.Hours()))

	if hours == 1 {
		return fmt.Sprint("an hour ago")
	} else if hours < 24 {
		return fmt.Sprintf("%d hours ago", hours)
	}

	days := int64(math.Round(duration.Hours() / 24))

	if days == 1 {
		return fmt.Sprint("a day ago")
	} else if days < 7 {
		return fmt.Sprintf("%d days ago", days)
	}

	weeks := int64(math.Round(duration.Hours() / 168))

	if weeks == 1 {
		return fmt.Sprint("a week ago")
	} else if weeks < 7 {
		return fmt.Sprintf("%d weeks ago", weeks)
	}

	return humanized
}
