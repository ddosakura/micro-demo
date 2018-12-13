package main

import (
	"errors"
	"os"

	"gopkg.in/mgo.v2"
)

// DB_HOST: "datastore:27017"/"mongo:27017"

func CreateSession() (*mgo.Session, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, errors.New("need to set DB_HOST")
	}
	s, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	s.SetMode(mgo.Monotonic, true)
	return s, nil
}
