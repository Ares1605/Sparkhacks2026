package db

import (
	"errors"
	"fmt"
	"time"
)

type datefmt struct {
	time.Time
}

func (d datefmt) String() string {
	return time.Time(d.Time).Format(time.DateOnly)
}

func (d *datefmt) From(data string) error {
	formats := []string{
		time.DateOnly,
		time.DateTime,
	}

	for _, f := range formats {
		parsed, err := time.Parse(f, data)
		if err != nil {
			continue
		}
		d.Time = parsed
		return nil
	}

	return errors.New("Invalid date format")
}

func (d *datefmt) UnmarshalJSON(data []byte) error {
	return d.From(string(data))
}

func (d datefmt) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (y *datefmt) Scan(value any) error {
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("Bad value provided to scan")
	}

	return y.From(s)
}
