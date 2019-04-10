// file: services/poster_service.go
/*eslint-disable */
package services

import (
	db "Aramooz/dataBaseServices"
	"Aramooz/dataModels"
	"Aramooz/helperfunc"
	"Aramooz/services/response"
	"encoding/json"
	"fmt"
)

type UserService interface {
	BasicOuth(string, string) response.Response
	GetAcl() map[string]map[string]AclVal
}

type userService struct {
	key   string
	Uid   string
	UID   uint64
	Token string
}

func NewUserService(uid string, token string) *userService {
	return &userService{
		key:   "poster",
		Uid:   uid,
		Token: token,
	}
}

func GetAcl(uid string) map[string]AclVal {
	myg := db.NewDgraphTrasn()
	q := fmt.Sprintf(`
		{
			user(func: uid(%s)) @filter(eq(key,"user")) {
				acl
			}
		}
		`, uid)

	resb, _ := myg.Query(q)
	var resstrc struct {
		User []dataModels.User `json:"user"`
	}
	json.Unmarshal(resb, &resstrc)
	useraclval, _ := helperfunc.UID(resstrc.User[0].ACL)
	acl := Acl()
	return acl.Privileges(useraclval)

}

//BasicOuth : outhenticate the user with uid and token
func BasicOuth(uid string, token string) (response.Response, bool, string) {
	Uid, err := helperfunc.UIDStrX(uid)
	if err != nil {
		res := response.Response{
			Message: err.Error(),
			State:   -1,
			Code:    -1,
		}
		return res, false, ""
	}
	myg := db.NewDgraphTrasn()
	q := fmt.Sprintf(`
	 	{
	 		user(func: uid(%s)) @filter(eq(kind,"User")) {
				uid
				token
				
				acl
	 		}
	 	}
		 `, Uid)
	dbresstr, _ := myg.Query(q)
	var dbres struct {
		User []dataModels.User
	}
	//ctx.Write(dbresstr)
	if err := json.Unmarshal(dbresstr, &dbres); err != nil {
		res := response.Response{
			Message: err.Error(),
			State:   -1,
			Code:    -1,
		}
		return res, false, ""
	}
	if len(dbres.User) < 1 {
		res := response.Response{
			State:   -1,
			Message: "Invalid Username Or Password",
			Code:    -1,
		}
		return res, false, ""
	}
	dbUser := dbres.User[0]
	if dbUser.Token == token {
		res := response.Response{
			State: 1,
			Code:  1,
		}
		return res, true, dbUser.ACL
	} else {
		res := response.Response{
			State: -1,
			Code:  -1,
		}
		return res, false, ""
	}

}
