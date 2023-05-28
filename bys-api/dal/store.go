package dal

import (
	"bys/bootstrap"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	LocalDB *gorm.DB
)

func Init(cfg *bootstrap.Config) {
	MustInitMysql(cfg, "local")
}

func MustInitMysql(cfg *bootstrap.Config, name string) {
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: getMysqlDSN(cfg.Mysql[name]),
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "t_",
				SingularTable: true,
			},
		},
	)
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	LocalDB = db
}

func getMysqlDSN(conf bootstrap.MysqlConfig) string {
	return fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Addr, conf.DBName)
}
