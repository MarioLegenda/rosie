package models

type Sample struct {
	Name     string `faker:"name"`
	LastName string `faker:"lastName"`
	Email    string `faker:"email"`

	Min int
	Max int
}
