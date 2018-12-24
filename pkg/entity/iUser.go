package entity

// IUser represents an IUser. These will hold the persisted information for every user from generic-filler.
type IUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// IUserService serves the DB queries for IUser.
type IUserService interface {
	CreateUser(u *IUser) error
	GetByUsername(username string) (*IUser, error)
}
