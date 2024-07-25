package main

import (
	"fmt"
	//"hello/iteration"
)

const (
	spanish = "Spanish"
	french  = "French"

	frenchHelloPrefix  = "Bonjour, "
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {

	/*
		if language == spanish {
			return spanishHelloPrefix + name
		}
		if language == french {
			return frenchHelloPrefix + name
		}
		return englishHelloPrefix + name
	*/

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
		break
	case french:
		prefix = frenchHelloPrefix
		break
	default:
		prefix = englishHelloPrefix
		break
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", ""))

	//iteration.ForLoops()

}
