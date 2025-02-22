package model

type User struct {
	ID          uint
	Username    string
	Email string
}

type VerifyCode struct {
	Email     string
	Code      int
	Timestamp int64
}
