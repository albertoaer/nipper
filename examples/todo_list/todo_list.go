package main

import (
	"github.com/albertoaer/nipper"
	"github.com/gin-gonic/gin"
)

type Item struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func main() {
	items := make(map[string]Item)

	router := gin.Default()
	api := nipper.NewHttpRouter(router.Group("/api"))
	api.Route("items").Get(func() []Item {
		elems := make([]Item, 0)
		for _, v := range items {
			elems = append(elems, v)
		}
		return elems
	}).Post(func(item Item) int {
		if _, e := items[item.Name]; e {
			return 400
		} else {
			items[item.Name] = item
			return 200
		}
	})
	router.Run()
}
