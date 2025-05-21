package main

import "fmt"

func main() {

	MaMap := make(map[int]string)

	MaMap[1] = "yacine"
	MaMap[2] = "Norhane"
	MaMap[1] = "Nadia"

	fmt.Println(MaMap[1])

	for personne := range MaMap {
		fmt.Println(personne, "est la clé de :", MaMap[personne])
	}

}
