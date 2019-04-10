package controllers

import (
	"Aramooz/dataModels"
	"Aramooz/services/response"

	"github.com/kataras/iris"
)

type QuestionController struct {
}

func (c *QuestionController) Post(ctx iris.Context) response.Response {
	var req dataModels.Question
	var res response.Response
	err := ctx.ReadJSON(&req)
	res.HandleErr(err)
	return res
}
