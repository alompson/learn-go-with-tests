package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const worldFrench = "Monde"
const worldSpanish = "Mundo"
const worldEnglish = "World"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
func Hello(name string, language string) string {
	if name == ""{
		name = worldLanguage(language)
	}
	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string){
	switch language{
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func worldLanguage(language string) (world string){
	switch language{
	case spanish:
		world = worldSpanish
	case french:
		world = worldFrench
	default:
		world = worldEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", ""))
}