package entities

type Approval struct {
	ID      string `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   string `bson:"owner,omitempty" json:"owner,omitempty"`
	Spender string `bson:"spender,omitempty" json:"spender,omitempty"`
	Value   string `bson:"value,omitempty" json:"value,omitempty"`

	Timestamp   int64  `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	BlockNumber uint64 `bson:"block_number,omitempty" json:"block_number,omitempty"`
}

// TableName returns the name of the table for Approval entities.
func (x *Approval) TableName() string {
	return "approvals"
}
