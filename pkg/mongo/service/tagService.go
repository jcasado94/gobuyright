package service

import (
	"gobuyright/pkg/entity"
	"gobuyright/pkg/mongo"
	"gobuyright/pkg/mongo/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TagService serves the mongo operations for Tag
type TagService struct {
	collection *mgo.Collection
}

// NewTagService creates a new TagService given database and collection names, connection through session.
func NewTagService(session *mongo.Session, dbName string, colName string) *TagService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(model.TagModelIndex())
	return &TagService{collection}
}

// CreateTag inserts t into the collection.
func (ts *TagService) CreateTag(t *entity.Tag) error {
	tag := model.NewTagModel(t)
	return ts.collection.Insert(&tag)
}

// GetTagByID retrieves the Tag with ID tagID from the collection.
func (ts *TagService) GetTagByID(tagID string) (*entity.Tag, error) {
	tm := model.TagModel{}
	err := ts.collection.Find(bson.M{"tagID": tagID}).One(&tm)
	if err != nil {
		return nil, err
	}
	return tm.ToTag(), err
}
