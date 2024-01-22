package helloworld

import "fmt"

const EnglishHelloPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		return EnglishHelloPrefix + name
	}
	return languagePrefix(language) + name
}
func languagePrefix(language string) (prefix string) {
	switch language {
		case "Spanish", "spanish", "espanol", "Espanol":
			prefix = "Hola, "
		case "French", "french", "francais", "Francais":
			prefix = "Bonjour, "
		case "Nepali", "nepali", "Nepal", "nepal":
			prefix = "Namaste, "
		default:
			prefix = "Hello, "
	}
	return
}
func main() {
	fmt.Println(Hello("Kumar Krishna Tripathi", "English"))
}