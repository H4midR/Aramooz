package controllers

import (
	db "Aramooz/dataBaseServices"
	"Aramooz/helperfunc"
	"errors"
	"fmt"
	"log"

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

//

//PostLogin : sign in
func (c *UserController) PostLogin(ctx iris.Context) response.Response { // must cheack uniq mobile
	//UID, _ := strconv.ParseUint(uid, 16, 64)
	var req dataModels.User // make(map[string]string)
	ctx.ReadJSON(&req)
	var res response.Response
	myg := db.NewDgraphTrasn()
	q := fmt.Sprintf(`
	 	{
	 		user(func: eq(mobile,"%s")) @filter(eq(kind,"User")) {
				uid
				password
				token
	 		}
	 	}
		 `, req.Mobile)

	dbresstr, err := myg.Query(q)
	if res.HandleErr(err) {
		return res
	}
	var dbres struct {
		User []dataModels.User
	}
	//ctx.Write(dbresstr)
	if err := json.Unmarshal(dbresstr, &dbres); err != nil {
		log.Print(err)
	}

	if len(dbres.User) < 1 {
		res = response.Response{
			State:   -1,
			Message: "Invalid Username Or Password",
			Code:    -1,
		}
		return res
	}

	dbUser := dbres.User[0]

	// ? uncomment if using hash password
	//if  datamodels.CompairPassword(dbUser.Password, req.Password) == nil {
	if dbUser.Password == req.Password {

		dbres.User[0].Token = helperfunc.Tokengenerator()
		dbqry, _ := json.Marshal(dbres)
		myg.Mutate(dbqry)
		/*var UData struct {
			Uid   string
			Token string
			User  datamodels.User
		}
		UData.Token = dbres.User[0].Token
		UData.Uid = dbres.User[0].Uid
		UData.User= dbres.User[]*/

		fulluser, _ := c.getUserData(dbres.User[0].Uid)
		var dbfuser struct {
			User []dataModels.User
		}
		if err := json.Unmarshal(fulluser, &dbfuser); err != nil {
			log.Print(err)
		}

		res = response.Response{
			Data:    dbfuser.User[0],
			Message: "ورود موفقیت آمیز",
			State:   1,
			Code:    1,
		}
		return res

	}
	res = response.Response{
		State:   -1,
		Message: "Invalid Username Or Password",
		Code:    -1,
	}
	return res
}

//GetBy : get all data of user		- 	uid user id
func (c *UserController) getUserData(Uid string) ([]byte, error) {
	//res,c := services.BasicOuth()
	myg := db.NewDgraphTrasn()
	q := fmt.Sprintf(`
		{
			user(func: uid(%s)) @filter(eq(kind,"User")) {
				uid
				expand(_all_)
				
			}
		}
		`, Uid)

	return myg.Query(q)
}
