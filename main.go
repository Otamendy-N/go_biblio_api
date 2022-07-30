package main

import (
	"go_biblio_api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controllers.ConfigureRouter(router)

	router.Run("localhost:5000")
}
