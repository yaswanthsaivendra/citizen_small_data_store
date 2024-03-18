package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

type CitizenModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName   string             `bson:"firstName" json:"firstName"`
	LastName    string             `bson:"lastName" json:"lastName"`
	DateOfBirth time.Time          `bson:"dateOfBirth" json:"dateOfBirth"`
	Gender      string             `bson:"gender" json:"gender"`
	Address     string             `bson:"address" json:"address"`
	City        string             `bson:"city" json:"city"`
	State       string             `bson:"state" json:"state"`
	Pincode     string             `bson:"pincode" json:"pincode"`
}

type CitizenSerializer struct {
	FirstName   string    `bson:"firstName" json:"firstName"`
	LastName    string    `bson:"lastName" json:"lastName"`
	DateOfBirth time.Time `bson:"dateOfBirth" json:"dateOfBirth"`
	Gender      string    `bson:"gender" json:"gender"`
	Address     string    `bson:"address" json:"address"`
	City        string    `bson:"city" json:"city"`
	State       string    `bson:"state" json:"state"`
	Pincode     string    `bson:"pincode" json:"pincode"`
}
