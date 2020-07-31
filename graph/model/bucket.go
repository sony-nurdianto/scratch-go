package model

//UserBucket Model
type UserBucket struct {
	ID       string     `json:"id"`
	Products []*Product `json:"products"`
	UserName *Users     `json:"user_name"`
}
