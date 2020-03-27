package main

import (
	"time"
)

func getCurrentTime(utc bool, tzDataBaseName string) (time.Time, error) {
	t := time.Now()

	if utc {
		return t.UTC(), nil
	}

	if len(tzDataBaseName) > 0 {
		l, err := time.LoadLocation(tzDataBaseName)
		if err != nil {
			return t, err
		}
		return t.In(l), nil
	}

	return t, nil
}
