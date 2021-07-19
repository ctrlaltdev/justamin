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
		hour:    "in an hour",
		hours:   "in %d hours",
		day:     "in a day",
		days:    "in %d days",
		week:    "in a week",
		weeks:   "in %d weeks",
	},
}

var DEBUG = false

func EnableDebug() {
	DEBUG = true
}

func debug(msg string) {
	if DEBUG {
		fmt.Println(msg)
	}
}

func Duration(moment time.Time) (humanized string) {
	now := time.Now()
	duration := moment.Sub(now)

	debug(fmt.Sprintf("time: %+v - now: %+v : duration: %+v", moment, now, duration))

	humanized = Humanize(duration)
	debug(fmt.Sprintf("humanized: %+v", humanized))

	return humanized
}

func Humanize(duration time.Duration) (humanized string) {
	var format timeFormat
	if duration < 0 {
		format = formats.past
		debug("PAST")
		duration = -duration
	} else {
		format = formats.future
		debug("FUTURE")
	}
	debug(fmt.Sprintf("duration: %+v", duration))

	if duration < 1000000000 {
		return format.now
	}

	seconds := int64(math.Round(duration.Seconds()))
	debug(fmt.Sprintf("seconds: %+v", seconds))

	if seconds == 1 {
		humanized = fmt.Sprint(format.second)
		debug(humanized)
		return humanized
	} else if seconds < 60 {
		humanized = fmt.Sprintf(format.seconds, seconds)
		debug(humanized)
		return humanized
	}

	minutes := int64(math.Round(duration.Minutes()))
	debug(fmt.Sprintf("minutes: %+v", minutes))

	if minutes == 1 {
		humanized = fmt.Sprint(format.minute)
		debug(humanized)
		return humanized
	} else if minutes < 60 {
		humanized = fmt.Sprintf(format.minutes, minutes)
		debug(humanized)
		return humanized
	}

	hours := int64(math.Round(duration.Hours()))
	debug(fmt.Sprintf("hours: %+v", hours))

	if hours == 1 {
		humanized = fmt.Sprint(format.hour)
		debug(humanized)
		return humanized
	} else if hours < 24 {
		humanized = fmt.Sprintf(format.hours, hours)
		debug(humanized)
		return humanized
	}

	days := int64(math.Round(duration.Hours() / 24))
	debug(fmt.Sprintf("days: %+v", days))

	if days == 1 {
		humanized = fmt.Sprint(format.day)
		debug(humanized)
		return humanized
	} else if days < 7 {
		humanized = fmt.Sprintf(format.days, days)
		debug(humanized)
		return humanized
	}

	weeks := int64(math.Round(duration.Hours() / 168))
	debug(fmt.Sprintf("weeks: %+v", weeks))

	if weeks == 1 {
		humanized = fmt.Sprint(format.week)
		debug(humanized)
		return humanized
	} else {
		humanized = fmt.Sprintf(format.weeks, weeks)
		debug(humanized)
		return humanized
	}
}
