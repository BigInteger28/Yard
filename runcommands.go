package main

import (
	"bufio"
	"fmt"
	"os"
)

func plakInput(inputTxt string) []string {
	var list []string
	fmt.Print("Plak ", inputTxt, " in: ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		item := scanner.Text()
		if len(item) != 0 {
			list = append(list, item)
		} else {
			break
		}
	}
	return list
}

func main() {
	for {
		var keuze int
		fmt.Println("1. HH MODEL naar TOS MODEL")
		fmt.Println("2  FINISH AFO met ordernr")
		fmt.Println("3. VERANDER TIJDSTIP vins left")
		fmt.Println("4  ")		
		fmt.Println("5. ")
		fmt.Println("6. ")
		fmt.Println("7. ")
		fmt.Println("8. ")
		fmt.Println("9. ")
		fmt.Println("")
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)
		if keuze == 1 {
			fmt.Println("Als het woodrings zijn, zet een W achter de input")
			var inputmodel []string = []string {
				"A60H",	"A45GFS", "L220H", "A40G", "A40GW", "A45G", "A45GFS", "A35G", "L150H", "L180H",
				"L220HW", "L150HW", "L180HW", "L90H", "L350H", "L350HW", "L350H2", "", "", "",
				"", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "",
			}
			var outputmodel []string = []string {
				"VA60H", "VA45G", "VL220H", "VA40G", "VA40GW", "VA45G", "VA45G", "VA35G", "VL150H", "VL180H",
				"VL220HW", "VL150HW", "VL180HW", "VL90H", "VL350H", "VL350HW", "VL350H", "", "", "",
				"", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "",
			}
			var modellen []string = plakInput("modellen")
			for m := range modellen {
				for i := range inputmodel {
					if inputmodel[i] == modellen[m] {
						fmt.Println(outputmodel[i])
						break
					}
				}
			}
		} else if keuze == 2 {
			var ordernr []string
			ordernr = plakInput("ordernrs")
			for i := range ordernr {
				fmt.Print("finishafo, afo ", ordernr[i], "\n")
			}
		} else if keuze == 3 {
			var vins []string
			var dtsleft string
			var hrleft string
			fmt.Print("Datum left (dd.mm.yyyy): ")
			fmt.Scanln(&dtsleft)
			fmt.Print("Tijdstip left (hh:mm): ")
			fmt.Scanln(&hrleft)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", dts_left ", dtsleft, " ", hrleft, "\n")
			}
		}
	}
}