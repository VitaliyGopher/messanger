package model

type User struct {
	ID          uint
	Username    string
	PhoneNumber string
}

type Sms struct {
	Phone     string
	Code      int
	Timestamp int64
}
