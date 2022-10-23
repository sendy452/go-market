package controller

import (
	"gotest_superindo/pkg"
	"gotest_superindo/structs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListKategori(c *gin.Context) {
	db, err := pkg.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	var kategori []structs.Kategori

	db.Table("tb_kategori").Select("*").Scan(&kategori)

	result := structs.ResponseKategori{
		Status: "success",
		Data:   kategori,
	}

	c.JSON(http.StatusOK, result)
}
