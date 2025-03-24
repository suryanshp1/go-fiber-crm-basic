package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


var (
	// DB is the database connection
	DBConn *gorm.DB
)