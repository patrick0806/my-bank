package utils

import (
	"fmt"
	"time"
)

var layout = "2006-01-02"

func ParseStringToTIme(dateString string) (*time.Time, error) {

	t, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Printf("Fail to convert Date %v", err)
		return nil, err
	}
	return &t, nil
}

func FormatDate(date time.Time) string {
	return date.Format(layout)
}
