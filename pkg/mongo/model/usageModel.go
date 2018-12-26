package model

import (
	"github.com/jcasado94/gobuyright/pkg/entity"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UsageModel is the DB model for entity.Usage.
type UsageModel struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UsageID   string
	UsageName string
}

// UsageModelIndex constructs the mgo.Index for usageModel
func UsageModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"usageID"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// NewUsageModel creates a new UsageModel given a Usage.
func NewUsageModel(u *entity.Usage) *UsageModel {
	return &UsageModel{
		UsageID:   u.UsageID,
		UsageName: u.UsageName,
	}
}

// ToUsage creates an Usage from the UsageModel.
func (um *UsageModel) ToUsage() *entity.Usage {
	return &entity.Usage{
		ID:        um.ID.Hex(),
		UsageID:   um.UsageID,
		UsageName: um.UsageName,
	}
}
