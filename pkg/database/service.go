package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	// Postres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	monthNamesDE = [12]string{
		"Jänner",
		"Februar",
		"März",
		"April",
		"Mai",
		"Juni",
		"Juli",
		"August",
		"September",
		"Oktober",
		"November",
		"Dezember",
	}
)

// Service is the database service
type Service struct {
	db *gorm.DB
}

// NewService returns the database handler as DataAccessor. It also opens the database connection
// using the given config.
func NewService(config Config) *Service {

	var s Service
	var err error
	if config.Dialect == "postgres" {
		dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.Hostname,
			config.Port,
			config.Username,
			config.Name,
			config.Password)
		fmt.Println(dbURI)
		s.db, err = gorm.Open(config.Dialect, dbURI)
	}
	if err != nil {
		log.Fatal("Could not connect database:", err)
	}

	if config.Migrate {
		log.Println("Starting DB migration ...")
		s.db = Migrate(s.db)
	}

	return &s
}
