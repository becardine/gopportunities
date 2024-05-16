package handler

import (
	"github.com/becardine/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Initialize() error {
	logger = config.GetLogger("handler")
	var err error
	db, err = config.InitializeSQLite()
	if err != nil {
		return err
	}
	return nil
}
