package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         int(user.ID),
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token, //token tidak disimpan didatabase
	}

	return formatter
}
