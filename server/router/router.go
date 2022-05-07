package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jcbritobr/cstodo/model"
	"github.com/jcbritobr/cstodo/server/kvstore"
)

var (
	store = kvstore.NewKVStore()
)

func setupRoute(engine *gin.Engine) {
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	engine.POST("/insert", func(ctx *gin.Context) {
		var item model.Item
		err := ctx.BindJSON(&item)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%v", err)
			return
		}
		id := uuid.New()
		store.InsertItem(id.String(), item)
		ctx.JSON(200, gin.H{"uuid": id.String()})
	})

	engine.GET("/load", func(ctx *gin.Context) {
		uuid := ctx.DefaultQuery("uuid", "0")
		if uuid != "0" && len(uuid) != 0 {
			item, ok := store.LoadItem(uuid)
			if ok {
				ctx.JSON(http.StatusOK, item)
				return
			} else {
				ctx.JSON(http.StatusOK, model.ErrorMessage{Message: model.ErrNotFound})
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, model.ErrorMessage{Message: model.ErrEmptyQuery})
	})

	engine.GET("/list", func(ctx *gin.Context) {
		data := store.List()
		ctx.JSON(http.StatusOK, data)
	})

	engine.POST("/doneundone", func(ctx *gin.Context) {
		var message model.UuidMessage
		ctx.BindJSON(&message)
		result := store.DoneUndone(message.Uuid)
		ctx.JSON(200, model.DoneundoneResult{Done: result})
	})
}

func Boostrap() *gin.Engine {
	e := gin.Default()
	setupRoute(e)
	return e
}
