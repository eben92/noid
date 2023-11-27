package database

type SignupPayload struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Msisdn   string `json:"msisdn"`
}

type SigninPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(fullName, password, email, msisdn string) (*User, error) {

	// encrypt password

	return &User{
		FullName: fullName,
		Password: password,
		Email:    email,
		Msisdn:   msisdn,
	}, nil
}
