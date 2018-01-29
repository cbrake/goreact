package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/cbrake/goreact/db"
	"github.com/cbrake/goreact/frontend"
	docopt "github.com/docopt/docopt-go"
	assetfs "github.com/elazarl/go-bindata-assetfs"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var port = ":8090"

func vcConfigGetAllHander(rw http.ResponseWriter, req *http.Request) {
	mSes := db.GetSession()
	defer mSes.Close()

	devices := []db.Sample{}

	iter := mSes.Samples.Find(nil).Iter()
	device := db.Sample{}

	for iter.Next(&device) {
		devices = append(devices, device)
	}

	data, err := json.Marshal(devices)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("vcConfigGetAllHander: ", string(data))

	rw.Header().Set("Connect-Type", "text/json")
	var reader = bytes.NewBuffer(data)
	io.Copy(rw, reader)
}

func vcConfigGetHander(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("vcConfigGetHander")
}

func vcConfigPostHander(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("vcConfigPostHander")
	params := mux.Vars(req)
	sn, ok := params["sn"]
	if !ok {
		fmt.Println("missing serial number")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("sn: ", sn)
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error reading body", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("body: ", string(data))

	var sample db.Sample

	if json.Unmarshal(data, &sample) != nil {
		fmt.Println("error parsing json response", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	sample.SerialNumber = sn

	fmt.Printf("sample: %+v\n", sample)
	mSes := db.GetSession()
	defer mSes.Close()

	_, err = mSes.Samples.Upsert(bson.M{"serialNumber": sn}, bson.M{"$set": sample})
	if err != nil {
		fmt.Println("Error saving sample to mongodb", err)
	}
}

func main() {
	fmt.Println("Go React demo app, v1")

	usage := `Go/React demo application.

Usage:
	app [--dev]

Options:
	--dev  start in development mode
`

	arguments, err := docopt.Parse(usage, nil, true, "Go React", false)

	if err != nil {
		fmt.Println("Error parsing arguments: ", err)
		return
	}

	dev := arguments["--dev"].(bool)

	if dev {
		fmt.Println("Starting in DEVELOPMENT mode")
	}

	/*
		err := db.Connect(false)

		if err != nil {
			fmt.Println("Error connecting to mongo")
			return
		}
	*/

	r := mux.NewRouter()

	if dev {
		r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))
	} else {
		afs := assetfs.AssetFS{
			Asset:     frontend.Asset,
			AssetDir:  frontend.AssetDir,
			AssetInfo: frontend.AssetInfo,
			Prefix:    "",
		}

		r.PathPrefix("/").Handler(http.FileServer(&afs))
	}
	r.HandleFunc("/sample", vcConfigGetAllHander).Methods("GET")
	r.HandleFunc("/sample/{sn}", vcConfigGetHander).Methods("GET")
	r.HandleFunc("/sample/{sn}", vcConfigPostHander).Methods("POST")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Println("starting server on: http://localhost" + port)
	err = http.ListenAndServe(port, loggedRouter)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

	fmt.Println("Done ...")
}
