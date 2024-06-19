package initialize

import (
	"fmt"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/tools/gormx/dmclient"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

const (
	DefaultLogName = "gorm"
)

func InitGorm() {
	if global.DBCli != nil {
		return
	}

	switch global.DBConfig.Type {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
			global.DBConfig.Username,
			global.DBConfig.Password,
			global.DBConfig.Host,
			global.DBConfig.Port,
			global.DBConfig.DBName,
			true,
			"Local")

		gormDial(mysql.Open(dsn))

	case "dm":
		//return fmt.Sprintf("dm://ops:123456789@192.168.119.82:5236/ops?columnNameCase=lower&charset=utf8mb4&parseTime=True&loc=Local")
		dsn := fmt.Sprintf("dm://%s:%s@%s:%d/%s?columnNameCase=lower&charset=utf8mb4&parseTime=True&loc=Local",
			global.DBConfig.Username,
			global.DBConfig.Password,
			global.DBConfig.Host,
			global.DBConfig.Port,
			global.DBConfig.DBName)
		gormDial(dmclient.Open(dsn))
	}
}

func gormDial(dial gorm.Dialector) {
	var err error
	global.DBCli, err = gorm.Open(dial, &gorm.Config{
		//为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）
		//如果没有这方面的要求，可以设置SkipDefaultTransaction为true来禁用它。
		//SkipDefaultTransaction: true,
		//Logger: Log,
		//执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续执行的效率
		PrepareStmt: global.DBConfig.PrepareStmt,
		NamingStrategy: schema.NamingStrategy{
			//使用单数表名,默认为复数表名，即当model的结构体为User时，默认操作的表名为users
			//设置	SingularTable: true 后当model的结构体为User时，操作的表名为user
			SingularTable: true,
		},
	})

	if err != nil {
		logx.Fatal(err)
	}

	global.DBCli.Set("gorm:table_options", "CHARSET=utf8mb4")

	settingDB, err := global.DBCli.DB()
	if err != nil {
		logx.Fatal(err)
	}

	settingDB.SetMaxOpenConns(global.DBConfig.MaxOpenConn)
	settingDB.SetMaxIdleConns(global.DBConfig.MaxIdleConn)
	settingDB.SetConnMaxLifetime(time.Duration(global.DBConfig.ConnMaxLifeSecond) * time.Minute)

	err = global.DBCli.Callback().Create().After("gorm:after_create").Register(DefaultLogName, afterLog)
	if err != nil {
		logx.Errorf("Register Create error, %s", err)
	}
	err = global.DBCli.Callback().Query().After("gorm:after_query").Register(DefaultLogName, afterLog)
	if err != nil {
		logx.Errorf("Register Query error", err)
	}
	err = global.DBCli.Callback().Update().After("gorm:after_update").Register(DefaultLogName, afterLog)
	if err != nil {
		logx.Errorf("Register Update error", err)
	}
	err = global.DBCli.Callback().Delete().After("gorm:after_delete").Register(DefaultLogName, afterLog)
	if err != nil {
		logx.Errorf("Register Delete error", err)
	}
}

func afterLog(db *gorm.DB) {
	err := db.Error
	//ctx := db.Statement.Context
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	if err != nil {
		logx.Error(sql, err)
	} else {
		logx.Infof("[ SQL语句: %s]", sql)
	}
}
