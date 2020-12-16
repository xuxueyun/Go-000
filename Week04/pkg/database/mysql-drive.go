package database

import (
	"Week04/global"
	"Week04/pkg/config"
	"Week04/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/os/glog"
	"github.com/jinzhu/gorm"
)

type Mysql struct {
}

func (e *Mysql) Setup() {
	var err error
	var db Database

	db = new(Mysql)
	global.Source = db.GetConnect()
	global.Eloquent, err = db.Open(db.GetDriver(), db.GetConnect())
	if err != nil {
		glog.Fatal(utils.Red(db.GetDriver()+" connect error :"), err)
	} else {
		glog.Info(utils.Green(db.GetDriver() + " connect success."))
	}

	if global.Eloquent.Error != nil {
		glog.Fatal(utils.Red(" database error :"), global.Eloquent.Error)
	}
}

// Open 打开数据库连接
func (e *Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	return gorm.Open(dbType, conn)
}

// GetConnect 获取数据库连接
func (e *Mysql) GetConnect() string {
	return config.DatabaseConfig.Source
}

// GetDriver 获取 Driver
func (e *Mysql) GetDriver() string {
	return config.DatabaseConfig.Driver
}
