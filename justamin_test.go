package justamin

import (
	"testing"
	"time"
)

type HumanizeCase struct {
	duration time.Duration
	expected string
}

func TestHumanize(t *testing.T) {
	var cases = []HumanizeCase{
		{time.Duration(1000), "now"},
		{time.Duration(1 * 1000000000), "a second ago"},
		{time.Duration(2 * 1000000000), "2 seconds ago"},
		{time.Duration(1 * 60 * 1000000000), "a minute ago"},
		{time.Duration(4 * 60 * 1000000000), "4 minutes ago"},
		{time.Duration(1 * 60 * 60 * 1000000000), "an hour ago"},
		{time.Duration(3.1 * 60 * 60 * 1000000000), "3 hours ago"},
		{time.Duration(1 * 24 * 60 * 60 * 1000000000), "a day ago"},
		{time.Duration(6.2 * 24 * 60 * 60 * 1000000000), "6 days ago"},
		{time.Duration(1 * 7 * 24 * 60 * 60 * 1000000000), "a week ago"},
		{time.Duration(2.13 * 7 * 24 * 60 * 60 * 1000000000), "2 weeks ago"},
	}

	for _, c := range cases {

		res := Humanize(c.duration)

		if res != c.expected {
			t.Logf("should return %s instead of %s", c.expected, res)
			t.Fail()
		}

	}
}
