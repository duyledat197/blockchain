package config

import (
	"fmt"
	"net/url"
)

// Database represents the configuration details for a database connection.
type Database struct {
	Schema        string `mapstructure:"schema,omitempty"`
	Host          string `mapstructure:"host,omitempty"`
	Port          string `mapstructure:"port,omitempty"`
	User          string `mapstructure:"user,omitempty"`
	Password      string `mapstructure:"password,omitempty"`
	Database      string `mapstructure:"database,omitempty"`
	MaxConnection int32  `mapstructure:"max_connection,omitempty"`
}

// Address returns the formatted string for the database connection address.
func (e *Database) Address() string {
	var uInfo *url.Userinfo
	if e.User != "" && e.Password != "" {
		uInfo = url.UserPassword(e.User, e.Password)
	}
	dbURL := &url.URL{
		Scheme: e.Schema, // or "mysql", "sqlite3", etc.
		User:   uInfo,
		Host:   fmt.Sprintf("%s:%s", e.Host, e.Port),
		Path:   e.Database,
	}

	query := dbURL.Query()
	query.Set("sslmode", "disable") // or "require", "verify-full", etc.
	dbURL.RawQuery = query.Encode()

	return dbURL.String()
}
