package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (gaurav *Postgres) NewStore() error {
	dsn := "host=localhost user=gaurav password=gaurav12 dbname=singh port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	gaurav.DB = db
	fmt.Printf("db: %+v\n", db)
	return nil
}

type StoreOperation interface {
	NewStore()error
}
