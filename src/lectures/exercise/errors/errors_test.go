package timeparse

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:45:49", true},
		{"23:59:59", true},
		{"00:00:00", true},
		{"1:2:3", true},
		{"1:-2:3", false},
		{"00:00", false},
		{"00:12:", false},
		{"", false},
		{"fail", false},
		{"aa:bb:cc", false},
		{"9999:12:12", false},
		{"12:9999:12", false},
		{"12:12:9999", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)

		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		} else if !data.ok && err == nil {
			t.Errorf("%v: error should have been raised", data.time)
		}
	}
}
