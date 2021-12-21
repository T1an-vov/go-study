package module

type Course struct {
	ID         uint
	Coursename string `form:"coursename" json:"coursename" binding:"required" gorm:"unique"` //课程名
	Credit     int    `form:"credit" json:"credit" bind:"required"` //学分
	Number     int `form:"number" json:"number"`  //目前选课人数
	Maxnumber  int `form:"maxnumber" json:"maxnumber" binding:"required"`  //最大选课人数
}
