package database

import (
	"mathgpt/configs/database/mongodb"
	"mathgpt/configs/database/mysql"

	"gorm.io/gorm"
)

var Database *DB

type DB struct {
	MyDB *gorm.DB
	MDB  *mongodb.MongoDB
}

func init() {
	Database.MDB = mongodb.Init()
	Database.MyDB = mysql.Init()

}
