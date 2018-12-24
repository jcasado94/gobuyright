package service

import (
	"gobuyright/pkg/entity"
	"gobuyright/pkg/mongo"
	"gobuyright/pkg/mongo/model"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// IUserService serves the mongo operations for User.
type IUserService struct {
	collection *mgo.Collection
}

// NewIUserService creates a new UserService given database and collection names, connection through session.
func NewIUserService(session *mongo.Session, dbName string, colName string) *IUserService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(model.IUserModelIndex())
	return &IUserService{collection}
}

// CreateUser inserts u into the collection.
func (us *IUserService) CreateUser(u *entity.IUser) error {
	user := model.NewIUserModel(u)
	return us.collection.Insert(&user)
}

// GetByUsername retrieves the User with Username username from the collection.
func (us *IUserService) GetByUsername(username string) (*entity.IUser, error) {
	um := model.IUserModel{}
	err := us.collection.Find(bson.M{"username": username}).One(&um)
	return um.ToIUser(), err
}
