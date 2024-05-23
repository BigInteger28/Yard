package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func getUserInput(text string) string {
	var input string
	fmt.Print(text)
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	input = scanner.Text()
	//Nu moeten we de nieuwe lijn verwijderen uit de string
	input = strings.TrimSuffix(input, "\n")
	return input
}

func main() {
	for {
		var keuze int
		var vins []string
		fmt.Println("\n0. AANPASSEN X VAN VIN")
		fmt.Println("1. HH MODEL naar TOS MODEL")
		fmt.Println("2. AFO edit End_Position")
		fmt.Println("3  FINISH AFO")
		fmt.Println("4. VERANDER TIJDSTIP vins left")
		fmt.Println("5  ARRIVALSCAN VAN VINS")	
		fmt.Println("6  POSITIONSCAN POSITIE VAN VINS")		
		fmt.Println("7. Zoekers VIN blad afdrukken")
		fmt.Println("8. Gele cargo sticker afdrukken")
		fmt.Println("9. Block position")
		fmt.Println("10. Wijzig positie meerdere vins")
		fmt.Println("11. Inventory Check")
		fmt.Println("12. Vins opladen onder laatste 7 of 8")
		fmt.Println("13. CATEGORY 2 aanpassen")
		fmt.Println("14. ReadyForTransport/ReadyForPickup")
		fmt.Println("15. Next Pos aanpassen")
		fmt.Println("16. Position aanpassen vins (inclusief CAT1/CAT10)")
		fmt.Println("17. Extra 1 aanpasen")
		fmt.Println("18. Vins op Canadakaai zetten (stellantis)")
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)
		if keuze == 0 {
			item := getUserInput("Item dat je wil aanpassen (bv weight): ")
			vins = plakInput("vins")
			waarde := getUserInput("Waarde: ")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", ", item, " ", waarde, "\n")
			}
		} else if keuze == 1 {
			fmt.Println("Als het woodrings zijn, zet een W achter de input")
			var inputmodel []string = []string {
				"A60H",	"A45GFS", "L220H", "A40G", "A40GW", "A45G", "A45GFS", "A35G", "L150H", "L180H",
				"L220HW", "L150HW", "L180HW", "L90H", "L350H", "L350HW", "L350H2", "FH380", "FH400", "FH420",
				"LC450H", "LC450HW", "", "", "", "", "", "", "", "",
				"L260H", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "",
			}
			var outputmodel []string = []string {
				"VA60H", "VA45G", "VL220H", "VA40G", "VA40GW", "VA45G", "VA45G", "VA35G", "VL150H", "VL180H",
				"VL220HW", "VL150HW", "VL180HW", "VL90H", "VL350H", "VL350HW", "VL350H", "VFH380", "VFH400", "VFH420",
				"VLC450H", "VLC450HW", "", "", "", "", "", "", "", "",
				"VL260H", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "",
			}
			var modellen []string = plakInput("modellen")
			for s := range modellen {
				modellen[s] = strings.ToUpper(modellen[s])
			}
			for m := range modellen {
				for i := range inputmodel {
					if inputmodel[i] == modellen[m] {
						fmt.Println(outputmodel[i])
						break
					}
				}
			}
		} else if keuze == 2 {
			afos := plakInput("AFO's: ")
			endPosition := getUserInput("End_Position: ")
			for i := range afos {
				fmt.Print("AFO_EDIT ", afos[i], ", END_POSITION ", endPosition, "\n")
			}
		} else if keuze == 3 {
			var ordernr []string
			ordernr = plakInput("ordernrs")
			for i := range ordernr {
				fmt.Print("finishafo, afo ", ordernr[i], "\n")
			}
		} else if keuze == 4 {
			dtsleft := getUserInput("Datum left (dd.mm.yyyy): ")
			hrleft := getUserInput("Tijdstip left (hh:mm): ")
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", dts_left ", dtsleft, " ", hrleft, "\n")
			}
		} else if keuze == 5 {
			visit := getUserInput("Visit: ")
			visit = strings.ToUpper(visit)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("ArrivedScan, vin ", vins[i], ", visit ", visit, "\n")
			}
		} else if keuze == 6 {
			var slot int
			position := getUserInput("Scan op positie: ")
			position = strings.ToUpper(position)
			fmt.Print("Slot: ")
			fmt.Scanln(&slot)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("PositionScan, vin ", vins[i], ", position ", position, " ,slot ", slot, "\n")
			}
		} else if keuze == 7 {
			var keuze int
			var printer string
			fmt.Println("1. PKY406\n2.HAN07")
			fmt.Print("Keuze: ")
			fmt.Scanln(&keuze)
			if keuze == 1 {
				printer = "PKY406"
			} else {
				printer = "HAN07"
			}
			var vins []string
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("PrintNextPosStickerexecution,vin ", vins[i], ", printerName ", printer,"\n")
			}
		} else if keuze == 8 {
			var keuze int
			var printer string
			fmt.Println("1. PTO425\n2.PTOHAN08")
			fmt.Print("Keuze: ")
			fmt.Scanln(&keuze)
			if keuze == 1 {
				printer = "PTO425"
			} else {
				printer = "PTOHAN08"
			}
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("PrintCargoLabelexecution,vin ", vins[i], ", printerName ", printer,"\n")
			}
		} else if keuze == 9 {
			var positions []string
			var starttijd string
			var eindtijd string
			scanner := bufio.NewScanner(os.Stdin)			
			fmt.Print("Start tijd (dd.mm.yyyy hh:mm): ")
			scanner.Scan()
			starttijd = scanner.Text()
			scanner = bufio.NewScanner(os.Stdin)
			fmt.Print("Eind tijd (dd.mm.yyyy hh:mm): ")
			scanner.Scan()
			eindtijd = scanner.Text()
			positions = plakInput("positions")
			for i := range positions {
				fmt.Print("BlockPosition, position", positions[i], ",start_time", starttijd, ",end_time", eindtijd,"\n")
			}
		} else if keuze == 10 {
			var startS int
			var vis bool
			position := getUserInput("Scan op positie: ")
			position = strings.ToUpper(position)
			fmt.Print("Start slot (-1 voor visgraat): ")
			fmt.Scanln(&startS)
			if startS == - 1 {
				vis = true
			}
			vins = plakInput("vins")
			if vis {
				for i := range vins {
					fmt.Print("vin_edit ", vins[i], ", Position ", position, ", SLOT ", 1, "\n")
				}
			} else {
				for i := range vins {
					fmt.Print("vin_edit ", vins[i], ", Position ", position, ", SLOT ", startS+i, "\n")
				}
			}			
		} else if keuze == 11 {
			check := getUserInput("Inventory Check (TRUE/FALSE): ")
			check = strings.ToUpper(check)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", Inventory_Check ", check, "\n")
			}
		} else if keuze == 12 {
			var laatste int
			fmt.Print("Onder laatste: ")
			fmt.Scanln(&laatste)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", additional_id ", vins[i][len(vins[i])-laatste:], "\n")
			}
		} else if keuze == 13 {
			waarde := getUserInput("Waarde: ")
			vins = plakInput("vins")			
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", CATEGORY_2 ", waarde, "\n")
			}
		} else if keuze == 14 {
			var menu int
			var sbool, sbool2 string
			fmt.Println("1. RFT")
			fmt.Println("2. RFP")
			fmt.Println("3. Beide")
			fmt.Print("Keuze: ")
			fmt.Scanln(&menu)
			vins = plakInput("vins")
			if menu == 1 {
				sbool = getUserInput("RFT (true / false): ")
				for i := range vins {
					fmt.Print("SetRFT ,vin ", vins[i], ", remove ", sbool, "\n")
				}
			} else if menu == 2 {
				sbool = getUserInput("RFP (true / false): ")
				for i := range vins {
					fmt.Print("SetRFP ,vin ", vins[i], ", remove ", sbool, "\n")
				}

			} else {
				sbool = getUserInput("RFT (true / false): ")
				sbool2 = getUserInput("RFP (true / false): ")
				for i := range vins {
					fmt.Print("SetRFT ,vin ", vins[i], ", remove ", sbool, "\n")
					fmt.Print("SetRFP ,vin ", vins[i], ", remove ", sbool2, "\n")
				}
			}
		} else if keuze == 15 {
			nextpos := getUserInput("Next Pos: ")
			nextpos = strings.ToUpper(nextpos)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", Next_Pos ", nextpos, "\n")
			}
		} else if keuze == 16 {
			position := getUserInput("Position: ")
			position = strings.ToUpper(position)
			categorie1 := getUserInput("Categorie 1 (BAT/NIT/AERTS/CHZ/CSPICO/SHUNTING/ZWKICO/OSTICO): ")
			categorie1 = strings.ToUpper(categorie1)
			categorie10 := getUserInput("Categorie 10 (BATBAT/BATHTZ/CHZ_SUB/CSPICOYARD/NIT1/NIT2/ZWKICO/ZWKICOYARD/OSTOSTICO): ")
			categorie10 = strings.ToUpper(categorie10)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", CATEGORIE_1 ", categorie1, "\n")
				fmt.Print("vin_edit ", vins[i], ", CATEGORIE_10 ", categorie10, "\n")
			}
		} else if keuze == 17 {
			vins = plakInput("vins")
			extra1 := getUserInput("EXTRA 1: ")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", EXTRA_1 \"", extra1, "\"\n")
			}
		} else if keuze == 18 {
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", LOCATION CANA", "\n")
				fmt.Print("vin_edit ", vins[i], ", CATEGORY_1 CANA", "\n")
			}
		}
	}
}
