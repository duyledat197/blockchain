package config

import "fmt"

// Database represents the configuration details for a database connection.
type Database struct {
	Host          string `mapstructure:"host,omitempty"`
	Port          string `mapstructure:"port,omitempty"`
	User          string `mapstructure:"user,omitempty"`
	Password      string `mapstructure:"password,omitempty"`
	Database      string `mapstructure:"database,omitempty"`
	MaxConnection int32  `mapstructure:"max_connection,omitempty"`
}

// Address returns the formatted string for the database connection address.
func (e *Database) Address() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		e.Host,
		e.Port,
		e.User,
		e.Password,
		e.Database,
	)
}
