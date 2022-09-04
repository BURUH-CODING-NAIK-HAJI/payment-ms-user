package responseentity

import "time"

type User struct {
	Id        string      `json:"id"`
	Username  string      `json:"username"`
	Password  interface{} `json:"password,omitempty"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Profile   Profile     `json:"profile,omitempty"`
}
