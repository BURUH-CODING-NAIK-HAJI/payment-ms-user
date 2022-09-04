package requestentity

import validation "github.com/go-ozzo/ozzo-validation/v4"

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Validate() error {
	payload := *user
	return validation.ValidateStruct(
		&payload,
		validation.Field(&payload.Name, validation.Required),
		validation.Field(&payload.Username, validation.Required),
		validation.Field(&payload.Password, validation.Required, validation.Length(8, 32)),
	)
}
