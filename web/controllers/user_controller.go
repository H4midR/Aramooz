package controllers

import (
	db "Aramooz/dataBaseServices"

	"Aramooz/dataModels"
	"Aramooz/services/response"
	"encoding/json"

	"github.com/kataras/iris"
)

//UserController : http://URL:9090/user
type UserController struct {
}

//
// ──────────────────────────────────────────────── I ──────────
//   :::::: U S E R : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────
//

// ─── CURD ───────────────────────────────────────────────────────────────────────

//Options : for allow crs
func (c *UserController) Options(ctx iris.Context) {}

//Get : get /user/ : show user data
//TODO: 0% -
// ? headers:(X-USER,TOKEN)
func (c *UserController) Get(ctx iris.Context) {
}

//Post : post /user/ : add new user
//TODO: 50% -
// *TODO : uniq user checking , recaptcha , required param checking , ...
// ? headers:(NULL)
// ? @Params(mobile,password,name)
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

//Put : Put /user/ : edit a user
//TODO: 0% -
// ? headers:(X-USER,TOKEN)
func (c *UserController) Put(ctx iris.Context) {
}

//Delete : get /user/ : delete a user (delete account && adminSuperDelete)
//TODO: 0% -
// ?headers:(X-USER,TOKEN)
func (c *UserController) Delete(ctx iris.Context) {
}
