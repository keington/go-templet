package zlog

import (
	"encoding/xml"
	"fmt"
	"os"
	"reflect"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/20 21:08
 * @file: log_cfg_xml.go
 * @description: 读取xml配置文件, 并将配置文件的值赋值给结构体
 */

// Xml read and load xml config file
// path: xml file path
// values: struct
func Xml(path string, values interface{}) interface{} {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error reading file:", err)
		return LogConfig{}
	}

	err = xml.Unmarshal(data, &values)
	if err != nil {
		fmt.Printf("error: %v", err)
		return LogConfig{}
	}

	// 反射获取结构体的字段
	t := reflect.TypeOf(values)
	v := reflect.ValueOf(values)

	if v.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Println("values is not a struct")
		return nil
	}

	for i := 0; i < t.NumField(); i++ {
	}

	return values
}
