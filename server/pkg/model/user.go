package model

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Password  string             `json:"password" bson:"password"`
	AuthInfos []AuthInfo         `json:"authInfos" gorm:"foreignKey:UserID;references:ID"`

	BaseModel
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

type AuthInfo struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID       string             `json:"userId" bson:"userId"`
	URL          string             `json:"url" bson:"url"`
	AuthType     string             `json:"authType" bson:"authType"`
	AuthID       string             `json:"authId" bson:"authId"`
	AccessToken  string             `json:"-" bson:"-"`
	RefreshToken string             `json:"-" bson:"-"`
	Expiry       time.Time          `json:"-" bson:"-"`

	BaseModel
}

type CreatedUser struct {
	Name      string     `json:"name" bson:"name"`
	Password  string     `json:"password" bson:"password"`
	AuthInfos []AuthInfo `json:"authInfos" bson:"authInfos"`
}

func (u *CreatedUser) GetUser() *User {
	return &User{
		Name:      u.Name,
		Password:  u.Password,
		AuthInfos: u.AuthInfos,
	}
}

type UpdatedUser struct {
	Name      string     `json:"name" bson:"name"`
	Password  string     `json:"password" bson:"password"`
	AuthInfos []AuthInfo `json:"authInfos" bson:"authInfos"`
}

func (u *UpdatedUser) GetUser() *User {
	return &User{
		Name:      u.Name,
		Password:  u.Password,
		AuthInfos: u.AuthInfos,
	}
}

type AuthUser struct {
	Name      string `json:"name" bson:"name"`
	Password  string `json:"password" bson:"password"`
	SetCookie bool   `json:"setCookie" bson:"setCookie"`
	AuthType  string `json:"authType" bson:"authType"`
	AuthCode  string `json:"authCode" bson:"authCode"`
}

type Users []User
