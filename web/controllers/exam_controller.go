package controllers

import (
	"github.com/kataras/iris"

	db "Aramooz/dataBaseServices"
	"Aramooz/dataModels"
	"Aramooz/services/response"
	"encoding/json"

	
)

type QuestionController struct{

}

func(c *QuestionController) Get (ctx iris.context){}
func(c *QuestionController) Options (ctx iris.context){}
func(c *QuestionController) Post (ctx iris.context) response.Response {
	var req dataModels.Question
	var res response.Response
	err := ctx.ReadJSON(&req)

	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	req.Kind = dataModels.QType
	mgt := db.NewDgraphTrasn()
	q, err := json.Marshal(req)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	_, Uids, err := mgt.Mutate(q)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	res.Data = Uids
	return res
}
func(c *QuestionController) Put (ctx iris.context){}
func(c *QuestionController) Delete (ctx iris.context){}