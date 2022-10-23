package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func announceCargo(vin string, model string, bestemming string, custommer string) {
	fmt.Print("AnnounceCargo,allowReannounce true,vin", vin, ",cargotype VEHICLE,port_of_loading BEZEE, customer", custommer, "port_of_destination", bestemming, ",model", model, "\n")
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
		fmt.Println("2 CRO --> WWL")
		fmt.Println("3. CRO --> NIT")
		fmt.Println("9. OWN SHUNT")
		fmt.Println("")
		fmt.Println("")
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)
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
				vins = append(vins, vin)
			} else {
				break
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
		} else if keuze == 9 {
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
		}
	}
}
