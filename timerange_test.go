package daterange

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		Start string
		End   string
	}{
		{"2021-01-01T05:00:00Z", "2021-01-01T05:00:00Z"},
		{"2021-01-01T05:00:00Z", "2021-02-01T05:00:00Z"},
		{"2021-12-31T05:00:00Z", "2022-01-01T05:00:00Z"},
		{"2024-02-29T05:00:00Z", "2024-02-29T06:00:00Z"},
	}
	for _, test := range tests {
		if _, err := New(test.Start, test.End); err != nil {
			t.Errorf("Error parsing %v: %v", test, err)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		first  string
		second string
		third  string
		fourth string
		want   bool
	}{
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-03T05:00:00Z", "2021-01-04T05:00:00Z", false},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-02T06:00:00Z", "2021-01-04T05:00:00Z", false},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-02T05:00:01Z", "2021-01-04T05:00:00Z", false},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-02T04:00:00Z", "2021-01-04T05:00:00Z", true},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-04T05:00:00Z", false},
	}
	for i, test := range tests {
		range1, _ := New(test.first, test.second)
		range2, _ := New(test.third, test.fourth)
		if got := range1.Intersect(*range2); got != test.want {
			t.Errorf("test #%v: wanted %v but got %v", i, test.want, got)
		}
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		first  string
		second string
		third  string
		fourth string
		want   bool
	}{
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", true},
		{"2024-01-01T05:00:00Z", "2024-02-29T05:00:00Z", "2024-01-01T05:00:00Z", "2024-02-29T05:00:00Z", true},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-01T05:00:00Z", "2022-01-02T05:00:00Z", false},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-01T05:00:00Z", "2021-01-02T05:00:01Z", false},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-02T05:00:00Z", "2021-01-04T05:00:00Z", false},
	}
	for i, test := range tests {
		range1, _ := New(test.first, test.second)
		range2, _ := New(test.third, test.fourth)
		if got := range1.Equal(*range2); test.want != got {
			t.Errorf("test #%v: wanted %v but got %v", i, test.want, got)
		}
	}
}

func TestLengthen(t *testing.T) {
	tests := []struct {
		first  string
		second string
		dur    time.Duration
		third  string
		fourth string
	}{
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", time.Hour, "2021-01-01T05:00:00Z", "2021-01-02T06:00:00Z"},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", time.Hour * 3, "2021-01-01T05:00:00Z", "2021-01-02T08:00:00Z"},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", time.Hour * 24, "2021-01-01T05:00:00Z", "2021-01-03T05:00:00Z"},
	}
	for i, test := range tests {
		range1, _ := New(test.first, test.second)
		want, _ := New(test.third, test.fourth)
		if range1.Lengthen(test.dur); !range1.Equal(*want) {
			t.Errorf("test #%v: didn't lengthen range properly \n want: %v \n got:  %v \n duration: %v", i, want, range1, test.dur)
		}
	}
}

func TestShorten(t *testing.T) {
	tests := []struct {
		first  string
		second string
		dur    time.Duration
		third  string
		fourth string
	}{
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", time.Hour, "2021-01-01T05:00:00Z", "2021-01-02T04:00:00Z"},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", time.Hour * 3, "2021-01-01T05:00:00Z", "2021-01-02T02:00:00Z"},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", time.Hour * 24, "2021-01-01T05:00:00Z", "2021-01-01T05:00:00Z"},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", time.Second * 30, "2021-01-01T05:00:00Z", "2021-01-02T04:59:30Z"},
		{"2021-01-01T05:00:00Z", "2021-01-02T05:00:00Z", time.Hour * 28, "2021-01-01T05:00:00Z", "2021-01-01T05:00:00Z"},
	}
	for i, test := range tests {
		range1, _ := New(test.first, test.second)
		want, _ := New(test.third, test.fourth)
		if range1.Shorten(test.dur); !range1.Equal(*want) {
			t.Errorf("test #%v: didn't shorten range properly \n want: %v \n got:  %v \n duration: %v", i, want, range1, test.dur)
		}
	}
}
