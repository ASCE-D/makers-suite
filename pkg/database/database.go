package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	
	dsn := "postgresql://neondb_owner:hIFRT9Gb2ayr@ep-broad-bonus-a13difo7.ap-southeast-1.aws.neon.tech/neondb?sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}