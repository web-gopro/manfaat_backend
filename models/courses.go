package models

type CoursesReq struct {
	Title       string `json:"title"`
	Teacher_id  string `json:"teacher_id"`
	Description string `json:"description"`
}

type Course struct {
	Course_id   string `json:"course_id"`
	Title       string `json:"title"`
	Teacher_id  string `json:"teacher_id"`
	Description string `json:"description"`
}
