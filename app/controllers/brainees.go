package controllers

import (
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"time"
	"strconv"
	"techTest/app/models"
)

type Brainees struct {
	*revel.Controller
}

type NewBraineeReq struct {
   Text string `json: "text"`
   Author string `json: "author"`
   Brand string `json: "brand"`
}

// NewBrainee adds a brainee in the cache
func (c Brainees) NewBrainee() revel.Result {
	var reqData NewBraineeReq
    c.Params.BindJSON(&reqData)

	_, err := models.PostBrainee(reqData.Author, reqData.Text, reqData.Brand)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON("Error")
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

	var inCacheBrainees []models.Brainee
	
	if errCache := cache.Get("brainees", &inCacheBrainees); errCache == nil {
		for _, brainee := range inCacheBrainees {
			if (brainee.ID == reqData) {
				return c.RenderJSON(brainee)
			} 
		}
	} 

	dbBrainee, err := models.FindBrainee(reqData)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(err)
	} else {
		inCacheBrainees = append(inCacheBrainees, dbBrainee)
		cache.Set("brainees", inCacheBrainees, 30*time.Minute)
		
		return c.RenderJSON(dbBrainee)
	}

	
	c.Response.Status = 500
	return c.RenderJSON("Can't find any brainee with this ID")
}