package main

import "fmt"

func Texte(mot, mot1 string) string {
	return mot + " " + mot1
}

func main() {

	phrase := Texte("yacine", "norhane")
	fmt.Println(phrase)

}
