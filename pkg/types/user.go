package types

type User struct {
	ID          int64 `binding:"required"`
	FirstName   string
	LastName    string
	DisplayName string
	UserName    string
}
