package model

type UserGenderEnum int32

const (
	USER_GENDER_UNSPECIFIED UserGenderEnum = iota
	USER_GENDER_MALE
	USER_GENDER_FEMALE
)

type User struct {
	Id       string         `json:"id"`
	Username string         `json:"username"`
	Password string         `json:"password,omitempty"`
	Name     string         `json:"name"`
	Dob      string         `json:"dob"`
	Gender   UserGenderEnum `json:"gender"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at,omitempty"`
}
