package task1

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// TimeNow prints current time or returns error code 1
func TimeNow(server string) {
	if server == "" {
		server = "0.beevik-ntp.pool.ntp.org"
	}

	response, err := ntp.Query(server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error message: %s\n", err)
		os.Exit(1)
	}

	currentTime := time.Now()
	exactTime := time.Now().Add(response.ClockOffset)

	fmt.Println(currentTime)
	fmt.Println(exactTime)
}
