package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err =
		gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/laravel")
	if err != nil {
		panic("Gagal Conect Ke Database")
	}
}

type (
	propinsi struct {
		ID       int    `json:"id"`
		Propinsi string `json:"propinsi"`
	}
)

func fetchAllPropinsi(c *gin.Context) {
	var model []propinsi

	db.Find(&model)

	if len(model) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data Tidak Ada"})
	}

	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": model})
}

func main() {

	router := gin.Default()
	v1 := router.Group("/api/propinsi")
	{
		v1.GET("", fetchAllPropinsi)
	}
	router.Run(":20001")
}
