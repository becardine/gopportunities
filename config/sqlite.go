package config

import (
	"os"

	"github.com/becardine/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// check if db already exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database does not exist, creating new database...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		logger.Info("database created")
		file.Close()
	}

	// create db connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect database: %v", err)
		return nil, err
	}
	logger.Info("database connection established")

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("failed to migrate schema: %v", err)
		return nil, err
	}
	logger.Info("schema migrated")
	return db, nil
}
