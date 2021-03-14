package tools

import (
	"strings"
)

/*
* String工具类
 */

/*
* 首字母大写，其余字母一律小写
 */
func Capitalize(str string) string {
	if len(str) < 1 {
		return ""
	}
	// 对其进行切片，转换成ASCII
	strArry := []rune(str)
	for index := range strArry {
		if index == 0 {
			// 首字母按大写处理
			if strArry[0] >= 97 && strArry[0] <= 122 {
				strArry[0] -= 32
			}
		} else {
			// 其它字母按小写处理
			if strArry[index] >= 65 && strArry[index] <= 90 {
				strArry[index] += 32
			}
		}
	}
	return string(strArry)
}

/*
* 根据表名生成实体对象名
* 采用驼峰命名的方式，首字母大写，后缀加上Entity
 */
func CreateEntityName(tableName string) string {
	// 切割字符串
	splitStr := strings.Split(tableName, "_")
	for index := range splitStr {
		// 首字母大写
		splitStr[index] = Capitalize(splitStr[index])
	}
	return strings.Join([]string(splitStr), "") + "Entity"
}

/*
* 根据字段名生成字段名
* 采用驼峰命名，首字母小写，_后的第一个字符大写
 */
func CreateFieldName(fieldName string) string {
	// 切割字符串
	splitStr := strings.Split(fieldName, "_")
	for index := range splitStr {
		// 第一个字符串全小写
		if index == 0 {
			splitStr[0] = strings.ToLower(splitStr[0])
		} else {
			// 首字母大写
			splitStr[index] = Capitalize(splitStr[index])
		}
	}
	return strings.Join([]string(splitStr), "")
}

/*
* 字段类型转换，
* 根据数据库字段属性，自动转换为，java基本数据类型
 */
func FieldTypeChange(fieldType string) string {
	// 全部转换为小写
	localType := strings.ToLower(fieldType)
	switch {
	case localType == "int", localType == "tinyint":
		return "Integer"
	case localType == "float":
		return "Float"
	case localType == "double":
		return "Double"
	case localType == "varchar", localType == "date", localType == "datetime":
		return "String"
	default:
		// 其它未定义类型
		return "String"
	}
}
