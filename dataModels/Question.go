package dataModels

type Choice struct {
	Uid   string `json:"uid,omitempty" form:"Uid"`
	Num   int    `json:"num,omirempty" form:"Num"`
	Title string `json:"title,omitempty" form:"Title"`
	Value int    `json:"value,omitempty" form:"Value"`
	Kind  string `choices`
}

type Question struct {
	Uid     string   `json:"uid,omitempty" from:"Uid"`
	Title   string   `json:"title,omitempty" form:"Title"`
	Qnum    string   `json:"qnum,omitempty" form:"Qnum"`
	Ltr     bool     `json:"ltr,omitempty" form:"Ltr"`
	Kind    string   `json:"kind,omitempty"`
	Choices []Choice `json:"choices,omitempty" form:"Choices"`
}

const QuestionType = "question"
const ChoiceType = "choices"
