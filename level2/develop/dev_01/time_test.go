package task1

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeNow(t *testing.T) {
	TimeNow("")
	fmt.Printf("\n")

	time.Sleep(time.Second * 1)

	TimeNow("0.beevik-ntp.pool.ntp.org")
	fmt.Printf("\n")
}
