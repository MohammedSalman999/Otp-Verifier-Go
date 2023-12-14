package data

type OTPData struct {
	PhoneNumber string `json:"phonenumber,omitempty" validate:"reqired" `
}

type VarifyData struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string   `json:"code,omitempty" validate:"required"`
}

