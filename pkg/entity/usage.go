package entity

// Usage represents a product usage.
type Usage struct {
	ID        string `json:"id"`
	UsageID   string `json:"usageId"`
	UsageName string `json:"usageName"`
}

// UsageService serves the DB queries for Usage.
type UsageService interface {
	CreateUsage(us *Usage) error
	GetByUsageID(usageID string) (*Usage, error)
	GetAllUsages() []*Usage
}
