package model

type OtpRequest struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}
