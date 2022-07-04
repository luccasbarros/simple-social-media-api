package models

// Password represents the request format to update an user's password
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
