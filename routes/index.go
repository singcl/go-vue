package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gonum/stat"
	"github.com/singcl/go-vue/db"
)

// MemData is POST 的数据结构体
type MemData struct {
	Data []float64 `binding:"required"` // tag key:value中间不需要空格
}

// DbSchema 数据存入数据库之前的处理
type DbSchema struct {
	Id   uint
	Data []float64 `binding:"required"`
}

// init 函数 import 的时候会执行 详情查看init() main() 相关函数
func init() {
	db.SetupDb()
}

// Persist 路由handler 处理请求数据
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

// Mean is a router handler
// 重启后 内存数据库中数据就没了。必须重启生成数据（调用上面的接口）该接口才会正确返回
func Mean(c *gin.Context) {
	memDb := db.Database

	// 创建只读事务
	txn := memDb.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("data", "id", uint(1))
	if err != nil {
		panic(err)
	}

	mean := stat.Mean(raw.(*DbSchema).Data, nil)

	c.JSON(200, mean)
}
