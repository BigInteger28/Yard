package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func announceCargo(vin string, model string, bestemming string, custommer string) {
	fmt.Print("AnnounceCargo,allowReannounce true,vin", vin, ",cargotype VEHICLE,port_of_loading BEZEE, customer", custommer, ",port_of_destination", bestemming, ",final_destination", bestemming, ",model", model, "\n")
}

func editShuntOrder() {
	var origineel string
	var nieuw string
	var ordercodes []string
	fmt.Print("Originele naam na YARD_SHUNT_ : ")
	fmt.Scanln(&origineel)
	origineel = strings.ToUpper(origineel)
	fmt.Print("Nieuwe naam na YARD_SHUNT_ : ")
	fmt.Scanln(&nieuw)
	nieuw = strings.ToUpper(nieuw)
	fmt.Print("Plak ordercodes in: ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		ordercode := scanner.Text()
		if len(ordercode) != 0 {
			ordercodes = append(ordercodes, ordercode)
		} else {
			break
		}
	}
	fmt.Print("\nOPERATION_EDIT YARD_SHUNT_", origineel, ",CODE YARD_SHUNT_", nieuw, "\n")
	for i := range ordercodes {
		fmt.Print("CUSTOMERORDER_EDIT ", ordercodes[i], ", ORDERCODETYPE SHUNT_", nieuw, "\n")
		fmt.Print("AFO_EDIT ", ordercodes[i], ", ORDERCODETYPE SHUNT_", nieuw, "\n")
	}
}

func askModel() string {
	var model string
	fmt.Println("Modellen --> ")
	fmt.Print("XPENG G3,\tXPG3\nXPENG P5,\tXPP5\nXPENG P7,\tXPP7\n--------------------\n")
	fmt.Print("MG3,\tMG3\nMG4 ELECTRIC,\tMMG4_EV\nMG5 ELECTRIC,\tMMG5_EV\n--------------------\n")
	fmt.Print("TOYOTA YARIS,\tTYARI\nTOYOTA YARIS CROSS,\tTYARIC\nTOYOTA SUPRA,\tTSUPR\nTOYOTA PROACE CITY,\tTPROAC\nTOYOTA PROACE CITY ELECTRIC,\tTPROACEV\nTOYOTA PROACE ELECTRIC,\tTPROAEV\nTOYOTA PROACE,\tTPROA\nTOYOTA HIGHLANDER,\tTHIGH\nTOYOTA AYGO,\tTAYGO\n--------------------\n")
	fmt.Print("\nModel: ")
	fmt.Scanln(&model)
	return strings.ToUpper(model)
}

func plakVins() (string, []string) {
	var bestemming string
	var vins []string
	fmt.Print("\nBestemming: ")
	fmt.Scanln(&bestemming)
	bestemming = strings.ToUpper(bestemming)
	fmt.Print("Plak vins in: ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		vin := scanner.Text()
		if len(vin) != 0 {
			if len(vin) > 17 {
				vin = vin[:17] // Neem alleen de eerste 17 tekens
			}
			vins = append(vins, vin)
		} else {
			break
		}
	}
	return bestemming, vins
}

func onlyVins() []string {
	var vins []string
	fmt.Print("Plak vins in: ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		vin := scanner.Text()
		if len(vin) != 0 {
			if len(vin) > 17 {
				vin = vin[:17] // Neem alleen de eerste 17 tekens
			}
			vins = append(vins, vin)
		} else {
			break
		}
	}
	return vins
}

func doPoleStar(vins []string, bestemming string, vkaaiInput string) {
	var model string
	var vkaai, nkaai, executor string
	if vkaai == "WWL" {
		executor = "WWL"
		vkaai = vkaaiInput
	} else if vkaai == "BAT" {
		executor = "CRO"
		vkaai = vkaaiInput
	} else {
		vkaai = "NIT"
	}	
	fmt.Print("Naar Kaai (BAT/HANZE/NIT=BLANCO): ")
	fmt.Scanln(&nkaai)
	nkaai = strings.ToUpper(nkaai)
	for i := range vins {
		if vins[i][5:7] == "FC" || vins[i][5:7] == "FD" || vins[i][5:7] == "FE" || vins[i][5:7] == "FE" || vins[i][5:7] == "FF" || vins[i][5:7] == "FG" {
			model = "POLE2_EV"
		} else if vins[i][5:7] == "EA" || vins[i][5:7] == "EE" || vins[i][5:7] == "EJ" {
			model = "POLE3_EV"
		} else {
			model = "POLE4_EV"
		}
		announceCargo(vins[i], model, bestemming, "ICO")
		if vkaai == "NIT" {
			fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_Code SHUNT_WWL_NIT,executorWWL,customerICO\n")
		} else {
			fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_Code SHUNT_", vkaai, "_", nkaai, ",executor", executor, ",customerICO\n")
		}		
	}
}

func main() {
	for {
		var keuze int
		fmt.Println("---------- NIT ----------")
		fmt.Println("1. Polestar WWL --> NIT (MET REANNOUCECARGO)")
		fmt.Println("2. WWL    --> NIT")
		fmt.Println("3. CRO_CANADAKAAI    --> NIT")
		fmt.Println("\n------------------------")
		fmt.Println("4. Polestar WWL/CRO --> BAT/HANZE (MET REANNOUCECARGO)")
		fmt.Println("\n---------- BAT ----------")
		fmt.Println("5. WWL    --> BAT (MET REANNOUCECARGO)")
		fmt.Println("6. WWL    --> BAT")
		fmt.Println("7. CRO_BRTT    --> BAT")
		fmt.Println("8. CLDN(PSA) --> BAT")
		fmt.Println("\n---------- HANZE ----------")
		fmt.Println("8. WWL --> HANZE")
		fmt.Println("9. HANZE    --> WWL")
		fmt.Println("10. CRO_BRTT    --> HANZE")
		fmt.Println("11. H&H eerst aanmaken CLDN(PSA) --> HANZE")
		fmt.Println("12. HANZE --> TOYOTA")
		fmt.Println("13. TOYOTA --> HANZE")
		fmt.Println("\n---------- ANDERE ----------")
		fmt.Println("14. OWN SHUNT")
		fmt.Println("15. TOYOTA --> TERMINAL")
		fmt.Println("16. Edit shuntopdracht naam en ordertype")
		fmt.Println("")
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)
		var bestemming string
		var vins []string		
		if keuze == 1 {
			bestemming, vins = plakVins()			
			doPoleStar(vins,bestemming,"")
		} else if keuze == 2 {
			bestemming, vins = plakVins()
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_Code SHUNT_WWL_NIT,executorWWL,customerICO\n")
			}
		} else if keuze == 3 {
			vins = onlyVins()
			for i := range vins {
				fmt.Println(vins[i])
			}
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_Code SHUNT_CANADAKAAI_NIT,executorCRO,customerICO\n")
			}
		} else if keuze == 4 {
			var vkaai string
			bestemming, vins = plakVins()
			fmt.Print("Van Kaai (WWL/BRITT): ")
			fmt.Scanln(&vkaai)
			vkaai = strings.ToUpper(vkaai)
			doPoleStar(vins,bestemming,vkaai)
		} else if keuze == 5 {
			bestemming, vins = plakVins()
			model := askModel()
			for i := range vins {
				announceCargo(vins[i], model, bestemming, "ICO")
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_CodeSHUNT_WWL_BAT,executorWWL,customerICO\n")
			}
		} else if keuze == 6 {
			vins = onlyVins()
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_Code SHUNT_WWL_BAT,executorWWL,customerICO\n")
			}
		} else if keuze == 7 {
			vins = onlyVins()
			for i := range vins {
				fmt.Println(vins[i])
			}
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_Code SHUNT_BRITT_BAT,executorCRO,customerICO\n")
			}
		} else if keuze == 8 {
			var keuze int
			var executor string
			fmt.Println("Executor -->")
			fmt.Println("1. MACO")
			fmt.Println("2. AUTOLUC")
			fmt.Print("Kies: ")
			fmt.Scanln(&keuze)
			if keuze == 1 {
				executor = "MACO"
			} else {
				executor = "AUTOLUC"
			}
			vins = onlyVins()
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_Code SHUNT_CLDN_BAT,executor ", executor, ",customerICO\n")
			}
		} else if keuze == 9 {
			vins = onlyVins()
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_CodeSHUNT_WWL_HANZE,executor WWL,customerICO\n")
			}
		} else if keuze == 10 {
			vins = onlyVins()
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_CodeSHUNT_HANZE_WWL,executor WWL,customerICO\n")
			}			
		} else if keuze == 11 {
			vins = onlyVins()
			for i := range vins {
				fmt.Println(vins[i])
			}
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_CodeSHUNT_BRITT_HANZE,executor CRO,customerICO\n")
			}
		} else if keuze == 12 {
			var keuze int
			var executor string
			fmt.Println("Executor -->")
			fmt.Println("1. MACO")
			fmt.Println("2. AUTOLUC")
			fmt.Print("Kies: ")
			fmt.Scanln(&keuze)
			if keuze == 1 {
				executor = "MACO"
			} else {
				executor = "AUTOLUC"
			}
			vins = onlyVins()
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_Code SHUNT_CLDN_HANZE,executor", executor, ",customerICO\n")
			}
		} else if keuze == 13 {
			vins = onlyVins()
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_Code SHUNT_HANZE_TOYO,executorICO,customerTESLA\n")
			}
		} else if keuze == 14 {
			vins = onlyVins()
			for i := range vins {
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_Code SHUNT_TOYO_HANZE,executorICO,customerTESLA\n")
			}
		} else if keuze == 15 {
			var ordercode string
			var executor string
			bestemming, vins = plakVins()
			model := askModel()
			fmt.Println("Order Code -->")
			fmt.Println("WWL_NIT")
			fmt.Println("CHZ_NIT")
			fmt.Println("...")
			fmt.Print("\n Order code: ")
			fmt.Scanln(&ordercode)
			ordercode = strings.ToUpper(ordercode)
			fmt.Print("\nExecutor: ")
			fmt.Scanln(&executor)
			executor = strings.ToUpper(executor)
			for i := range vins {
				announceCargo(vins[i], model, bestemming, "ICO")
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_CodeSHUNT_", ordercode, ",executor", executor, ",customerICO\n")
			}
		} else if keuze == 16 {
			var vins [][]string
			var terminal string
			var aantalmodellen int
			var modellen []string
			var bestemming []string
			fmt.Print("Naar welke terminal (bat/nit/htz): ")
			fmt.Scanln(&terminal)
			terminal = strings.ToUpper(terminal)
			fmt.Print("Hoeveel Toyota Modellen: ")
			fmt.Scanln(&aantalmodellen)
			modellen = make([]string, aantalmodellen)
			bestemming = make([]string, aantalmodellen)
			vins = make([][]string, aantalmodellen)
			for i := 0; i < aantalmodellen; i++ {
				modellen[i] = askModel()
				bestemming[i], vins[i] = plakVins()
			}
			for a := 0; a < aantalmodellen; a++ {
				for b := 0; b < len(vins[a]); b++ {
					announceCargo(vins[a][b], modellen[a], bestemming[a], "UECC")
					fmt.Print("AssignCustomerOrder,vin", vins[a][b], ",order_CodeSHUNT_TMME_", terminal, ",executorAUTOLUC,customerICO\n")
				}
			}
		} else if keuze == 17 {
			editShuntOrder()
		}
	}
}
