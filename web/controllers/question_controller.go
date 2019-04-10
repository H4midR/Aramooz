package controllers

import (
	db "Aramooz/dataBaseServices"
	"Aramooz/dataModels"
	"Aramooz/services/response"
	"encoding/json"

	"github.com/kataras/iris"
)

type AddQuestion struct {
}

func (c *AddQuestion) Post(ctx iris.Context) response.Response {
	var req dataModels.AddQuestion
	var res response.Response
	err := ctx.ReadJSON(&req)
	res.HandleErr(err)

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
