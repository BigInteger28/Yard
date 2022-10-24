package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func announceCargo(vin string, model string, bestemming string, custommer string) {
	fmt.Print("AnnounceCargo,allowReannounce true,vin", vin, ",cargotype VEHICLE,port_of_loading BEZEE, customer", custommer, ",port_of_destination", bestemming, ",model", model, "\n")
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
	fmt.Print("..., RMARX")
	fmt.Print("\nModel: ")
	fmt.Scanln(&model)
	return strings.ToUpper(model)
}

func main() {
	for {
		var keuze int
		fmt.Println("1. Polestar WWL --> NIT")
		fmt.Println("2  WWL --> NIT")
		fmt.Println("3. CRO --> NIT")
		fmt.Println("4. OWN SHUNT")
		fmt.Println("5. Edit shuntopdracht naam en ordertype")
		fmt.Println("")
		fmt.Println("")
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)
		var bestemming string
		var vins []string		
		if keuze < 4  {
			fmt.Print("\nBestemming: ")
			fmt.Scanln(&bestemming)
			bestemming = strings.ToUpper(bestemming)
			fmt.Print("Plak vins in: ")
			scanner := bufio.NewScanner(os.Stdin)
			for {
				scanner.Scan()
				vin := scanner.Text()
				if len(vin) != 0 {
					vins = append(vins, vin)
				} else {
					break
				}
			}
		}		
		if keuze == 1 {
			for i := range vins {
				announceCargo(vins[i], "POLE2", bestemming, "POLESTAR")
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_CodeSHUNT_WWL_NIT,executorWWL,customerICO\n")
			}
		} else if keuze == 2 {
			model := askModel()
			for i := range vins {
				announceCargo(vins[i], model, bestemming, "ICO")
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_CodeSHUNT_WWL_NIT,executorWWL,customerICO\n")
			}
		} else if keuze == 3 {
			model := askModel()
			for i := range vins {
				announceCargo(vins[i], model, bestemming, "ICO")
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_CodeSHUNT_CANADAKAAI_NIT,executorCRO,customerICO\n")
			}
		} else if keuze == 4 {
			var ordercode string
			model := askModel()
			fmt.Println("Order Code -->")
			fmt.Println("WWL_NIT")
			fmt.Println("CHZ_NIT")
			fmt.Println("...")
			fmt.Print("\n Order code: ")
			fmt.Scanln(&ordercode)
			ordercode = strings.ToUpper(ordercode)
			for i := range vins {
				announceCargo(vins[i], model, bestemming, "ICO")
				fmt.Print("AssignCustomerOrder,vin", vins[i], ",order_CodeSHUNT_", ordercode, ",executorCRO,customerICO\n")
			}
		} else if keuze == 5 {
			editShuntOrder()
		}
	}
}
