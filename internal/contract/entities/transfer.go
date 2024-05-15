package entities

type Transfer struct {
	ID    string `bson:"_id,omitempty" json:"id,omitempty"`
	From  string `bson:"from,omitempty" json:"from,omitempty"`
	To    string `bson:"to,omitempty" json:"to,omitempty"`
	Value string `bson:"value,omitempty" json:"value,omitempty"`

	Timestamp   int64  `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	BlockNumber uint64 `bson:"block_number,omitempty" json:"block_number,omitempty"`
}

func (x *Transfer) TableName() string {
	return "transfers"
}
