package route

import (
	"gotest_superindo/controller"
	"gotest_superindo/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	api := router.Group("api/v1")
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)

		secured := api.Use(middleware.Middleware())
		{
			secured.GET("/listkategori", controller.GetListKategori)

			secured.GET("/:kategori/listproduk", controller.GetProdukByKategori)
			secured.GET("/listproduk/:id", controller.GetProdukById)

			secured.POST("/cart", controller.PostToCart)
		}
	}

	return router
}
