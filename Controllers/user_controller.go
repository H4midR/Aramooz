package Controllers

import (
	"Aramooz/DataBaseServices"
	"Aramooz/Services/Response"
	"Aramooz/dataModels"
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
func (c *UserController) Post(ctx iris.Context) Response.Response {
	var req dataModels.User
	var res Response.Response
	err := ctx.ReadJSON(&req)

	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	req.Kind = dataModels.UserType
	mgt := DataBaseServices.NewDgraphTrasn()
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
