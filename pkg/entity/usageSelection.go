package entity

// UsageSelection represents a user's usage categorization on a product tag (set).
type UsageSelection struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	TagIDs   []string `json:"tagIDs"`
	UsageIDs []string `json:"usageIDs"`
}

// UsageSelectionService serves the DB queries for Usage.
type UsageSelectionService interface {
	CreateUsageSelection(us *UsageSelection) error
	GetByUsernameAndTags(username string, tagIDs []string) (*UsageSelection, error)
}
