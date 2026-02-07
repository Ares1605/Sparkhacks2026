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
	return time.Time(d.Time).Format(`"2006-01-02"`)
}

func (d *datefmt) From(data string) error {
	var err error

	d.Time, err = time.Parse("2006-01-02", data)
	if err != nil {
		return errors.New("Invalid date format")
	}

	return nil
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
