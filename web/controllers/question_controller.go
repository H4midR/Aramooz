package controllers

import (
	db "Aramooz/dataBaseServices"
	"Aramooz/dataModels"
	"Aramooz/services/response"
	"encoding/json"

	"github.com/kataras/iris"
)

type QuestionController struct {
}

func (c *QuestionController) Get(ctx iris.Context) {
	ctx.WriteString("Add Question ..!")
}
func (c *QuestionController) PostBy(userId string, ctx iris.Context) response.Response {
	var req dataModels.Question
	var res response.Response
	req.Kind = dataModels.QuestionType
	err := ctx.ReadJSON(&req)
	Freq := dataModels.Exam{
		Uid: userId, //TODO: ExamUID from form

	}
	Freq.Questions = append(Freq.Questions, req)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	dbwork := db.NewDgraphTrasn()
	query, err := json.Marshal(Freq)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	_, uids, err := dbwork.Mutate(query)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}

	/*r1 := fmt.Sprintf(`<%s> <owner> <%s> .
					`, userId, uids["blank-0"])
	dbwork.MutateRDF([]byte(r1), "set")
	*/
	/*
		addOwner := `{
					set{
						<` + userId + `> <owner> <` + uids["blank-0"] + ` >
						}
					}`
	*/
	addOwner := `{
				"uid" : "` + userId + `"
				"owner" :{
					"uid": "` + uids["blank-0"] + `"
				}
			}`
	addQuery, err := json.Marshal(addOwner)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	_, finalUids, err := dbwork.Mutate(addQuery)
	res.HandleErr(err)
	if res.Code < 1 {
		return res
	}
	res.Data = finalUids
	return res
}
