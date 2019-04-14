package controllers

import (
	db "Aramooz/dataBaseServices"
	"Aramooz/helperfunc"
	"Aramooz/services"

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

	if res.HandleErr(err) {
		return res
	}
	req.Kind = dataModels.UserType
	s, _ := dataModels.GeneratePassword(req.Password)
	req.Password = string(s)
	//
	// ──────────────────────────────── CHEACK NO OTHER USER WITH THE SAME MOBILE ─────
	//

	m := db.NewDgraphTrasn()
	if req.Mobile == "" || req.Mobile == "0" {
		res.Data = req
		res.Message = "فیلد موبایل خالی است"
		res.Code = -1
		res.State = -1
		return res
	}
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
func (c *UserController) Put(ctx iris.Context) response.Response {

	acl := services.Acl()
	res := services.Authentication(ctx, acl.ACLDef["User"], true)
	if res.Code < 1 {
		return res
	}

	myg := db.NewDgraphTrasn()

	var dbUser dataModels.User
	ctx.ReadJSON(&dbUser)

	dbUser.Uid = ctx.GetHeader("X-USER")

	dbqry, _ := json.Marshal(dbUser)
	_, _, err := myg.Mutate(dbqry)
	if res.HandleErr(err) {
		return res
	}

	fulluser, _ := c.getUserData(dbUser.Uid)
	var dbfuser struct {
		User []dataModels.User
	}
	if err := json.Unmarshal(fulluser, &dbfuser); err != nil {
		log.Print(err)
	}

	res = response.Response{
		Data:    dbfuser.User[0],
		Message: "ویرایش اطلاعات با موفقیت انجام شد",
		State:   1,
		Code:    1,
	}
	return res
	//UUID := ctx.GetHeader("X-USER")
}

//Delete : get /user/ : delete a user (delete account && adminSuperDelete)
//TODO: 0% -
// ?headers:(X-USER,TOKEN)
func (c *UserController) Delete(ctx iris.Context) {
}

//

//
// ────────────────────────────────────────────────────────────────────── II ──────────
//   :::::: S E S S I O N   S E T T I N G : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────────────────
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
			Message: "شماره موبایل یا کلمه عبور وارد شده نادرست است",
			Code:    -1,
		}
		return res
	}

	dbUser := dbres.User[0]

	// ? uncomment if using hash password
	if dataModels.CompairPassword(dbUser.Password, req.Password) == nil {
		// if dbUser.Password == req.Password {

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
		Message: "شماره موبایل یا کلمه عبور وارد شده نادرست است",
		Code:    -1,
	}
	return res
}

//
// ────────────────────────────────────────────────────────────────────── III ──────────
//   :::::: P R O F I L E   A N D   E X T : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────────────────
//

//PostProfile : Update User Profile
func (c *UserController) PostProfile(ctx iris.Context) response.Response {

	acl := services.Acl()
	res := services.Authentication(ctx, acl.ACLDef["User"], true)
	if res.Code < 1 {
		return res
	}

	myg := db.NewDgraphTrasn()

	var dbUser dataModels.User
	ctx.ReadJSON(&dbUser)

	dbUser.Uid = ctx.GetHeader("X-USER")

	dbqry, _ := json.Marshal(dbUser)
	_, _, err := myg.Mutate(dbqry)
	if res.HandleErr(err) {
		return res
	}

	fulluser, _ := c.getUserData(dbUser.Uid)
	var dbfuser struct {
		User []dataModels.User
	}
	if err := json.Unmarshal(fulluser, &dbfuser); err != nil {
		log.Print(err)
	}

	res = response.Response{
		Data:    dbfuser.User[0],
		Message: "ویرایش اطلاعات با موفقیت انجام شد",
		State:   1,
		Code:    1,
	}
	return res
}

//GetBy : get all data of user		- 	uid user id
//TODO : return Response vs []byte
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
