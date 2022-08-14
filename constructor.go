package nipper

import (
	"encoding/json"
	"reflect"

	"github.com/gin-gonic/gin"
)

type FieldSetter = func(val reflect.Value, data string)

func updateFieldParam(val reflect.Value, data string) {
	//TODO: Adapt all golang types
	val.SetString(data)
}

func updateFieldQuery(val reflect.Value, data []string) {
	//TODO: Adapt all golang types
	val.SetString(data[0])
}

type structField struct {
	structIdx int
	fieldIdx  int
}

func getInputConstructor(get func(int) reflect.Type, length int) func(*gin.Context) []reflect.Value {
	paramSetters := make(map[string][]structField)
	querySetters := make(map[string][]structField)
	for i := 0; i < length; i++ {
		tp := get(i)
		for f := 0; f < tp.NumField(); f++ {
			if p := tp.Field(f).Tag.Get("param"); len(p) > 0 {
				paramSetters[p] = append(paramSetters[p], structField{i, f})
			}
			if q := tp.Field(f).Tag.Get("query"); len(q) > 0 {
				querySetters[q] = append(querySetters[q], structField{i, f})
			}
		}
	}
	return func(c *gin.Context) []reflect.Value {
		decoder := json.NewDecoder(c.Request.Body)
		res := make([]reflect.Value, length)
		for i := 0; i < length; i++ {
			res[i] = reflect.New(get(i)).Elem()
			decoder.Decode(res[i].Interface())
		}
		for _, v := range c.Params {
			for _, sf := range paramSetters[v.Key] {
				updateFieldParam(res[sf.structIdx].Field(sf.fieldIdx), v.Value)
			}
		}
		for k, v := range c.Request.URL.Query() {
			for _, sf := range querySetters[k] {
				updateFieldQuery(res[sf.structIdx].Field(sf.fieldIdx), v)
			}
		}
		return res
	}
}
