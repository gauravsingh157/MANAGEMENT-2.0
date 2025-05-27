// package controller

// import (
// 	"net/http"
// 	"time"

// 	"MANAGEMENT-2.0/model"
// 	"MANAGEMENT-2.0/util"
// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// func (server Server) CreatUser(ctx *gin.Context) {

// 	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreatUser, " creating new user", nil)

// 	user := model.User{}
// 	err := ctx.BindJSON(&user)
// 	if err != nil {
// 		util.Log(model.LogLevelError, model.ControllerPackage, model.CreatUser, " error while json to binding ", err)
// 		//ctx.JSON(http.StatusInternalServerError, err)
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

// 		return
// 	}
// 	user.ID = uuid.New()
// 	user.CreatedAt = time.Now()

// 	err = server.PostgressDb.CreatUser(&user)
// 	if err != nil {
// 		util.Log(model.LogLevelError, model.ControllerPackage, model.CreatUser, " error while inserting record", err)
// 		ctx.JSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, user)

// }
package controller

import (
	"fmt"
	"net/http"
	"time"

	"MANAGEMENT-2.0/model"
	"MANAGEMENT-2.0/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server Server) CreatUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreatUser, " creating new user", nil)

	user := model.User{}
	//err := ctx.BindJSON(&user)
	err := ctx.ShouldBindJSON(&user)

	// if err != nil {
	// 	util.Log(model.LogLevelError, model.ControllerPackage, model.CreatUser, " error while json to binding", err)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	// 	return

	// }
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreatUser, "error while binding JSON", err)
		fmt.Println("Step 1: Binding success, user =", user)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400 bad request is better here
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	err = server.PostgressDb.CreatUser(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreatUser, " error while inserting record", err)
		//ctx.JSON(http.StatusInternalServerError, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	ctx.JSON(http.StatusCreated, user)
}
