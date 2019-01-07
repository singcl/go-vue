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

// init函数import的时候会执行 详情查看init() main() 相关函数
func init() {
	db.SetupDb()
}

func Persist(c *gin.Context) {
	data := new(MemData)
	c.Bind(&data)

	fmt.Printf("Persist数为：%+v", data)

	memDB := db.Database

	// 创建写事务
	txn := memDB.Txn(true)
	p := &DbSchema{uint(1), data.Data}
	// 插入记录
	if err := txn.Insert("data", p); err != nil {
		panic(err)
	}
	// Commit
	txn.Commit()
}

// // 创建只读事务
// txn = db.Txn(false)
// defer txn.Abort()

// // 返回第一个符合的记录
// raw, err := txn.First("person", "id", "joe@aol.com")
// if err != nil {
//     panic(err)
// }

// // Say hi!
// fmt.Printf("Hello %s!", raw.(*Person).Name)
