package database

import (
	"database/jinzhu/gorm"
	_ "database/jinzu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)
