package models

type TeacherReq struct {
	Teacher_name    string `json:"teacher_name"`
	Teacher_surname string `json:"teacher_surname"`
	Teacher_number  string `json:"teacher_number"`
	Teachers_tg     string `json:"teacher_tg"`
	Teacher_bio     string `json:"bio"`
	Rating          string `json:"rating"`
}

type Teacher struct {
	Teacher_id      string `json:"teacher_id"`
	Teacher_name    string `json:"teacher_name"`
	Teacher_surname string `json:"teacher_surname"`
	Teacher_number  string `json:"teacher_number"`
	Teachers_tg     string `json:"teacher_tg"`
	Teacher_bio     string `json:"bio"`
	Rating          string `json:"rating"`
}
