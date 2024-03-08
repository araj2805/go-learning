package main

import "fmt"

const (
	englishHello = "Hello "
	spanishHello = "Hola "
	frenchHello  = "Bonjor "

	spanish = "Spanish"
	french  = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	/*if language == spanish {
		return spanishHello + name
	} else if language == french {
		return frenchHello + name
	}*/

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHello
	case spanish:
		prefix = spanishHello
	default:
		prefix = englishHello
	}

	return

}

func main() {
	fmt.Println(Hello("Ankit", "Spanish"))
}
