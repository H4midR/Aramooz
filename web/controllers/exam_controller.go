/****** SHAKOURI *****
Exam Controller defined in exam_controller
---------------------------------------
***********************/

package controllers

import (
	"github.com/kataras/iris"

	
	"fmt"
	

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


//GetBy : get all data of Exam		- 	uid exam id
func (c *ExamController) getExamData(UID string) ([]byte, error) {
	//res,c := services.BasicOuth()
	myg := db.NewDgraphTrasn()
	q := fmt.Sprintf(`
		{
			exam(func: uid(%s)) @filter(eq(kind,"Exam")) {
				uid
				expand(_all_)
				
			}
		}
		`, UID)

	return myg.Query(q)
}
