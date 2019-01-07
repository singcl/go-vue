package db

import memdb "github.com/hashicorp/go-memdb"

var Database *memdb.MemDB

func SetupDb() {
	// 创建数据库的 DBSchema , 它包括多个表结构：  Tables map[string]*TableSchema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			// 每个表结构是 TableSchema, 它定义对应的index
			"data": &memdb.TableSchema{
				Name: "data",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "Id"}, // 从代码来看，是基于 Id 来建立 index
					},
				},
			},
		},
	}

	// 基于schema 创建数据库
	Database, _ = memdb.NewMemDB(schema)
}
