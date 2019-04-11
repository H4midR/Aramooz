/****** SHAKOURI *****
Exam Controller defined in exam_controller
---------------------------------------
***********************/

package controllers

import (
	"github.com/kataras/iris"

	db "Aramooz/dataBaseServices"
	"Aramooz/dataModels"
	"Aramooz/services/response"
	"encoding/json"
)

//ExamController : http://URL:9090/exam
type ExamController struct{}

//
// ──────────────────────────────────────────────── I ──────────
//   :::::: E X A M : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────
//
func (c *ExamController) Options(ctx iris.Context) {}

func (c *ExamController) Get(ctx iris.Context) {}

func (c *ExamController) Put(ctx iris.Context)    {}

func (c *ExamController) Delete(ctx iris.Context) {}

func (c *ExamController) Post(ctx iris.Context) response.Response {
	var req dataModels.Exam
	var res response.Response
	err := ctx.ReadJSON(&req)

	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	req.Kind = dataModels.ExamType
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
//
// ──────────────────────────────────────────────────────── II ──────────
//   :::::: Q U E S T I O N : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────
//

//
// ──────────────────────────────────────────────────── III ──────────
//   :::::: C H O I C E : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────
//
