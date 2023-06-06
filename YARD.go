package main

import (
	"fmt"
	"math"
	"strings"
)

var BlokIndex = []string{"BA", "BB", "BC", "CA", "CB", "CC", "DA", "DB", "DC", "E", "HLN", "SUI", "PL", "PL100", "F", "G", "HHN", "H", "I", "KCA", "J", "K", "L", "M", "KCB", "KD", "N", "O", "P", "Q2", "Q9", "Q10", "R", "S", "SP45", "U", "U2", "W", "Z01", "Z02", "Z03", "Z04", "K405", "K408", "K409", "HHW", "KA", "LOODS-C", "KE", "KB"}
var Lijnplaatsen = [][]int{
	{4, 43},
	{4, 43},
	{4, 43},
	{4, 79},
	{4, 80},
	{4, 80},
	{4, 83},
	{4, 80},
	{4, 81},
	//E
	{20, 86},
	{15, 40},
	//SUI
	{0, 0},
	{21, 40},
	{20, 5},
	{19, 20},
	{12, 27},
	//HHN
	{8, 8},
	{24, 19},
	{12, 19},
	{18, 20},
	{30, 20},
	{12, 20},
	{30, 20},
	{12, 20},
	{18, 20},
	//KD
	{30, 4},
	{23, 21},
	{20, 26},
	{20, 47},
	//Q2
	{20, 47},
	{20, 5},
	{20, 6},
	//R
	{29, 21},
	{8, 20},
	{95, 4},
	{9, 168},
	{13, 168},
	//W
	{2, 21},
	{2, 78},
	{2, 29},
	{2, 59},
	{2, 40},
	//K405
	{38, 8},
	{47, 6},
	{50, 6},
	{18, 11},
	{30, 12},
	{11, 18},
	{18, 9},
	{8, 6},
}

type Auto struct {
	merk      string
	model     string
	lengtecat int
	hoogte    float64
}

type Blok struct {
	id    int
	naam  string
	score int
}

func geefPlaatsenOpBlok(blok string, aantalLijnen int, vermenigvuldiger float64) [4]int {
	var blokindex int
	var totaalPlaatsen [4]int
	var percentages = [4]float64{1, 0.75, 0.5, 1.8}
	for s := range BlokIndex {
		if BlokIndex[s] == strings.ToUpper(blok) {
			blokindex = s
			break
		}
	}
	if vermenigvuldiger == 0 {
		for v := range percentages {
			var plaatsen float64 = percentages[v] * float64(Lijnplaatsen[blokindex][0]) * float64(aantalLijnen)
			totaalPlaatsen[v] += int(math.Round(plaatsen))
		}
	} else {
		var plaatsen float64 = vermenigvuldiger * float64(Lijnplaatsen[blokindex][0]) * float64(aantalLijnen)
		totaalPlaatsen[0] += int(math.Round(plaatsen))
	}
	return totaalPlaatsen
}

func rekenBlok(locatie int) {
	var BootBlokken = []string{"U", "SP45", "KA", "K409", "LOODS-C", "HHW", "K408", "Q10", "Q9", "Q2", "R", "N", "L", "J", "KCB", "KD", "KCA", "H", "F", "HHN", "K405", "PL", "HLN", "E", "SUI", "DC", "DB", "DA"}
	var afstand405 = []int{0}
	var afstand408 = []int{12, 0, 15, 16, 17, 18, 22, 19, 20, 21, 9, 11, 10, 8, 14, 13, 7, 5, 4, 6, 1, 2, 3}
	blokken := make([]Blok, len(BootBlokken))
	for i := range BootBlokken {
		blokken[i].id = i
		blokken[i].naam = BootBlokken[i]
		if locatie == 405 {
			blokken[i].score += afstand405[i]
		} else {
			blokken[i].score += afstand408[i]
		}
	}

}

func hoeveelLijnenOpBlok(aantal int) int {
	var blok string
	var blokindex int
	var lengte float64
	fmt.Print("Lengte (m): ")
	fmt.Scanln(&lengte)
	fmt.Print("Blok: ")
	fmt.Scanln(&blok)
	for s := range BlokIndex {
		if BlokIndex[s] == strings.ToUpper(blok) {
			blokindex = s
			break
		}
	}
	var vermenigvuldiger float64 = 4.9 / lengte
	if aantal%int(float64(Lijnplaatsen[blokindex][0])*vermenigvuldiger) == 0 {
		return aantal / int(float64(Lijnplaatsen[blokindex][0])*vermenigvuldiger)
	} else {
		return (aantal / int(float64(Lijnplaatsen[blokindex][0])*vermenigvuldiger)) + 1
	}
}

func main() {
	for {
		var keuze int
		fmt.Println("\n1. Kies kaai")
		fmt.Println("2. Aantal plaats voor blok")
		fmt.Println("3. Hoeveel lijnen")
		fmt.Println("9. Overzicht aantal plaatsen per blok")
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)

		if keuze == 1 {

		} else if keuze == 2 {
			var lijnen int
			var blok string
			var eigenlengte float64
			var tot [4]int
			fmt.Print("Vermenigvuldiger lengte (0 is geen eigen lengte): ")
			fmt.Scanln(&eigenlengte)
			fmt.Print("Blok: ")
			fmt.Scanln(&blok)
			fmt.Print("Aantal lijnen beschikbaar: ")
			fmt.Scanln(&lijnen)
			var totL [4]int
			if eigenlengte > 0 {
				totL = geefPlaatsenOpBlok(blok, lijnen, eigenlengte)
			} else {
				totL = geefPlaatsenOpBlok(blok, lijnen, 0)
			}
			for leng := 0; leng < 4; leng++ {
				tot[leng] += totL[leng]
			}

			if eigenlengte > 0 {
				fmt.Println("\nOp blok ", strings.ToUpper(blok), ":")
				el := fmt.Sprintf("%.1f", eigenlengte)
				fmt.Println("voor auto's (", el, "m ) zijn er ", tot[0], " plaatsen")
			} else {
				fmt.Println("\nOp blok ", strings.ToUpper(blok), ":")
				fmt.Println("voor 100% marge zijn er ", tot[0], " plaatsen")
				fmt.Println("voor 75% marge (camionet) zijn er ", tot[1], " plaatsen")
				fmt.Println("voor 50% marge (lang) zijn er ", tot[2], " plaatsen")
				fmt.Println("voor 180% marege (heel kort) zijn er ", tot[3], " plaatsen")
			}
		} else if keuze == 3 {
			var aantal int
			fmt.Print("Hoeveel auto's: ")
			fmt.Scanln(&aantal)
			fmt.Println("Aantal lijnen nodig: ", hoeveelLijnenOpBlok(aantal))
		} else if keuze == 9 {
			for i := range BlokIndex {
				fmt.Println(BlokIndex[i], ":  ", Lijnplaatsen[i][0]*Lijnplaatsen[i][1], " plaatsen")
			}
		}
	}
}
