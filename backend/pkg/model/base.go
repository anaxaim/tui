package model

import (
	"time"
)

type BaseModel struct {
	CreatedAt       time.Time  `json:"createdAt" bson:"createdAt"`
	CreatedAtString string     `json:"createdAtString" bson:"createdAtString"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	UpdatedAtString string     `json:"updatedAtString,omitempty" bson:"updatedAtString,omitempty"`
}
