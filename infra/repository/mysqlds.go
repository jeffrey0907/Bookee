package repository

import (
    "Bookee/infra/config"
    "Bookee/infra/logger"
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "strings"
    "sync"
)

var (
    mysqldb     *gorm.DB
    onceMysqldb sync.Once
)

func DB() *gorm.DB {
    onceMysqldb.Do(func() {
        activedb := config.GetStringOrEmpty(`datasource.active`)

        mysqlFunc := func() *gorm.DB {
            url := config.GetStringOrEmpty(`datasource.mysql.url`)
            um := config.GetStringOrEmpty(`datasource.mysql.username`)
            pwd := config.GetStringOrEmpty(`datasource.mysql.password`)
            db, err := gorm.Open("mysql", fmt.Sprint(`%s:%s@%s`, um, pwd, url))
            if err != nil {
                logger.L().Println(err.Error())
            } else {
                maxOpenConns := config.GetInt(`datasource.mysql.maxopenconns`, 5)
                db.DB().SetMaxOpenConns(maxOpenConns)
                maxIdleConns := config.GetInt(`datasource.mysql.maxidleconns`, 2)
                db.DB().SetMaxIdleConns(maxIdleConns)
            }
            return db
        }
        sqliteFunc := func() *gorm.DB {
            url := config.GetStringOrEmpty(`datasource.sqlite.url`)
            db, err := gorm.Open("sqlite3", url)
            if err == nil {
                return db
            } else {
                panic(err)
            }
        }

        if strings.EqualFold(`sqlite`, activedb) {
            mysqldb = mysqlFunc()
        } else {
            mysqldb = sqliteFunc()
        }
    })
    return mysqldb
}
