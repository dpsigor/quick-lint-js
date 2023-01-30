package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func main() {
	p := message.NewPrinter(language.English)

	p.Printf("%v of gophers lose too much fur\n", number.Percent(0.12))
}
