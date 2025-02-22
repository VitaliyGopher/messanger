package postgres

import "github.com/VitaliyGopher/messanger/internal/pkg/model"

type VerifyCodeRepo struct {
	store *Storage
}

func NewVerifyCodeRepo(s *Storage) *VerifyCodeRepo {
	return &VerifyCodeRepo{
		store: s,
	}
}

func (r *VerifyCodeRepo) CreateCode(sms *model.VerifyCode) {
	r.store.DB.QueryRow(
		"INSERT INTO verification_codes (email, code, time_expire) VALUES ($1, $2, $3);",
		sms.Email,
		sms.Code,
		sms.Timestamp,
	)
}

func (r *VerifyCodeRepo) FindVerifyCode(VerifyCode *model.VerifyCode) (*model.VerifyCode, error) {
	newCode := &model.VerifyCode{}
	
	if err := r.store.DB.QueryRow(
		"SELECT email, code, time_expire FROM verification_codes WHERE email = $1;",
		VerifyCode.Email,
	).Scan(
		&newCode.Email,
		&newCode.Code,
		&newCode.Timestamp,
	); err != nil {
		return VerifyCode, err
	}

	return newCode, nil
}

func (r *VerifyCodeRepo) DeleteCode(VerifyCode *model.VerifyCode) {
	r.store.DB.QueryRow(
		"DELETE FROM verification_codes WHERE email = $1",
		VerifyCode.Email,
	)
}
