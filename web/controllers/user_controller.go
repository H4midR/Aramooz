package controllers

import (
	db "Aramooz/dataBaseServices"
	"errors"
	"fmt"

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
	//
	// ──────────────────────────────── CHEACK NO OTHER USER WITH THE SAME MOBILE ─────
	//

	m := db.NewDgraphTrasn()
	dbStrRes, err := m.Query(fmt.Sprintf(`
		{
			data(func:eq(mobile,"%s")){
				uid
			}
		}
	`, req.Mobile))
	if res.HandleErr(err) {
		return res
	}
	var dbStc struct {
		Data []struct {
			Uid string `json:"uid"`
		} `json:"data"`
	}
	err = json.Unmarshal(dbStrRes, &dbStc)
	if res.HandleErr(err) {
		return res
	}
	if len(dbStc.Data) > 0 {
		res.HandleErr(errors.New("قبلا کاربری با این شماره موبایل ثبت نام کرده است."))
		return res
	}

	//
	// ───────────────────────────────────────────────────── ADD USER TO DATABASE ─────
	//

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
	//UUID := ctx.GetHeader("X-USER")
}

//Delete : get /user/ : delete a user (delete account && adminSuperDelete)
//TODO: 0% -
// ?headers:(X-USER,TOKEN)
func (c *UserController) Delete(ctx iris.Context) {
}
