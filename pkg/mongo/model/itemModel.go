package model

import (
	"github.com/jcasado94/gobuyright/pkg/entity"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ItemModel is the DB model for entity.Item
type ItemModel struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	ProdId string        `bson:"prodId"`
	Name   string        `bson:"name"`
	Img    string        `bson:"img"`
}

// ItemModelIndex constructs the mgo.Index for itemModel
func ItemModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"ID"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// NewItemModel creates a new ItemModel given an Item.
func NewItemModel(i *entity.Item) *ItemModel {
	return &ItemModel{
		ProdId: i.ProdId,
		Name:   i.Name,
		Img:    i.Img,
	}
}

// ToItem creates an Item from the ItemModel.
func (im *ItemModel) ToItem() *entity.Item {
	return &entity.Item{
		ID:     im.ID.Hex(),
		ProdId: im.ProdId,
		Name:   im.Name,
		Img:    im.Img,
	}
}
