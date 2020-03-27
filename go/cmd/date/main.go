package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	logger := log.New(os.Stderr, "", 0)

	t, err := getCurrentTime(flags.utc, flags.tzDataBaseName)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}

	fmt.Println(formatTime(t, flags.format))
}
