package nipper

import (
	"encoding/json"
	"reflect"

	"github.com/gin-gonic/gin"
)

func updateFieldValue(val reflect.Value, raw string) {
	//TODO: Adapt all golang types
	val.SetString(raw)
}

func constructInput(c *gin.Context, get func(int) reflect.Type, last int) []reflect.Value {
	decoder := json.NewDecoder(c.Request.Body)
	res := make([]reflect.Value, last)
	for i := 0; i < last; i++ {
		tp := get(i)
		res[i] = reflect.New(tp).Elem()
		for f := 0; f < tp.NumField(); f++ {
			if p := tp.Field(f).Tag.Get("param"); len(p) > 0 && len(c.Param(p)) > 0 {
				updateFieldValue(res[i].Field(f), c.Param(p))
			}
			if q := tp.Field(f).Tag.Get("query"); len(q) > 0 && len(c.Query(q)) > 0 {
				updateFieldValue(res[i].Field(f), c.Query(q))
			}
		}
		decoder.Decode(res[i].Interface())
	}
	return res
}
