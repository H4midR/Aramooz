package Controllers

import (
	"Aramooz/DataBaseServices"
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
func (c *UserController) Post(ctx iris.Context) interface{} {
	var req dataModels.User
	err := ctx.ReadJSON(&req)
	if err != nil {
		return err
	}
	req.Kind = dataModels.UserType
	mgt := DataBaseServices.NewDgraphTrasn()
	q, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, Uids, err := mgt.Mutate(q)
	if err != nil {
		return err
	}
	return Uids

}
