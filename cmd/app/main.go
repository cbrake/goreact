package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goelm/db"
	"goelm/frontend"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

var port = ":8090"

func rootHandler(rw http.ResponseWriter, req *http.Request) {
	if bs, err := frontend.Asset("frontend/index.html"); err != nil {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		var reader = bytes.NewBuffer(bs)
		io.Copy(rw, reader)
	}
}

func elmJsHandler(rw http.ResponseWriter, req *http.Request) {
	if bs, err := frontend.Asset("frontend/elm.js"); err != nil {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		var reader = bytes.NewBuffer(bs)
		io.Copy(rw, reader)
	}
}

func bootstrapHandler(rw http.ResponseWriter, req *http.Request) {
	if bs, err := frontend.Asset("frontend/bootstrap.min.css"); err != nil {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		rw.Header().Set("Content-Type", "text/css")
		var reader = bytes.NewBuffer(bs)
		io.Copy(rw, reader)
	}
}

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
	fmt.Println("OKI Portal, v4")

	err := db.Connect(false)

	if err != nil {
		fmt.Println("Error connecting to mongo")
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/elm.js", elmJsHandler)
	r.HandleFunc("/bootstrap.min.css", bootstrapHandler)
	r.HandleFunc("/sample", vcConfigGetAllHander).Methods("GET")
	r.HandleFunc("/sample/{sn}", vcConfigGetHander).Methods("GET")
	r.HandleFunc("/sample/{sn}", vcConfigPostHander).Methods("POST")

	log.Println("starting server on: http://localhost" + port)
	err = http.ListenAndServe(port, r)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

	fmt.Println("Done ...")
}
