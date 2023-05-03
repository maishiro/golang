package entity

type User struct {
	Id int64 `xorm:"id"`

	Username string `xorm:"username"`

	FirstName string `xorm:"firstname"`

	LastName string `xorm:"lastname"`

	Email string `xorm:"email"`

	Password string `xorm:"password"`

	Phone string `xorm:"phone"`

	UserStatus int32 `xorm:"userstatus"`
}
