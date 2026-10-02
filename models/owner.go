package models

type Owner struct {
	Fullname    string `json:"fullname"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
	Gmail       string `json:"gmail"`
	Telegram    string `json:"telegram"`
	Github      string `json:"github"`
	LinkedIn    string `json:"linked_in"`
	Leetcode    string `json:"leetcode"`
	AboutMe     string `json:"about_me"`
}

type LoginOwn struct {
	PhoneNumber string `json:"phone_number"`
	Gmail       string `json:"gmail"`
	Password    string `json:"password"`
}