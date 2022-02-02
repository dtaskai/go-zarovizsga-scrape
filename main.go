package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	var old = ""
	var new = ""

	durationArg := flag.String("d", "10m", "Duration between scrapes")
	url := flag.String("w", "https://info.mik.pte.hu/hu/zarovizsga", "Website to scrape")
	flag.Parse()

	for {
		newResp, err := http.Get(*url)
		getTime()
		if err != nil {
			fmt.Println("Something went wrong!:(")
		}

		bodyOld, err := io.ReadAll(newResp.Body)
		if err != nil {
			fmt.Println("Something went wrong!:(")
		}
		new = string(bodyOld)

		duration, _ := time.ParseDuration(*durationArg)

		time.Sleep(duration)

		if old != new && old != "" {
			fmt.Println("Fent az eredmény!:)")
			os.Exit(0)
		}

		old = new
	}

}

func getTime() {
	hours, minutes, sec := time.Now().Clock()
	fmt.Printf("%d:%02d:%02d - Scraped \r\n", hours, minutes, sec)
}
