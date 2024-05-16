package pubsub

// Pack represents the package data for a message. It includes key and value.
type Pack struct {
	Key []byte
	Msg []byte
}
