package model

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type VerifyCode struct {
	Email     string `json:"email"`
	Code      int    `json:"verification_code"`
	Timestamp int64  `json:"timestamp"`
}
