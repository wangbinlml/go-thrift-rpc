package core

import (
	"reflect"
	io "io/ioutil"
	"github.com/astaxie/beego/logs"
)

func Switch(value reflect.Value) {
	kind := value.Kind()
	switch kind {
	case reflect.String:
		logs.Info("string: ", value.String())
	case reflect.Struct:
		parseStruct(value)
	case reflect.Slice:
		parseSlice(value)
	case reflect.Interface:
		v := reflect.ValueOf(value.Interface())
		Switch(v)
	default:
		logs.Info("unknown type: ", kind.String())
	}
}
func parseSlice(value reflect.Value) {
	size := value.Len()
	for j := 0; j < size; j++ {
		field := value.Index(j)
		Switch(field)
	}
}
func parseStruct(value reflect.Value) {
	size := value.NumField()
	for i := 0; i < size; i++ {
		field := value.Field(i)
		Switch(field)
	}
}

func LoadFile(filename string) []byte {
	data, err := io.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	datajson := []byte(data)
	return datajson
}


//Usege
/*
var objs []interface{}
if err := json.Unmarshal([]byte(txt), &objs); err != nil {
fmt.Println(err)
return
}
for _, obj := range objs {
Switch(reflect.ValueOf(obj))
}*/