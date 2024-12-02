package models

type Category struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image []byte `json:"image"`
}
