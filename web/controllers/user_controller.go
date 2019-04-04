package controllers

import (
	db "Aramooz/dataBaseServices"

	"Aramooz/dataModels"
	"Aramooz/services/response"
	"encoding/json"

	"github.com/kataras/iris"
)

type UserController struct {
}

//Options : for allow crs
func (c *UserController) Options(ctx iris.Context) {}

//Get : get /user/
func (c *UserController) Get(ctx iris.Context) {

}

//Post : post /user/ : add new user
func (c *UserController) Post(ctx iris.Context) response.Response {
	var req dataModels.User
	var res response.Response
	err := ctx.ReadJSON(&req)

	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	req.Kind = dataModels.UserType
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
