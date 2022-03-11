package main

import "fmt"

const (
	french             = "French"
	italian            = "Italian"
	spanish            = "Spanish"
	frenchHelloPrefix  = "Bonjour, "
	italianHelloPrefix = "Ciao, "
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name, language string) string {
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
	case italian:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Gopher", ""))
}
