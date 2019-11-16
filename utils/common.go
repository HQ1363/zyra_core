package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Map map[string]interface{}

func (newMap Map) Merge(oldMap Map) {
	for k, v := range oldMap {
		newMap[k] = v
	}
}

/**
 * 根据path读取文件中的内容，返回字符串
 * 建议使用绝对路径，例如："./schema/search/appoint.json"
 */
func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func ReadJson(path string) Map {
	jsonStr := ReadFile(path)
	ret := Map{}
	err := json.Unmarshal([]byte(jsonStr), &ret)
	if err != nil {
		panic("文件[" + path + "]的内容不是json格式")
	}
	return ret
}

func JsonToMap(jsonStr string) Map {
	var mapResult Map
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMap err: ", err)
	}
	return mapResult
}

func StructToMap(source interface{}, withReflect bool) map[string]interface{} {
	target := make(Map)
	if withReflect {
		relVal := reflect.ValueOf(source)
		relType := reflect.TypeOf(source)
		if relType.Kind() != reflect.Struct {
			fmtString := fmt.Sprintf("待转换类型(%s)不为结构体类型, 参数错误.", relType.String())
			fmt.Println(fmtString)
			return target
		}
		for i := 0; i < relVal.NumField(); i++ {
			if relType.Field(i).Name == "BaseModel" {
				target.Merge(StructToMap(relVal.Field(i).Interface(), true))
			} else {
				key := relType.Field(i).Name
				target[key] = relVal.Field(i).Interface()
			}
		}
		return target
	}
	jsonStr, _ := json.Marshal(source)
	_ = json.Unmarshal(jsonStr, &target)
	return target
}

/**
 * 对象转换为string
 * 支持类型：int,float64,string,bool(true:"1";false:"0")
 * 其他类型报错
 */
func ToString(obj interface{}) string {
	switch obj.(type) {
	case int:
		return strconv.Itoa(obj.(int))
	case float64:
		return strconv.FormatFloat(obj.(float64), 'f', -1, 64)
	case string:
		return obj.(string)
	case bool:
		if obj.(bool) {
			return "1"
		} else {
			return "0"
		}
	default:
		panic("ToString出错")
	}
}

/**
 * 对象转换为bool
 * 支持类型：int,float64,string,bool
 * 其他类型报错
 */
func ToBool(obj interface{}) bool {
	switch obj.(type) {
	case int:
		if obj.(int) == 0 {
			return false
		} else {
			return true
		}
	case float64:
		if obj.(float64) == 0 {
			return false
		} else {
			return true
		}
	case string:
		trues := map[string]int{"true": 1, "是": 1, "1": 1, "真": 1}
		if _, ok := trues[strings.ToLower(obj.(string))]; ok {
			return true
		} else {
			return false
		}
	case bool:
		return obj.(bool)
	default:
		panic("ToBool出错")
	}
}

//判断一个数据是否为空，支持int, float, string, slice, array, map的判断
func Empty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		if reflect.ValueOf(value).Len() == 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

//判断某一个值是否在列表(支持 slice, array, map)中
func InList(needle interface{}, haystack interface{}) bool {
	//interface{}和interface{}可以进行比较，但是interface{}不可进行遍历
	hayValue := reflect.ValueOf(haystack)
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice, reflect.Array:
		//slice, array类型
		for i := 0; i < hayValue.Len(); i++ {
			if hayValue.Index(i).Interface() == needle {
				return true
			}
		}
	case reflect.Map:
		//map类型
		var keys []reflect.Value = hayValue.MapKeys()
		for i := 0; i < len(keys); i++ {
			if hayValue.MapIndex(keys[i]).Interface() == needle {
				return true
			}
		}
	default:
		return false
	}
	return false
}
