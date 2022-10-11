package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var BlokIndex = []string{"E", "KCA", "KCB", "PL", "J", "L", "N", "R", "Q2", "K408", "K409", "LOODSC", "HLN"}
var Lijnplaatsen = [][]int{
	/*E*/ {19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 0, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 15, 9, 2},
	/*KCA*/ {20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 14, 14, 14, 14, 14, 14},
	/*KCB*/ {20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 14, 14, 14, 14, 14, 14},
	/*PL*/ {21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
	/*J*/ {27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
	/*L*/ {30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 28, 28, 28},
	/*N*/ {25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
	/*R*/ {21, 24, 26, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
	/*Q2*/ {20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
	/*K408*/ {39, 39, 39, 39, 39, 39},
	/*K409*/ {49, 49, 49, 49, 49, 49},
	/*LOODSC*/ {240},
	/*HLN*/ {},
}

// lengtecat --> veryshort/short/normal/long/verylong
type Auto struct {
	merk      string
	model     string
	lengtecat int
	hoogte    float64
}

func geefPlaatsenOpBlok(blok string, startlijn int, eindlijn int) [4]int {
	var blokindex int
	var totaalPlaatsen [4]int
	var vermenigvuldiger = [4]float64{1, 1.136, 0.877, 1.85}
	/*
		Normaal 	5m/slot
		Korte		4.4m/slot
		Lange		5.7m/slot
		Heel korte	2.7m/slot
	*/
	for s := range BlokIndex {
		if BlokIndex[s] == strings.ToUpper(blok) {
			blokindex = s
			break
		}
	}
	for i := startlijn; i <= eindlijn; i++ {
		for v := 0; v < 4; v++ {
			var plaatsen float64 = vermenigvuldiger[v] * float64(Lijnplaatsen[blokindex][i-1])
			totaalPlaatsen[v] += int(math.Round(plaatsen))
		}
	}
	return totaalPlaatsen
}

func main() {
	for {
		var keuze int
		fmt.Println("1. Kies kaai")
		fmt.Println("2. Aantal plaats voor blok")
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)

		if keuze == 1 {

		} else if keuze == 2 {
			var start string
			var blok string
			var tot [4]int
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
					totL := geefPlaatsenOpBlok(blok, startln, eindln)
					for leng := 0; leng < 4; leng++ {
						tot[leng] += totL[leng]
					}
				}
			}
			fmt.Println("Op blok ", strings.ToUpper(blok), ":")
			fmt.Println("voor VOLVO_XC60 / MG_5 auto's zijn er ", tot[0], " plaatsen")
			fmt.Println("voor PEUGEOT_208 auto's zijn er ", tot[1], " plaatsen")
			fmt.Println("voor FORD_RANGER auto's zijn er ", tot[2], " plaatsen")
			fmt.Println("voor CITROEN_AMI auto's zijn er ", tot[3], " plaatsen")
		}
	}
}
