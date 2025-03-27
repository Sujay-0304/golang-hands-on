package testable_program


func Greet(name string) string {
	return "Hello, " + name
}

func SendGreeting(name string) string {
	return Greet(name) + "! Welcome!"
}