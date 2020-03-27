package main

import (
	"strings"
	"time"
)

// https://golang.org/pkg/time/#pkg-constants
var predefinedTimeFormats = map[string]string{
	"default":     "2006-01-02 15:04:05.999999999 -0700 MST", // see https://golang.org/pkg/time/#Time.String
	"ansic":       time.ANSIC,
	"unixdate":    time.UnixDate,
	"rubydate":    time.RubyDate,
	"rfc822":      time.RFC822,
	"rfc822z":     time.RFC822Z,
	"rfc850":      time.RFC850,
	"rfc1123":     time.RFC1123,
	"rfc1123z":    time.RFC1123Z,
	"rfc3339":     time.RFC3339,
	"rfc3339nano": time.RFC3339Nano,
	"kitchen":     time.Kitchen,
	"stamp":       time.Stamp,
	"stampmilli":  time.StampMilli,
	"stampmicro":  time.StampMicro,
	"stampnano":   time.StampNano,
}

func predefinedTimeFormatKeys() []string {
	keys := []string{}
	for k := range predefinedTimeFormats {
		keys = append(keys, k)
	}
	return keys
}

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

func formatTime(t time.Time, format string) string {

	if len(format) == 0 {
		return t.String()
	}

	if v, ok := predefinedTimeFormats[strings.ToLower(format)]; ok {
		return t.Format(v)
	}
	return t.Format(format)
}
