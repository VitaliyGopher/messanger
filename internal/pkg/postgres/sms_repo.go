package postgres

import "github.com/VitaliyGopher/messanger/internal/pkg/model"

type SmsRepo struct {
	store *Storage
}

func NewSmsRepo(s *Storage) *SmsRepo {
	return &SmsRepo{
		store: s,
	}
}

func (r *SmsRepo) CreateSmsCode(sms *model.Sms) {
	r.store.DB.QueryRow(
		"INSERT INTO sms_codes (phone, code, time_expire) VALUES ($1, $2, $3);",
		sms.Phone,
		sms.Code,
		sms.Timestamp,
	)
}

func (r *SmsRepo) FindSmsCode(sms *model.Sms) (*model.Sms, error) {
	newSms := &model.Sms{}
	
	if err := r.store.DB.QueryRow(
		"SELECT phone, code, time_expire FROM sms_codes WHERE phone = $1;",
		sms.Phone,
	).Scan(
		&newSms.Phone,
		&newSms.Code,
		&newSms.Timestamp,
	); err != nil {
		return sms, err
	}

	return newSms, nil
}

func (r *SmsRepo) DeleteSmsCode(sms *model.Sms) {
	r.store.DB.QueryRow(
		"DELETE FROM sms_codes WHERE phone = $1",
		sms.Phone,
	)
}
