package service

import (
	"github.com/jcasado94/gobuyright/pkg/entity"
	"github.com/jcasado94/gobuyright/pkg/mongo"
	"github.com/jcasado94/gobuyright/pkg/mongo/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UsageService serves the mongo operations for Usage.
type UsageService struct {
	collection *mgo.Collection
}

// NewUsageService creates a new UsageService given database and collection names, connection through session.
func NewUsageService(session *mongo.Session, dbName string, colName string) *UsageService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(model.UsageModelIndex())
	return &UsageService{collection}
}

// CreateUsage inserts u into the collection.
func (us *UsageService) CreateUsage(u *entity.Usage) error {
	usage := model.NewUsageModel(u)
	return us.collection.Insert(&usage)
}

// GetByUsageID retrieves the Usage with ID usageID from the collection.
func (us *UsageService) GetByUsageID(usageID string) (*entity.Usage, error) {
	um := model.UsageModel{}
	err := us.collection.Find(bson.M{"usageID": usageID}).One(&um)
	if err != nil {
		return nil, err
	}
	return um.ToUsage(), err
}

// GetAllUsages retrieves all Usages from the collection.
func (us *UsageService) GetAllUsages() ([]*entity.Usage, error) {
	var ums []*model.UsageModel
	err := us.collection.Find(nil).All(&ums)
	if err != nil {
		return nil, err
	}
	var usages []*entity.Usage
	for _, um := range ums {
		usages = append(usages, um.ToUsage())
	}
	return usages, err
}
