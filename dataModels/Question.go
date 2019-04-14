package dataModels

type Choice struct {
	Uid   string `json:"uid,omitempty" form:"Uid"`
	Num   int    `json:"num,omirempty" form:"Num"`
	Title string `json:"title,omitempty" form:"Title"`
	Value bool   `json:"value,omitempty" form:"Value"`
}

type Question struct {
	Uid     string   `json:"uid,omitempty" from:"Uid"`
	Title   string   `json:"title,omitempty" form:"Title"`
	Qnum    string   `json:"qnum,omitempty" form:"Qnum"`
	Ltr     bool     `json:"ltr,omitempty" form:"Ltr"`
	Choices []Choice `json:"choices,omitempty" form:"Choices"`
}
