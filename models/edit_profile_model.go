package models

type EditProfile struct {
	UserName     string `json:"user_name" bson:"user_name" validate:"required,max=30"`
	ProfileImage string `json:"profile_image" bson:"profile_image" validate:"required,uri"`
}
