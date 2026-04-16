package main

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	russianHelloPrefix = "Привет, "

	french  = "French"
	spanish = "Spanish"
	russian = "Russian"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	default:
		prefix = "Hello, "
	}
	return
}
func main() {

}
