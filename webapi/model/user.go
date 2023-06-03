package model

type User struct {
	Id int64 `json:"id" xml:"id" form:"id"`

	Username string `json:"username" xml:"username" form:"username"`

	FirstName string `json:"firstName" xml:"firstName" form:"firstName"`

	LastName string `json:"lastName" xml:"lastName" form:"lastName"`
}
