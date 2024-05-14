package entities

type User struct {
	ID             string
	UserName       string
	HashedPassword string
}

// TableName returns the table name for the User entity.
//
// No parameters.
// Returns a string.
func (x *User) TableName() string {
	return "users"
}
