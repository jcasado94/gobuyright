package entity

// Usage represents a product usage.
type Usage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UsageService serves the DB queries for Usage.
type UsageService interface {
	CreateUsage(us *Usage) error
	GetByUsageID(usageID string) (*Usage, error)
	GetAllUsages() ([]*Usage, error)
}
