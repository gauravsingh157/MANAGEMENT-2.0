package store

import (
	"fmt"

	"MANAGEMENT-2.0/model"
	"MANAGEMENT-2.0/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

// func (store *Postgres) NewStore() error {
// 	dsn := "host=localhost user=gaurav password=gaurav12 dbname=singh port=5432 sslmode=disable"

// 	util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, " creating new store", nil)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, " error while creating ", err)
// 		return err
// 	}
// 	err = db.AutoMigrate(
// 		model.User{},
// 	)
// 	if err != nil {
// 		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, " error while running automigration", err)

// 		return err
// 	}
// 	fmt.Printf("db: %+v\n", db)
// 	return nil
// }

// type StoreOperation interface {
// 	NewStore() error
// 	CreatUser(user *model.User) error
// }

func (store *Postgres) NewStore() error {
	dsn := "host=localhost user=gaurav password=gaurav12 dbname=singh port=5432 sslmode=disable"
	util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, " creating new store", nil)

	var err error
	store.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, " error while creating ", err)
		return err
	}

	err = store.DB.AutoMigrate(model.User{})
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, " error while running automigration", err)
		return err
	}

	fmt.Printf("db: %+v\n", store.DB)
	return nil
}

type StoreOperation interface {
	NewStore() error
	CreatUser(user *model.User) error
}
