package controllers

import (
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"time"
	"fmt"
	"strconv"
)

type Brainees struct {
	*revel.Controller
}

type NewBraineeReq struct {
   Text string `json: "text"`
   Author string `json: "author"`
   Brand string `json: "brand"`
}

type FullBrainees struct {
	ID int
	Text string
   	Author string
   	Brand string
}

// NewBrainee adds a brainee in the cache
func (c Brainees) NewBrainee() revel.Result {
	var reqData NewBraineeReq
    c.Params.BindJSON(&reqData)
	fmt.Println(reqData)

	var inCacheBrainees []FullBrainees

	if err := cache.Get("brainees", &inCacheBrainees); err != nil {
		var newCache = []FullBrainees{
			FullBrainees{
				ID: 0,
				Text: reqData.Text,
				Author: reqData.Author,
				Brand: reqData.Brand,
			},
		}
		cache.Set("brainees", newCache, 30*time.Minute)
	} else {
		fullNewBrainee := FullBrainees{
			ID: len(inCacheBrainees),
			Text: reqData.Text,
			Author: reqData.Author,
			Brand: reqData.Brand,
		}
		inCacheBrainees = append(inCacheBrainees, fullNewBrainee)
		fmt.Println(inCacheBrainees)
		cache.Set("brainees", inCacheBrainees, 30*time.Minute)
	}
	return c.RenderJSON("OK")
}

// GetBraineeByID returns a brainee based on the get data
func (c Brainees) GetBraineeByID() revel.Result{
	reqData, err := strconv.Atoi(c.Params.Route.Get("id"))
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON("Error, the id is not valid")
	}

	var inCacheBrainees []FullBrainees
	
	if errCache := cache.Get("brainees", &inCacheBrainees); errCache != nil {
		c.Response.Status = 500
	return c.RenderJSON("Can't get cache")
	} else {
		for _, brainee := range inCacheBrainees {
			if (brainee.ID == reqData) {
				return c.RenderJSON(brainee)
			}
		}
	}
	c.Response.Status = 500
	return c.RenderJSON("Can't find any brainee with this ID")
}//