package service

import (
	"github.com/jcasado94/gobuyright/pkg/entity"
	"github.com/jcasado94/gobuyright/pkg/mongo"
	"github.com/jcasado94/gobuyright/pkg/mongo/model"
	mgo "gopkg.in/mgo.v2"
)

// ItemService serves the mongo operations for Item.
type ItemService struct {
	collection *mgo.Collection
}

// NewItemService creates a new ItemService given db and collection names.
func NewItemService(session *mongo.Session, dbName string, colName string) *ItemService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(model.ItemModelIndex())
	return &ItemService{collection}
}

// CreateItem inserts it into the collection.
func (is *ItemService) CreateItem(it *entity.Item) error {
	item := model.NewItemModel(it)
	return is.collection.Insert(&item)
}

// GetAllItems retrieves all Items from the collection.
func (is *ItemService) GetAllItems() ([]*entity.Item, error) {
	var ims []*model.ItemModel
	err := is.collection.Find(nil).All(&ims)
	if err != nil {
		return nil, err
	}
	var items []*entity.Item
	for _, im := range ims {
		items = append(items, im.ToItem())
	}
	return items, err
}
