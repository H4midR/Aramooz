package dataModels

type Exam struct {
	Uid                   string     `json:"uid,omitempty" form:"Uid"`
	Title                 string     `json:"title,omitempty" form:"Title"`
	Cost                  string     `json:"cost,omitempty" form:"Cost"`
	ExamStatus            string     `json:"examStatus,omitempty" form:"examStatus"`
	SelectCourse          string     `json:"selectcourse" form:"selectcourse"`
	NumberofFinalQuestion string     `json:"numberofFinalQuestion" form:"numberofFinalQuestion"`
	Duration              string     `json:"duration" form:"duration"`
	StartDate             string     `json:"startDate" form:"startDate"`
	EndDate               string     `json:"endDate" form:"endDate"`
	startTime             string     `json:"startTime" form:"startTime"`
	EndTime               string     `json:"endTime" form:"endTime"`
	BeforeMessage         string     `json:"beforeMessage" form:"beforeMessage"`
	AfterMessage          string     `json:"afterMessage" form:"afterMessage"`
	MultiCourse           string     `json:"multiCourse" form:"multiCourse"`
	SpecificTime          string     `json:"specificTime" form:"specificTime"`
	RandomShow            string     `json:"randomShow" form:"randomShow"`
	NegativeScore         string     `json:"negativeScore" form:"negativeScore"`
	Kind                  string     `json:"kind,omitempty"`
	Questions             []Question `json:"question,omitempty"`
}

const ExamType = "Exam"
