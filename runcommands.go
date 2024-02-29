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

func main() {
	for {
		var keuze int
		var vins []string
		fmt.Println("\n0. AANPASSEN X VAN VIN")
		fmt.Println("1. HH MODEL naar TOS MODEL")
		fmt.Println("2  FINISH AFO")
		fmt.Println("3. VERANDER TIJDSTIP vins left")
		fmt.Println("4  ARRIVALSCAN VAN VINS")	
		fmt.Println("5  POSITIONSCAN POSITIE VAN VINS")		
		fmt.Println("6. Zoekers VIN blad afdrukken")
		fmt.Println("7. Gele cargo sticker afdrukken")
		fmt.Println("8. Block position")
		fmt.Println("9. Wijzig positie meerdere vins")
		fmt.Println("10. Inventory Check")
		fmt.Println("11. Vins opladen onder laatste 7 of 8")
		fmt.Println("12. CATEGORY 2 aanpassen")
		fmt.Println("13. ReadyForTransport/ReadyForPickup")
		fmt.Println("14. Next Pos aanpassen")
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)
		if keuze == 0 {
			var item string
			var waarde string
			fmt.Print("Item dat je wil aanpassen (bv weight): ")
			fmt.Scanln(&item)
			vins = plakInput("vins")
			fmt.Print("Waarde: ")
			fmt.Scanln(&waarde)
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
			var ordernr []string
			ordernr = plakInput("ordernrs")
			for i := range ordernr {
				fmt.Print("finishafo, afo ", ordernr[i], "\n")
			}
		} else if keuze == 3 {
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
		} else if keuze == 4 {
			var visit string
			fmt.Print("Visit: ")
			fmt.Scanln(&visit)
			visit = strings.ToUpper(visit)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("ArrivedScan, vin ", vins[i], ", visit ", visit, "\n")
			}
		} else if keuze == 5 {
			var position string
			var slot int
			fmt.Print("Scan op positie: ")
			fmt.Scanln(&position)
			position = strings.ToUpper(position)
			fmt.Print("Slot: ")
			fmt.Scanln(&slot)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("PositionScan, vin ", vins[i], ", position ", position, " ,slot ", slot, "\n")
			}
		} else if keuze == 6 {
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
		} else if keuze == 7 {
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
		} else if keuze == 8 {
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
		} else if keuze == 9 {
			var position string
			var startS int
			var vis bool
			fmt.Print("Scan op positie: ")
			fmt.Scanln(&position)
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
		} else if keuze == 10 {			
			var check string
			fmt.Print("Inventory Check (TRUE/FALSE): ")
			fmt.Scanln(&check)
			check = strings.ToUpper(check)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", Inventory_Check ", check, "\n")
			}
		} else if keuze == 11 {
			var laatste int
			fmt.Print("Onder laatste: ")
			fmt.Scanln(&laatste)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", additional_id ", vins[i][len(vins[i])-laatste:], "\n")
			}
		} else if keuze == 12 {
			var waarde string
			fmt.Print("Waarde: ")
			fmt.Scanln(&waarde)
			vins = plakInput("vins")			
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", CATEGORY_2 ", waarde, "\n")
			}
		} else if keuze == 13 {
			var menu int
			var sbool, sbool2 string
			fmt.Println("1. RFT")
			fmt.Println("2. RFP")
			fmt.Println("3. Beide")
			fmt.Print("Keuze: ")
			fmt.Scanln(&menu)
			vins = plakInput("vins")
			if menu == 1 {
				fmt.Print("RFT (true / false): ")
				fmt.Scanln(&sbool)
				for i := range vins {
					fmt.Print("SetRFT ,vin ", vins[i], ", remove ", sbool, "\n")
				}
			} else if menu == 2 {
				fmt.Print("RFP (true / false): ")
				fmt.Scanln(&sbool)
				for i := range vins {
					fmt.Print("SetRFP ,vin ", vins[i], ", remove ", sbool, "\n")
				}

			} else {
				fmt.Print("RFT (true / false): ")
				fmt.Scanln(&sbool)
				fmt.Print("RFP (true / false): ")
				fmt.Scanln(&sbool2)
				for i := range vins {
					fmt.Print("SetRFT ,vin ", vins[i], ", remove ", sbool, "\n")
					fmt.Print("SetRFP ,vin ", vins[i], ", remove ", sbool2, "\n")
				}
			}
		} else if keuze == 14 {
			var nextpos string
			fmt.Print("Next Pos: ")
			fmt.Scanln(&nextpos)
			nextpos = strings.ToUpper(nextpos)
			vins = plakInput("vins")
			for i := range vins {
				fmt.Print("vin_edit ", vins[i], ", Next_Pos ", nextpos, "\n")
			}
		}
	}
}
