package main

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	store := newStore()
	router := gin.Default()

	addRoutes(router, store)
	AddRoutes(router)

	router.Run(":4000")
}
