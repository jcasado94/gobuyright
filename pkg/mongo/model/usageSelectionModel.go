package model

import (
	"github.com/jcasado94/gobuyright/pkg/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UsageSelectionModel is the DB model for entity.UsageSelectionModel.
type UsageSelectionModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	TagIDs   []string
	UsageIDs []string
}

// UsageSelectionModelIndex costructs the mgo.Index for UsageSelectionModel.
func UsageSelectionModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username", "tagids"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// NewUsageSelectionModel creates a new UsageSelectionModel given a UsageSelection.
func NewUsageSelectionModel(us *entity.UsageSelection) *UsageSelectionModel {
	return &UsageSelectionModel{
		Username: us.Username,
		TagIDs:   us.TagIDs,
		UsageIDs: us.UsageIDs,
	}
}

// ToUsageSelection creates an UsageSelection from the UsageSelectionModel.
func (us *UsageSelectionModel) ToUsageSelection() *entity.UsageSelection {
	return &entity.UsageSelection{
		ID:       us.ID.Hex(),
		Username: us.Username,
		TagIDs:   us.TagIDs,
		UsageIDs: us.UsageIDs,
	}
}
