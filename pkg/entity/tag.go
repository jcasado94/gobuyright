package entity

// Tag represents a product tag.
type Tag struct {
	ID      string `json:"id"`
	TagID   string `json:"tagID"`
	TagName string `json:"tagName"`
}

// TagService serves the DB queries for Tag.
type TagService interface {
	CreateTag(t *Tag) error
	GetTagByID(tagID string) (*Tag, error)
}
