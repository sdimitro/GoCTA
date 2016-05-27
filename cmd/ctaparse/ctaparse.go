package main

import (
	"fmt"
	"github.com/sdimitro/gocta"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ctaparse: %v\n", err)
		os.Exit(1)
	}

	res, err := gocta.ParseCTAResponse(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ctaparse: %v\n", err)
		os.Exit(1)
	}

	now := time.Now()
	for _, eta := range res.PredictionList {
		var t string
		if eta.IsApp {
			t = "Due"
		} else {
			arrival, err := gocta.ParseCTATime(eta.Arrival)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ctaparse: %v\n", err)
				os.Exit(1)
			}

			prediction := arrival.Sub(now).Minutes()
			t = fmt.Sprintf("%.2f mins", prediction)
		}
		fmt.Println(eta.DestName, t)
	}
}
