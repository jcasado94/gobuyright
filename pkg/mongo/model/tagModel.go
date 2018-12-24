package model

import (
	"gobuyright/pkg/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TagModel is the DB model for entity.Tag.
type TagModel struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	TagID   string
	TagName string
}

// TagModelIndex constructs the mgo.Index for tagModel.
func TagModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"tagID"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// NewTagModel creates a new TagModel given a Tag.
func NewTagModel(t *entity.Tag) *TagModel {
	return &TagModel{
		TagID:   t.TagID,
		TagName: t.TagName,
	}
}

// ToTag creates an Tag from the TagModel.
func (t *TagModel) ToTag() *entity.Tag {
	return &entity.Tag{
		ID:      t.ID.Hex(),
		TagID:   t.TagID,
		TagName: t.TagName,
	}
}
