package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/singcl/go-vue/db"
)

type MemData struct {
	Data []float64 `binding: "required"`
}

type DbSchema struct {
	Id   uint
	Data []float64 `binding:"required"`
}

func init() {
	db.SetupDb()
}

func Persist(c *gin.Context) {
	data := new(MemData)
	c.Bind(&data)

	fmt.Printf("Persist数为：%+v", data)

	memDB := db.Database
	txn := memDB.Txn(true)

	p := &DbSchema{uint(1), data.Data}
	if err := txn.Insert("data", p); err != nil {
		panic(err)
	}

	txn.Commit()
}
