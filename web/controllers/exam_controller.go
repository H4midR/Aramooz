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

func (c *ExamController) Get() string{
	return "Hello" ;
}

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
func (c *ExamController) GetList() response.Response {
	//res,c := services.BasicOuth()
	var res response.Response
	myg := db.NewDgraphTrasn()
	q := fmt.Sprintf(`
		{
			exams(func:eq(kind,"Exam")){
				expand(_all_)
			  }
		}
		`)

	dbres,err := myg.Query(q)
	if res.HandleErr(err) {
		return res
		}
	var dbexams struct{
		Exams []dataModels.Exam `json:"exams"`
	}
	/*err = json.Unmarshal(dbres,&dbexams)
	if res.HandleErr(err) {
		return res
	}*/
	
	
	if err := json.Unmarshal(dbres, &dbexams); res.HandleErr(err) {
		return res
		}
	if len(dbexams.Exams) < 1 {
		res = response.Response{
			State:   -1,
			Message: "اطلاعات بدرستی بارگذاری نشده است",
			Code:    -1,
			}
		return res
		}
		res.Code=1
		res.Message="داده های امتحان"
		res.State=1
		res.Data = dbexams.Exams
		return res
	}
