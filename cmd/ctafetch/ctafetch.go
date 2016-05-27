package main

import (
	"fmt"
	"github.com/sdimitro/gocta"
	"os"
	"strconv"
)

func main() {
	apiKey := os.Getenv("CTA_KEY")
	if apiKey == "" {
		fmt.Fprintf(os.Stderr,
			"ctafetch: CTA API key not set (CTA_KEY)\n")
		os.Exit(1)
	}

	stationName := os.Getenv("CTA_STATION")
	if stationName == "" {
		fmt.Fprintf(os.Stderr,
			"ctafetch: station name not set (CTA_STATION)\n")
		os.Exit(1)
	}

	iMapID, ok := gocta.StationMapID[stationName]
	if !ok {
		fmt.Fprintf(os.Stderr,
			"ctafetch: station name does not exist (yet?)\n")
		os.Exit(1)
	}

	mapID := strconv.Itoa(iMapID)
	resp, err := gocta.GetPredictions(apiKey, mapID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ctafetch: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", resp)
}
