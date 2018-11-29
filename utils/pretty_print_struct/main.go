package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func prettyPrintStruct(i interface{}) {
	params := Flatten(i)
	var buffer bytes.Buffer
	for i := range params {
		buffer.WriteString("\n\t")
		buffer.WriteString(params[i])
	}
	fmt.Printf("Orderer config values:%s\n", buffer.String())
}

func Flatten(i interface{}) []string {
	var res []string
	flatten("", &res, reflect.ValueOf(i))
	return res
}

const DELIMITER = "."

func flatten(k string, m *[]string, v reflect.Value) {
	delimiter := DELIMITER
	if k == "" {
		delimiter = ""
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			*m = append(*m, fmt.Sprintf("%s =", k))
			return
		}
		flatten(k, m, v.Elem())
	case reflect.Struct:
		if x, ok := v.Interface().(fmt.Stringer); ok {
			*m = append(*m, fmt.Sprintf("%s = %v", k, x))
			return
		}

		for i := 0; i < v.NumField(); i++ {
			flatten(k+delimiter+v.Type().Field(i).Name, m, v.Field(i))
		}
	case reflect.String:
		// It is useful to quote string values
		*m = append(*m, fmt.Sprintf("%s = \"%s\"", k, v))
	default:
		*m = append(*m, fmt.Sprintf("%s = %v", k, v))
	}
}

type Person struct {
	Name  string
	Age   int
	Other Other
}

type Other struct {
	like []string
}

func main() {
	s1 := []string{"abc", "edf", "hig", "klm", "nop"}
	person := Person{"Simeon", 21, Other{s1}}
	prettyPrintStruct(person)

	// fmt.Println(reflect.ValueOf(person))
	// fmt.Println(person)
}
