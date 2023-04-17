package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credential struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	CreatedBy   string             `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	Secrets     Secrets            `json:"secrets,omitempty" bson:"secrets,omitempty"`

	BaseModel
}

type Secret struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

type (
	Credentials []Credential
	Secrets     []Secret
)
