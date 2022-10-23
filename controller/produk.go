package controller

import (
	"gotest_superindo/pkg"
	"gotest_superindo/structs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProdukByKategori(c *gin.Context) {
	db, err := pkg.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	kategori := c.Param("kategori")

	var produk []structs.Produk

	db.Table("tb_produk").Select("tb_produk.*, tb_kategori.kategori").Joins("left join tb_kategori on tb_kategori.id_kategori = tb_produk.id_kategori").Where("tb_produk.id_kategori = ?", kategori).Scan(&produk)

	result := structs.ResponseProduk{
		Status: "success",
		Data:   produk,
	}

	c.JSON(http.StatusOK, result)
}

func GetProdukById(c *gin.Context) {
	db, err := pkg.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	id := c.Param("id")

	var produk []structs.Produk

	db.Table("tb_produk").Select("tb_produk.*, tb_kategori.kategori").Joins("left join tb_kategori on tb_kategori.id_kategori = tb_produk.id_kategori").Where("tb_produk.id_produk = ?", id).Scan(&produk)

	result := structs.ResponseProduk{
		Status: "success",
		Data:   produk,
	}

	c.JSON(http.StatusOK, result)
}
