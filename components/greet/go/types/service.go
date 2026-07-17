package greet_types

type GreetService interface {
	Greet(name string) string
}
