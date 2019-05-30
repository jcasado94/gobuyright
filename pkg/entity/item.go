package entity

// Item represents an Item. These hold the persisted information for every item used.
type Item struct {
	ID     string `json:"id"`
	ProdId string `json:"prodId"`
	Name   string `json:"name"`
	Img    string `json:"img"`
}

// ItemService serves the DB queries for Item.
type ItemService interface {
	GetAllItems() ([]*Item, error)
}
