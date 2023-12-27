package model

type Category struct {
	ID          string `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//Courses     []*Course `json:"courses"` não precisa mais
}
