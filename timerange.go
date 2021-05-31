package daterange

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

func (dr DateRange) Equal(dr2 DateRange) bool {
	return dr.Start.Equal(dr2.Start) && dr.End.Equal(dr2.End)
}

func (dr DateRange) String() string {
	return fmt.Sprintf("%s -> %s", dr.Start.Format(time.RFC3339), dr.End.Format(time.RFC3339))
}

func (dr DateRange) Duration() time.Duration {
	return dr.End.Sub(dr.Start)
}

func (dr DateRange) Intersect(dr2 DateRange) bool {
	if dr2.Start.After(dr.Start) && dr2.Start.Before(dr.End) {
		return true
	}
	if dr2.End.After(dr.Start) && dr2.End.Before(dr.End) {
		return true
	}
	if dr2.Start.Before(dr.Start) && dr2.End.After(dr.End) {
		return true
	}
	return false
}

func (dr *DateRange) Lengthen(dur time.Duration) {
	dr.End = dr.End.Add(dur)
}

func (dr *DateRange) Shorten(dur time.Duration) {
	if dr.Duration() <= dur {
		dr.End = dr.Start
	} else {
		diff := dr.Duration() - dur
		dr.End = dr.Start.Add(diff)
	}
}

func New(startDate, endDate string) (*DateRange, error) {
	start, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		return nil, err
	}

	if end.Before(start) {
		return nil, errors.New("end date cannot be before start date")
	}

	return &DateRange{start, end}, nil
}
