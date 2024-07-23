package http

import (
	"time"
)

func timeConvert(strTime string) (time.Time, error) {
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, strTime)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}
