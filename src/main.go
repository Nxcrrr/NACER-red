package main

import (
	"fmt"
	"projet-red-N/src/utils"
	"time"
)

var p utils.Personnage

func main() {
	utils.AnimateText("Nacer présente ... Le Projet Red !")
	fmt.Println("")
	utils.AnimateText("Chargement ...")
	time.Sleep(3 * time.Second)
	p.Menu()
}
