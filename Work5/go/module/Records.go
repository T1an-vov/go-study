package module

type Record struct {
	ID uint
	Username string `form:"username" json:"username"`  //选课人
	Coursename string `form:"course" json:"course"` //所选课程
}
