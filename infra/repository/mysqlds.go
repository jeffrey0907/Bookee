package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var (
	mysqldb      *gorm.DB
	mysqldbMutex sync.Mutex
)

func DB() *gorm.DB {
	if mysqldb == nil {
		func() {
			mysqldbMutex.Lock()
			defer mysqldbMutex.Unlock()
			if mysqldb == nil {
				db, err := gorm.Open("mysql")
				if err == nil {
					mysqldb = db
				}
			}
		}()
	}
	return mysqldb
}
