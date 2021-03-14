package datasource

import (
	"iris-starter/tools"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

var (
	masterEngine *xorm.Engine
	lock         sync.Mutex
)

// 使用单例模式创建数据源
func InstanceMaster() *xorm.Engine {
	// TODO:
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()
	// 双重判断
	if masterEngine != nil {
		return masterEngine
	}
	//导入配置文件
	config := tools.InitConfig("./conf/application.properties")
	// 开发路径
	//config := tools.InitConfig("./conf/application.properties")
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		config["User"], config["Pwd"], config["Host"], config["Port"], config["DbName"], config["CharSet"])
	engine, err := xorm.NewEngine(config["DriverName"], driverSource)
	engine.ShowSQL(true)
	// 空闲连接池数
	engine.SetMaxIdleConns(3)
	// 最大连接池数
	engine.SetMaxOpenConns(10)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster error=", err)
		return nil
	} else {
		masterEngine = engine
		return masterEngine
	}
}