package ntp

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

func GetCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(time)
}
