package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserName       string             `bson:"username,omitempty" json:"user_name,omitempty"`
	HashedPassword string             `bson:"hashed_password,omitempty" json:"hashed_password,omitempty"`

	PrivateKey    string `bson:"private_key,omitempty" json:"private_key,omitempty"`
	WalletAddress string `bson:"wallet_address,omitempty" json:"wallet_address,omitempty"`
	Nonce         string `bson:"nonce,omitempty" json:"nonce,omitempty"`
}

// TableName returns the table name for the User entity.
//
// No parameters.
// Returns a string.
func (x *User) TableName() string {
	return "users"
}
