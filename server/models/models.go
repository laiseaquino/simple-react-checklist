package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// model for how data will be stored in database
// (uppercase because it will be exported)
type Checklist struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Number string             `json:"number,omitempty"`
	Status bool               `json:"status,omitempty"`
}
