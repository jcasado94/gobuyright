package model

import (
	"gobuyright/pkg/entity"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// IUserModel is the DB model for entity.IUser.
type IUserModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string
}

// IUserModelIndex constructs the mgo.Index for userModel
func IUserModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// NewIUserModel creates a new UserModel given a IUser.
func NewIUserModel(u *entity.IUser) *IUserModel {
	return &IUserModel{
		Username: u.Username,
	}
}

// ToIUser creates an IUser from the IUserModel.
func (um *IUserModel) ToIUser() *entity.IUser {
	return &entity.IUser{
		ID:       um.ID.Hex(),
		Username: um.Username,
	}
}
