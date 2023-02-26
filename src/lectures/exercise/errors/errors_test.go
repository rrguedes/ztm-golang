package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:3:12", true},
		{"1:-4:12", false},
		{"0:59:59", true},
		{"", false},
		{"aa:bb:cc", false},
		{"5:23:", false},
	}
	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, o err deve ser nil", data.time, err)
		}
	}
}
