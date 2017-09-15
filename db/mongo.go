package db

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

// Logger for db module
var L = log.WithField("ctx", "db")

var mgoSession *mgo.Session

func Connect(devMode bool) (err error) {
	var url string
	if devMode {
		fmt.Println("Not supported yet")
	} else {
		url = "mongodb://user:pass:57621/dbname"
	}

	L.WithField("url", url).Info("Connecting Mongo")

	mgoSession, err = mgo.Dial(url)
	if err != nil {
		L.WithError(err).Error("Error connecting")
		return
	}

	return
}

type Ses struct {
	session *mgo.Session
	Samples *mgo.Collection
}

// this session must be closed once you are done with database operations
func GetSession() (s Ses) {
	s.session = mgoSession.Copy()
	db := s.session.DB("")
	s.Samples = db.C("Samples")
	return
}

func (s *Ses) Close() {
	s.session.Close()
}
