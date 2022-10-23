package controller

import (
	"gotest_superindo/pkg"
	"gotest_superindo/structs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostToCart(c *gin.Context) {
	db, err := pkg.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	var cart structs.Cart

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	addCart := structs.Cart{Id_Produk: cart.Id_Produk, Qty: cart.Qty, Id_User: cart.Id_User}

	if err := db.Table("tb_cart").Create(&addCart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := structs.Response{
		Status:  "success",
		Message: "Berhasil ditambahkan ke cart",
	}

	c.JSON(http.StatusOK, response)
}
