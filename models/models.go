package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
var DB orm.Ormer
func init() {
	orm.Debug = true
	dbHost := beego.AppConfig.String("dbhost")
	dbPort := beego.AppConfig.String("dbport")
	dbUser := beego.AppConfig.String("dbuser")
	dbPassword := beego.AppConfig.String("dbpassword")
	dbName := beego.AppConfig.String("dbname")
	if dbPort == "" {
		dbPort = "3306"
	}
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Error(err)
	}
	err2 := orm.RegisterDataBase("default", "mysql", dsn)
	if err2 != nil {
		logs.Error(err2)
	}


	orm.RegisterModel(new(Users), new(Topic), new(SourceCode),
		new(Solution), new(Sim), new(ShareCode), new(Reply),
		new(Problem), new(Privilege), new(Printer), new(Online),
		new(Mail), new(News), new(LoginLog), new(Custominput),
		new(ContestProblem), new(Contest), new(Compileinfo), new(Balloon))
	_ = orm.RunSyncdb("default", false, true)

	DB = orm.NewOrm()
	DB.Using("default") // 默认使用 default，你可以指定为其他数据库

}