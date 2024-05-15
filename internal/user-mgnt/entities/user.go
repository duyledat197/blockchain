package entities

type User struct {
	ID             string `bson:"_id,omitempty" json:"id,omitempty"`
	UserName       string `bson:"username,omitempty" json:"user_name,omitempty"`
	HashedPassword string `bson:"hashed_password,omitempty" json:"hashed_password,omitempty"`
}

// TableName returns the table name for the User entity.
//
// No parameters.
// Returns a string.
func (x *User) TableName() string {
	return "users"
}
