package controllers

import (
	"github.com/kataras/iris"

	db "Aramooz/dataBaseServices"
	"Aramooz/dataModels"
	"Aramooz/services/response"
	"encoding/json"

	
)

type ExamController struct{

}

func(c *ExamController) Get (ctx iris.context){}
func(c *ExamController) Options (ctx iris.context){}
func(c *ExamController) Post (ctx iris.context) response.Response {
	var req dataModels.Exam
	var res response.Response
	err := ctx.ReadJSON(&req)

	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	req.Kind = dataModels.EType
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
func(c *ExamController) Put (ctx iris.context){}
func(c *ExamController) Delete (ctx iris.context){}