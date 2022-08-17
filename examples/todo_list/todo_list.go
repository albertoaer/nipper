package main

import (
	"github.com/albertoaer/nipper"
	"github.com/gin-gonic/gin"
)

type Item struct {
	Name string `json:"name" param:"name"`
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
		}
		items[item.Name] = item
		return 200
	})
	api.Route("items/:name").Get(func(ref Item) (int, *Item) {
		if item, e := items[ref.Name]; e {
			return 200, &item
		}
		return 404, nil
	})
	router.Run()
}
