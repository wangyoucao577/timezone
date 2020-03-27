package main

import (
	"flag"
	"fmt"
	"strings"
)

var flags struct {
	utc            bool
	tzDataBaseName string
	format         string
}

func init() {
	flag.BoolVar(&flags.utc, "u", false, "Display UTC time.")
	flag.StringVar(&flags.tzDataBaseName, "tz", "", "IANA Time Zone database Name, such as \"America/New_York\".")

	flag.StringVar(&flags.format, "f", "unixdate", fmt.Sprintf("Format textual representation of the time value. Support format like \"Mon Jan 2 15:04:05 MST 2006\", see more in https://golang.org/pkg/time/#Time.Format. \nSome predefined formats: %s. ", strings.Join(predefinedTimeFormatKeys(), ",")))
}
