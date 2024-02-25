package models

type Test struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	Id       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	TestId   string `json:"testId"`
}

type Enrollment struct {
	StudentId string `json:"studentId"`
	TestId    string `json:"testId"`
}
