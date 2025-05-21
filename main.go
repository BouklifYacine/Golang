package main

import "fmt"

type adresse struct {
	rue    string
	numero int
}

func main() {

	commande := adresse{
		rue:    "17 rue victor hugo",
		numero: 17,
	}

	fmt.Println(commande.numero)

}
