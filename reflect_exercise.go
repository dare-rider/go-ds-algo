package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Gender struct {
	Data string `ms:"data"`
	Male bool `ms:"male"`
}

type CustumMapStruct struct {
	Name string `ms:"name"`
	Age int `ms:"age"`
	Gender *Gender `ms:"gen"`
}

//out kind should be pointer
//out kind should be pointer to struct
func mapToStruct(data map[string]interface{}, out interface{}) error {
	if out == nil {
		return errors.New("Nil pointer")
	}
	
	valAtPtr := reflect.ValueOf(out)
	if valAtPtr.Type().Kind() == reflect.Ptr {
		ele := valAtPtr.Elem()
		valAtPtr = reflect.ValueOf(ele.Addr().Interface()).Elem()
	}
	
	if valAtPtr.Kind() != reflect.Struct {
		return errors.New("Output is not struct")
	}
	
	tagMp := make(map[string]reflect.Value)
	fieldCount := valAtPtr.NumField()
	for i:=0; i < fieldCount; i+= 1 {
		fldValue := valAtPtr.Field(i)
		fldType := valAtPtr.Type().Field(i)
		tag := fldType.Tag.Get("ms")
		tagMp[tag] = fldValue
	}
	for key, val := range data {
		fld, ok := tagMp[key]
		if !ok {
			continue
		}
		switch fld.Kind() {
		case reflect.Int:
			fld.SetInt(int64(val.(int)))
		case reflect.String:
			fld.SetString(val.(string))
		case reflect.Bool:
			fld.SetBool(val.(bool))
		case reflect.Ptr:
			newEle := reflect.New(fld.Type().Elem())
			err := mapToStruct(val.(map[string]interface{}), newEle.Interface())
			if err != nil {
				return err
			}
			fld.Set(newEle)
		case reflect.Struct:
			newEle := reflect.New(fld.Type())
			err := mapToStruct(val.(map[string]interface{}), newEle.Interface())
			if err != nil {
				return err
			}
			fld.Set(newEle.Elem())
		default:
			return errors.New(fmt.Sprintf("Invalid type for field %v", fld.Type().Name()))
		}
	}
	return nil
}

func main() {
	data := map[string]interface{}{
			"name": "harry",
			"age": 28,
			"gen": map[string]interface{}{
				"data": "M",
				"male": true,
			},
		}
	res := CustumMapStruct{}
	err := mapToStruct(data, &res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(res.Gender)
}
