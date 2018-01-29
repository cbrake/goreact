package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cbrake/goreact/db"

	docopt "github.com/docopt/docopt-go"
)

func main() {
	usage := `Sample post.

Usage:
	sample-post <sn> <paramA> <paramB> [--dev]

Options:
	--dev  post to dvelopment server
`
	arguments, err := docopt.Parse(usage, nil, true, "Sample Post 1.0", false)

	if err != nil {
		fmt.Println("Error parsing arguments: ", err)
		return
	}

	sn := arguments["<sn>"].(string)
	paramA := arguments["<paramA>"].(string)
	paramB_ := arguments["<paramB>"].(string)
	paramB, _ := strconv.Atoi(paramB_)
	dev := arguments["--dev"].(bool)

	var server string

	if dev {
		server = "http://localhost:8090"
	} else {
		server = "https://your-production-server.com"
	}

	base, err := url.Parse(server)
	if err != nil {
		fmt.Println("Error parsing URL", err)
		return
	}

	u, err := url.Parse("sample/" + sn)

	url := base.ResolveReference(u)

	fmt.Println("posting to: ", url)

	data := db.Sample{
		ParamA: paramA,
		ParamB: paramB,
	}

	data_, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Error encoding data", err)
		return
	}

	fmt.Println("encoded data: ", string(data_))

	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(data_))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
