package account

type Register struct {
	FullName     string `json:"fullname"`
	PasswordText string `json:"password"`
	Email        string `json:"email"`
}

type VerifyCode struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type EmailTemplate struct {
	To   string `json:"to"`
	Link string `json:"link"`
	Type string `json:"type"`
}

const (
	Verification  = "verification"
	PasswordReset = "passwordReset"
)
