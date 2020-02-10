package database

// Config defines the database configuration
type Config struct {
	Hostname string
	Port     int
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
	Migrate  bool
}
