package db

// common data structures

type Sample struct {
	SerialNumber string `json:"SerialNumber" bson:"SerialNumber"`
	ParamA       string `json:"paramA" bson:"paramA"`
	ParamB       int    `json:"paramB" bson:"paramB"`
}
