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
func (c *QuestionController) PostBy(userId string, ctx iris.Context) response.Response {
	var req dataModels.Question
	var res response.Response
	req.Kind = dataModels.QuestionType
	err := ctx.ReadJSON(&req)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
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
	qu := `{
			set{
				<` + userId + `> <uid> <` + uids["blank-0"] + `>
			}
		}`
	query2, err2 := json.Marshal(qu)
	res.HandleErr(err2)
	if res.Code < 1 {
		return res
	}
	_, uids2, err2 := dbwork.Mutate(query2)
	res.HandleErr(err2)
	if res.Code < 1 {
		return res
	}
	res.Data = uids2
	return res
}
