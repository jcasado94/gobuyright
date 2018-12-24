package service

import (
	"gobuyright/pkg/entity"
	"gobuyright/pkg/mongo"
	"gobuyright/pkg/mongo/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UsageSelectionService serves the mongo operations for UsageSelection.
type UsageSelectionService struct {
	collection *mgo.Collection
}

// NewUsageSelectionService creates a new UsageSelectionService given database and collection names, connection through session.
func NewUsageSelectionService(session *mongo.Session, dbName string, colName string) *UsageSelectionService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(model.UsageSelectionModelIndex())
	return &UsageSelectionService{collection}
}

// CreateUsageSelection inserts us into the collection.
func (uss *UsageSelectionService) CreateUsageSelection(us *entity.UsageSelection) error {
	usageSelection := model.NewUsageSelectionModel(us)
	return uss.collection.Insert(&usageSelection)
}

// GetByUsernameAndTags retrieves the UsageSelection with correspondant username and tagIDs.
func (uss *UsageSelectionService) GetByUsernameAndTags(username string, tagIDs []string) (*entity.UsageSelection, error) {
	usm := model.UsageSelectionModel{}
	err := uss.collection.Find(bson.M{"username": username, "tagids": tagIDs}).One(&usm)
	if err != nil {
		return nil, err
	}
	return usm.ToUsageSelection(), err
}
