package main

import (
	"github.com/gin-gonic/gin"

	"getAllUuid/handler"
	"getAllUuid/repository"
)

func main() {
	client, ctx, cancel := repository.ConnectDB()
	defer cancel()
	defer client.Disconnect(ctx)

	router := gin.Default()
	router.GET("/uuids", handler.GetUUIDsHandler(client))
	router.Run("localhost:7070")
}
