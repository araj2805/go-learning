package main

import "fmt"

const englishHello = "Hello "
const spanishHello = "Hola "
const frenchHello = "Bonjor "

const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHello + name
	} else if language == french {
		return frenchHello + name
	}
	return englishHello + name
}

func main() {
	fmt.Println(Hello("Ankit", "Spanish"))
}
