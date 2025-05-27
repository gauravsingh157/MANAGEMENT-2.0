// package store

// import (
// 	"MANAGEMENT-2.0/model"
// 	"MANAGEMENT-2.0/util"
// )

// func (store Postgres) CreatUser(user *model.User) error {

// 	util.Log(model.LogLevelInfo, model.StorePackage, model.CreatUser, " creating new user", nil)

// 	response := store.DB.Create(user) //यहां store.DB मतलब GORM का database object है।

// 	if response.Error != nil {

// 		util.Log(model.LogLevelInfo, model.StorePackage, model.CreatUser, " error while creating new user ", response.Error)
// 		return response.Error

// 	}
// 	util.Log(model.LogLevelInfo, model.StorePackage, model.CreatUser, " creating new user", user)
// 	return nil

// }
package store

import (
	"MANAGEMENT-2.0/model"
	"MANAGEMENT-2.0/util"
)

func (store Postgres) CreatUser(user *model.User) error {
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreatUser, " creating new user", nil)
	response := store.DB.Create(user) //यहां store.DB मतलब GORM का database object है।

	if response.Error != nil {

		util.Log(model.LogLevelInfo, model.StorePackage, model.CreatUser, " error while creating new user ", response.Error)
		return response.Error

	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreatUser, " creating new user", user)
	return nil

}
