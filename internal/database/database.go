package database

import (
	"github.com/arkurl/mygo-todo/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.Conf.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Conf.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.Conf.Database.MaxLifeTime)

	return db, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
