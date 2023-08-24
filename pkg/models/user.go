package models
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User - Data Model To Store user Data in
type user struct{
	Name string `json:"name"`
	City string `json:"city"`
	Age int `json:"age"`
	}