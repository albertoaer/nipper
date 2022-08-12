package nipper

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

/*
A nipper endpoint can have as input values as it wants as long they are structs,
they will be filled with data from the request

The response of the endpoint must be either nothing,
a status, a struct or both
*/

func validateInput(val reflect.Type) bool {
	for i := 0; i < val.NumIn(); i++ {
		if val.In(i).Kind() != reflect.Struct {
			return false
		}
	}
	return true
}

func validateOutput(val reflect.Type) bool {
	validOut := func(x reflect.Type) bool {
		return x.Kind() == reflect.Int || x.Kind() == reflect.Struct
	}
	switch val.NumOut() {
	case 1:
		return validOut(val.Out(0))
	case 2:
		return val.Out(0).Kind() != val.Out(1).Kind() && validOut(val.Out(0)) && validOut(val.Out(1))
	default:
		return false
	}
}

func injectFunction(val reflect.Value) gin.HandlerFunc {
	tp := val.Type()
	if !validateInput(tp) {
		panic("Invalid input from endpoint handler")
	}
	if !validateOutput(tp) {
		panic("Invalid output from endpoint handler")
	}
	return func(c *gin.Context) {
		//TODO: Use return type
		val.Call(constructInput(c, tp.In, tp.NumIn()))
	}
}

func injectInto(
	routeFunc func(string, ...gin.HandlerFunc) gin.IRoutes,
	route Route,
	handler interface{},
) {
	val := reflect.ValueOf(handler)
	switch val.Type().Kind() {
	case reflect.Func:
		routeFunc(route.Route(), injectFunction(val))
	default:
		panic("Type not supported as endpoint handler")
	}
}
