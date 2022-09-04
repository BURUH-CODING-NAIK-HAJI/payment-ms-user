package responseentity

import "time"

type Profile struct {
	Id        string    `json:"id"`
	UserId    string    `json:"userId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
