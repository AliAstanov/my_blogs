package models

type OtpData struct {
	Otp   string `json:"otp"`
	Gmail string `json:"gmail"`
}

type CheckOtp struct {
	Gmail string `json:"gmail"`
	Otp   string `json:"otp"`
}

type CheckOtpResp struct {
	IsRight bool `json:"is_right"`
}
