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

/*
This function must only be called once the function output has been validated
*/
func getOutputFunction(val reflect.Type) func([]reflect.Value, *gin.Context) {
	switch val.NumOut() {
	case 1:
		if val.Out(0).Kind() == reflect.Int {
			return func(out []reflect.Value, cnt *gin.Context) {
				cnt.Status(int(out[0].Int()))
			}
		} else {
			return func(out []reflect.Value, cnt *gin.Context) {
				cnt.JSON(200, out[0].Interface())
			}
		}
	case 2:
		if val.Out(0).Kind() == reflect.Int {
			return func(out []reflect.Value, cnt *gin.Context) {
				cnt.JSON(int(out[0].Int()), out[1].Interface())
			}
		} else {
			return func(out []reflect.Value, cnt *gin.Context) {
				cnt.JSON(int(out[1].Int()), out[0].Interface())
			}
		}
	default:
		return nil
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
	useOutput := getOutputFunction(tp)
	return func(c *gin.Context) {
		useOutput(
			val.Call(constructInput(c, tp.In, tp.NumIn())),
			c,
		)
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
