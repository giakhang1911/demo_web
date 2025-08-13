package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ClassItem struct {
	ID   string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}

type SurveyOpinion struct {
	Title    string `bson:"title" json:"title"`         // Tên mục (vd: "Đánh giá nội dung")
	Score    int    `bson:"score" json:"score"`         // 1–7; với câu text đặt 0
	Survey   string `bson:"survey" json:"survey"`       // Nội dung câu hỏi hoặc câu trả lời text
	ParentID string `bson:"parent_id" json:"parent_id"` // Nhóm: content|teacher|result|feedback
}

type CourseSurvey struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Code          string             `bson:"code" json:"code"`                     // Mã SV
	Name          string             `bson:"student_name" json:"student_name"`     // Họ tên SV
	Class         ClassItem          `bson:"class" json:"class"`                   // Lớp
	CourseName    string             `bson:"course_name" json:"course_name"`       // Tên môn
	Instructor    string             `bson:"instructor" json:"instructor"`         // Giảng viên
	SurveyOpinion []SurveyOpinion    `bson:"survey_opinion" json:"survey_opinion"` // Câu hỏi & trả lời
}
