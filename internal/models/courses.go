package models


type Courses struct{
	CourseName string	`json:"coursename"`
	CourseCode string	`json:"coursecode"`
	Credits    int		`json:"Credits"`
}