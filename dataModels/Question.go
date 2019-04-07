package dataModels

type Question struct{
	Uid 	string `json:"uid,omitempty" form:"Uid"`
	Title	string `json:"title,omitempty" form:"Title"`
	Cost	string `json:"cost,omitempty" form:"Cost"`
	Status	string `json:"status,omitempty" form:"Status"`
}

const QType = "Question"