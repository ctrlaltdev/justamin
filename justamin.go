package justamin

import (
	"fmt"
	"math"
	"time"
)

type timeFormat struct {
	now     string
	second  string
	seconds string
	minute  string
	minutes string
	hour    string
	hours   string
	day     string
	days    string
	week    string
	weeks   string
}
type timeFormats struct {
	past   timeFormat
	future timeFormat
}

var formats = timeFormats{
	past: timeFormat{
		now:     "now",
		second:  "a second ago",
		seconds: "%d seconds ago",
		minute:  "a minute ago",
		minutes: "%d minutes ago",
		hour:    "an hour ago",
		hours:   "%d hours ago",
		day:     "a day ago",
		days:    "%d days ago",
		week:    "a week ago",
		weeks:   "%d weeks ago",
	},
	future: timeFormat{
		now:     "now",
		second:  "in a second",
		seconds: "in %d seconds",
		minute:  "in a minute",
		minutes: "in %d minutes",
		hour:    "ain n hour",
		hours:   "in %d hours",
		day:     "in a day",
		days:    "in %d days",
		week:    "in a week",
		weeks:   "in %d weeks",
	},
}

func Duration(moment time.Time) (humanized string) {
	now := time.Now()
	duration := moment.Sub(now)
	fmt.Printf("%+v - %+v : %+v\n", moment, now, duration)

	return Humanize(duration)
}

func Humanize(duration time.Duration) (humanized string) {
	var format timeFormat
	if duration < 0 {
		format = formats.past
		duration = -duration
	} else {
		format = formats.future
	}

	if duration < 1000000000 {
		return format.now
	}

	seconds := int64(math.Round(duration.Seconds()))

	if seconds == 1 {
		return fmt.Sprint(format.second)
	} else if seconds < 60 {
		return fmt.Sprintf(format.seconds, seconds)
	}

	minutes := int64(math.Round(duration.Minutes()))

	if minutes == 1 {
		return fmt.Sprint(format.minute)
	} else if minutes < 60 {
		return fmt.Sprintf(format.minutes, minutes)
	}

	hours := int64(math.Round(duration.Hours()))

	if hours == 1 {
		return fmt.Sprint(format.hour)
	} else if hours < 24 {
		return fmt.Sprintf(format.hours, hours)
	}

	days := int64(math.Round(duration.Hours() / 24))

	if days == 1 {
		return fmt.Sprint(format.day)
	} else if days < 7 {
		return fmt.Sprintf(format.days, days)
	}

	weeks := int64(math.Round(duration.Hours() / 168))

	if weeks == 1 {
		return fmt.Sprint(format.week)
	} else if weeks < 7 {
		return fmt.Sprintf(format.weeks, weeks)
	}

	return humanized
}
