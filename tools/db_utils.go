package tools

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
* 数据库工具类
 */

func GetDBConnetion(userName string, passWord string, ipAddr string, port string, DBName string, tableName string) ([]map[string]interface{}, error) {
	log := GetLoggerInstance().Logger
	// 避免没有来得及运行就结束了，强制flush
	defer log.Flush()
	// 构造数据库连接url
	db_url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", userName, passWord, ipAddr, port, DBName)
	DB, err := sql.Open("mysql", db_url)
	if err != nil {
		log.Warnf("数据库连接发生错误: %s", err)
		return nil, err
	}
	defer DB.Close()

	rows, err := DB.Query("select column_name as columnName,data_type as dataType,column_comment as columnComment from information_schema.columns where table_name=?", tableName)
	if err != nil {
		log.Warnf("数据库查询失败: %s", err)
		return nil, err
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength) //临时存储每行数据
	for index, _ := range cache {              //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		list = append(list, item)
	}
	return list, nil
}
