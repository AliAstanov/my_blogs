package models

import "github.com/google/uuid"

type Viewer struct {
	ViewerID uuid.UUID `json:"viewr_id"`
	Fullname string    `json:"fullname"`
	Username string    `json:"username"`
	Gmail    string    `json:"gmail"`
	Password string    `json:"password"`
}

type CheckViewer struct {
	Gmail string `json:"gmail"`
}

type ViewerRegReq struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Gmail    string `json:"gmail"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
}
