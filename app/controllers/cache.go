package controllers

import (
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
)

type Cache struct {
	*revel.Controller
}

func (c Cache) Index() revel.Result {
	var result map[string]interface{}
	if err := cache.Get("test", &result); err == nil {
		return c.RenderJSON(result)
	} else {
		return c.RenderJSON(err)
	}
}