package pegawai

import (
	"fmt"
	"math/rand"
	"os"
	"pegawaimicroservice/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"go.opencensus.io/trace"
)

type Pegawai struct {
	DB *gorm.DB
}

type Units struct {
	Unit []Unit `json:"data"`
}
type Unit struct {
	Id       int    `json:"id"`
	NamaUnit string `json:"nama_unit"`
	KodeUnit string `json:"kode_unit"`
}

func (p *Pegawai) GetPegawaiUnit(c *gin.Context) {
	db := p.DB
	var pegawai []model.Users
	var unitdata Units
	db.Find(&pegawai)

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	unitUrl := os.Getenv("UNIT_URL")
	unitPort := ":" + os.Getenv("UNIT_PORT")
	r := req.New()
	req.Debug = true
	resp, err := r.Get(unitUrl + unitPort + "/unit/") // nembak ke microservice unit
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error get product data " + err.Error(),
		})
	}
	resp.ToJSON(&unitdata) //convert ke json

	for i, peg := range pegawai {
		for _, dataunit := range unitdata.Unit {
			if peg.IdUnit == dataunit.Id { // apabila id unit di db = id unit dari microservice
				pegawai[i].NamaUnit = dataunit.NamaUnit // nama unit dimasukkan ke struct users (ada pada model)
			}
		}

	}
	c.JSON(200, gin.H{
		"pegawai": pegawai,
	})

}

func (p *Pegawai) GetPegawai(c *gin.Context) {
	db := p.DB
	var pegawai []model.Users

	db.Find(&pegawai)

	c.JSON(200, gin.H{
		"data": pegawai,
	})

	serviceb(c)
}

// func (p *Product) GetProductById(c *gin.Context) {
// 	var product model.Product
// 	db := p.DB
// 	id := c.Param("id")

// 	db.Where("id = ?", id).Find(&product)

// 	c.JSON(200, gin.H{
// 		"product_name": product.ProductName,
// 		"sku":          product.SKU,
// 		"qty":          product.Qty,
// 	})
// }

// func (p *Product) CreateProduct(c *gin.Context) {
// 	var request product

// 	if err := c.ShouldBind(&request); err != nil {
// 		c.JSON(500, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	p.DB.Create(&request)
// 	c.JSON(200, gin.H{
// 		"message": "success",
// 	})
// }

func serviceb(c *gin.Context) {
	_, span := trace.StartSpan(c, "/products")
	defer span.End()
	time.Sleep(time.Duration(rand.Intn(800)+200) * time.Millisecond)
}
