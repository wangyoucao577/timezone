package main

import "flag"

var flags struct {
	utc            bool
	tzDataBaseName string
}

func init() {
	flag.BoolVar(&flags.utc, "u", false, "Display UTC time.")
	flag.StringVar(&flags.tzDataBaseName, "tz", "", "IANA Time Zone database Name, such as \"America/New_York\".")
}
