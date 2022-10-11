package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var BlokIndex = []string{"E", "F", "H", "KCA", "KD", "KCB", "PL", "J", "L", "N", "R", "Q2", "K408", "K409", "LOODSC", "HLN"}
var Lijnplaatsen = [][]int{
	/*E*/ {19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 0, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 15, 9, 2},
	/*F*/ {20, 21, 22, 22, 21, 21, 21, 21, 21, 21, 19, 19, 18, 16, 17, 22, 22, 18, 18, 18},
	/*H*/ {26, 26, 26, 26, 26, 25, 24, 25, 25, 25, 25, 25, 25, 25, 22, 22, 22, 22, 22, 22},
	/*KCA*/ {20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 14, 14, 14, 14, 14, 14},
	/*KD*/ {29, 29, 29, 29, 29},
	/*KCB*/ {20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 14, 14, 14, 14, 14, 14},
	/*PL*/ {21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
	/*J*/ {27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
	/*L*/ {30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 28, 28, 28},
	/*N*/ {25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
	/*R*/ {21, 24, 26, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
	/*Q2*/ {20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
	/*K408*/ {39, 39, 39, 39, 39, 39},
	/*K409*/ {49, 49, 49, 49, 49, 49},
	/*LOODSC*/ {10, 10, 10, 10, 10, 10, 10, 10, 0, 6, 12, 12, 12, 12, 12, 12, 12, 12, 12, 10, 10, 10},
	/*HLN*/ {15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
}

// lengtecat --> veryshort/short/normal/long/verylong
type Auto struct {
	merk      string
	model     string
	lengtecat int
	hoogte    float64
}

func geefPlaatsenOpBlok(blok string, startlijn int, eindlijn int, eigenlengte float64) [4]int {
	var blokindex int
	var totaalPlaatsen [4]int
	var vermenigvuldiger = [4]float64{1, 1.114, 0.86, 1.814}
	for s := range BlokIndex {
		if BlokIndex[s] == strings.ToUpper(blok) {
			blokindex = s
			break
		}
	}
	for i := startlijn; i <= eindlijn; i++ {
		if eigenlengte == 0 {
			for v := range vermenigvuldiger {
				var plaatsen float64 = vermenigvuldiger[v] * float64(Lijnplaatsen[blokindex][i-1])
				totaalPlaatsen[v] += int(math.Round(plaatsen))
			}
		} else {
			var plaatsen float64 = (4.9 / eigenlengte) * float64(Lijnplaatsen[blokindex][i-1])
			totaalPlaatsen[0] += int(math.Round(plaatsen))
		}
	}
	return totaalPlaatsen
}

func main() {
	for {
		var keuze int
		fmt.Println("\n1. Kies kaai")
		fmt.Println("2. Aantal plaats voor blok")
		fmt.Println("3. Overzicht aantal lijnen per blok")
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)

		if keuze == 1 {

		} else if keuze == 2 {
			var start string
			var blok string
			var eigenlengte float64
			var tot [4]int
			fmt.Print("Eigen lengte (0 is geen eigen lengte): ")
			fmt.Scanln(&eigenlengte)
			fmt.Print("Blok: ")
			fmt.Scanln(&blok)
			for start != "." {
				var eindln int
				fmt.Print("Start lijn: ")
				fmt.Scanln(&start)
				startln, _ := strconv.Atoi(start)
				if start != "." {
					fmt.Print("Stop lijn: ")
					fmt.Scanln(&eindln)
					var totL [4]int
					if eigenlengte > 0 {
						totL = geefPlaatsenOpBlok(blok, startln, eindln, eigenlengte)
					} else {
						totL = geefPlaatsenOpBlok(blok, startln, eindln, 0)
					}
					for leng := 0; leng < 4; leng++ {
						tot[leng] += totL[leng]
					}
				}
			}
			if eigenlengte > 0 {
				fmt.Println("\nOp blok ", strings.ToUpper(blok), ":")
				el := fmt.Sprintf("%.1f", eigenlengte)
				fmt.Println("voor auto's (", el, "m ) zijn er ", tot[0], " plaatsen")
			} else {
				fmt.Println("\nOp blok ", strings.ToUpper(blok), ":")
				fmt.Println("voor MG_5 auto's (4.9m) zijn er ", tot[0], " plaatsen")
				fmt.Println("voor PEUGEOT_208 (4.4m) auto's zijn er ", tot[1], " plaatsen")
				fmt.Println("voor FORD_RANGER (5.7m) auto's zijn er ", tot[2], " plaatsen")
				fmt.Println("voor CITROEN_AMI (2.7m) auto's zijn er ", tot[3], " plaatsen")
			}
		} else if keuze == 3 {
			for i := range BlokIndex {
				fmt.Println(BlokIndex[i], ":  ", len(Lijnplaatsen[i]), " lijnen")
			}
		}
	}
}
