package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cbrake/goreact/sms"
	log "github.com/sirupsen/logrus"
)

func main() {

	to := flag.String("to", "", "SMS recipient")
	msg := flag.String("msg", "", "Message to send")
	debug := flag.Bool("debug", false, "enable Debug logging")

	flag.Parse()

	if *to == "" || *msg == "" {
		fmt.Println("send-sms usage:")
		flag.PrintDefaults()
		os.Exit(-1)
	}

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	err := sms.Send(*to, *msg)

	if err != nil {
		fmt.Println("Error sending SMS: ", err)
	} else {
		fmt.Println("SMS message sent")
	}
}
