// file: services/poster_service.go
/*eslint-disable */
package services

import (
	
	"Aramooz/helperfunc"
	db "Aramooz/dataBaseServices"
	"Aramooz/dataModels"
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
	myg := NewMygraphService()
	q := fmt.Sprintf(`
		{
			user(func: uid(%s)) @filter(eq(key,"user")) {
				acl
			}
		}
		`, uid)

	resb := myg.Query(q)
	var resstrc struct {
		User []datamodels.User `json:"user"`
	}
	json.Unmarshal(resb, &resstrc)
	useraclval, _ := helperfunc.UID(resstrc.User[0].ACL)
	acl := Acl()
	return acl.Privileges(useraclval)

}

//BasicOuth : outhenticate the user with uid and token
func BasicOuth(uid string, token string) (Response, bool, string) {
	Uid, err := helperfunc.UIDStrX(uid)
	if err != nil {
		res := Response{
			Error:  err.Error(),
			Status: "Error",
			Code:   InvalidUID,
		}
		return res, false, ""
	}
	myg := NewMygraphService()
	q := fmt.Sprintf(`
	 	{
	 		user(func: uid(%s)) @filter(eq(kind,"User")) {
				uid
				token
				
				acl
	 		}
	 	}
		 `, Uid)
	dbresstr := myg.Query(q)
	var dbres struct {
		User []datamodels.User
	}
	//ctx.Write(dbresstr)
	if err := json.Unmarshal(dbresstr, &dbres); err != nil {
		res := Response{
			Error:  err.Error(),
			Status: "Error",
			Code:   UserOuthFail,
		}
		return res, false, ""
	}
	if len(dbres.User) < 1 {
		res := Response{
			Status: "Error",
			Error:  "Invalid Username Or Password",
			Code:   UserOuthFail,
		}
		return res, false, ""
	}
	dbUser := dbres.User[0]
	if dbUser.Token == token {
		res := Response{
			Status: "OK",
			Code:   UserOuthSucc,
		}
		return res, true, dbUser.ACL
	} else {
		res := Response{
			Status: "Error",
			Code:   UserOuthFail,
		}
		return res, false, ""
	}

}
