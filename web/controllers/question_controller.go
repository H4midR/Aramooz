package controllers

import (
	"Aramooz/dataModels"
	"Aramooz/services/response"

	"github.com/kataras/iris"
)

type AddQuestion struct {
}

func (c *AddQuestion) Post(ctx iris.Context) response.Response {
	var quest dataModels.AddQuestion
	var res response.Response
	err := ctx.ReadJSON(&quest)
	res.HandleErr(err)
	return res

}
