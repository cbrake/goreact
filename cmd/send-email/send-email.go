package main

import (
	"fmt"

	"github.com/cbrake/goreact/email"

	docopt "github.com/docopt/docopt-go"
)

func main() {
	usage := `Email test utility.

Usage:
	email-test --to=<to> --subject=<subject> --body=<body>

Options:
	--to=TO            email address to send to
	--subject=SUBEJCT  email subject
	--body=BODY        email BODY
`
	arguments, err := docopt.Parse(usage, nil, true, "Email test", false)

	if err != nil {
		fmt.Println("Error parsing arguments: ", err)
		return
	}

	fmt.Printf("arguments: %+v\n", arguments)

	to := arguments["--to"].(string)
	subject := arguments["--subject"].(string)
	body := arguments["--body"].(string)
	to_ := []string{to}

	err = email.Send(to_, subject, body)

	if err != nil {
		fmt.Println("Error sending mail: ", err)
	}

}
