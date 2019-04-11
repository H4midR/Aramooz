package controllers

import (
	db "Aramooz/dataBaseServices"
	"Aramooz/dataModels"
	"Aramooz/services/response"
	"encoding/json"

	"github.com/kataras/iris"
)

type QuestionController struct {
}

func (c *QuestionController) Get(ctx iris.Context) {
	ctx.WriteString("Add Question ..!")
}
func (c *QuestionController) PostBy(eid string, ctx iris.Context) response.Response {
	var req dataModels.Question
	var res response.Response

	err := ctx.ReadJSON(&req)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	res.Data = req
	return res
	dbwork := db.NewDgraphTrasn()
	query, err := json.Marshal(req)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	_, uids, err := dbwork.Mutate(query)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	res.Data = uids
	return res
}
