package entity

type Role struct {
	Name string `json:name,validate:required`
}
