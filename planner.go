package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Auto struct {
	Bestemming string
	Aantal     int
}

type AutoGroep struct {
	Bestemming string
	Aantal     int
}

type Lijn struct {
	BlokID         string
	LijnID         string
	Capaciteit     int
	Buitenlijn     bool
	Bezet          int    // Hoeveelheid auto's momenteel op de lijn
	Voorbestemming string // Nieuw veld voor vooraf ingevulde bestemming
}

type Blok struct {
	BlokID string
	Lijnen []Lijn
}

type Toewijzing struct {
	LijnID string
	Aantal int
}

var toewijzingenPerGroep map[string][]Toewijzing

func leesAutosBestand(bestandsnaam string) ([]Auto, error) {
	file, err := os.Open(bestandsnaam)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var autos []Auto
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		aantal, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		autos = append(autos, Auto{Bestemming: parts[1], Aantal: aantal})
	}
	return autos, scanner.Err()
}

func leesBlokkenBestand(bestandsnaam string) ([]Blok, error) {
	file, err := os.Open(bestandsnaam)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var blokken []Blok
	var huidigBlok *Blok

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		// Detecteer of de regel een nieuw blokID aangeeft
		if len(parts) == 1 {
			if huidigBlok != nil {
				// Voeg het vorige blok toe aan de lijst voordat je een nieuw blok start
				blokken = append(blokken, *huidigBlok)
			}
			// Start een nieuw blok
			huidigBlok = &Blok{BlokID: parts[0], Lijnen: []Lijn{}}
			continue
		}

		// Verwerk de lijn binnen een blok
		if len(parts) >= 3 && huidigBlok != nil {
			capaciteit, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("kon capaciteit niet lezen voor lijn %s in blok %s: %v", parts[0], huidigBlok.BlokID, err)
			}

			buitenlijn, err := strconv.ParseBool(parts[2])
			if err != nil {
				return nil, fmt.Errorf("kon buitenlijnstatus niet lezen voor lijn %s in blok %s: %v", parts[0], huidigBlok.BlokID, err)
			}

			voorbestemming := ""
			if len(parts) > 3 {
				voorbestemming = parts[3]
			}

			lijn := Lijn{
				LijnID:         parts[0],
				Capaciteit:     capaciteit,
				Buitenlijn:     buitenlijn,
				Voorbestemming: voorbestemming,
			}

			huidigBlok.Lijnen = append(huidigBlok.Lijnen, lijn)
		}
	}

	// Voeg het laatste blok toe aan de lijst na het einde van de lus
	if huidigBlok != nil {
		blokken = append(blokken, *huidigBlok)
	}

	return blokken, nil
}

func groepeerEnSorteerAutos(autos []Auto) []AutoGroep {
	// Groepeer auto's per bestemming
	groepMap := make(map[string]int)
	for _, auto := range autos {
		groepMap[auto.Bestemming] += auto.Aantal
	}

	// Zet de map om naar een slice van AutoGroep voor sortering
	var groepen []AutoGroep
	for bestemming, aantal := range groepMap {
		groepen = append(groepen, AutoGroep{Bestemming: bestemming, Aantal: aantal})
	}

	// Sorteer de groepen van groot naar klein op basis van het aantal auto's
	sort.Slice(groepen, func(i, j int) bool {
		return groepen[i].Aantal > groepen[j].Aantal
	})

	return groepen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func toonIngeplandeAutos(blokken []Blok) {
	fmt.Println("Eindresultaat van de ingeplande auto's:")
	for _, blok := range blokken {
		fmt.Printf("Blok %s:\n", blok.BlokID)
		for _, lijn := range blok.Lijnen {
			if lijn.Bezet > 0 {
				buitenlijnStatus := "Nee"
				bestemmingen := lijn.Voorbestemming // Gebruik Voorbestemming voor duidelijkheid
				if lijn.Buitenlijn {
					buitenlijnStatus = "Ja"
					if lijn.Voorbestemming == "" {
						bestemmingen = "Gemengd" // Aanduiding dat deze lijn gemengde bestemmingen heeft
					}
				}
				fmt.Printf("  Lijn %s: %d/%d auto's, Buitenlijn: %s, Bestemming(en): %s\n",
					lijn.LijnID, lijn.Bezet, lijn.Capaciteit, buitenlijnStatus, bestemmingen)
			} else if lijn.Buitenlijn { // Toon buitenlijnen ook als ze niet gebruikt zijn
				fmt.Printf("  Lijn %s: %d/%d auto's, Buitenlijn: Ja, Niet gebruikt\n",
					lijn.LijnID, lijn.Bezet, lijn.Capaciteit)
			}
		}
	}
}

//voorlopig nog niet nodig
/*
func vulBuitenlijnAan(lijn *Lijn, autosGroepen *[]AutoGroep) {
	beschikbareCapaciteit := lijn.Capaciteit - lijn.Bezet
	for i, groep := range *autosGroepen {
		rest := groep.Aantal % lijn.Capaciteit
		if rest > 0 && rest <= beschikbareCapaciteit {
			// Voeg toe aan buitenlijn
			lijn.Bezet += rest
			(*autosGroepen)[i].Aantal -= rest
			beschikbareCapaciteit -= rest
			if beschikbareCapaciteit == 0 {
				break
			}
		}
	}
}
*/

func maakAanvulPogingen(groep AutoGroep, blokken []Blok) []Toewijzing {
	// Implementeer logica om te proberen bestaande toewijzingen aan te vullen
}

func maakNieuweToewijzingen(groep AutoGroep, blokken []Blok, aanvulPogingen []Toewijzing) []Toewijzing {
	// Implementeer logica om nieuwe toewijzingen te maken na aanvulpogingen
}

func updateToewijzingenPerGroep(groep AutoGroep, aanvulPogingen, nieuweToewijzingen []Toewijzing) {
	// Voeg resultaten toe aan toewijzingenPerGroep
}

func controleerOnvolledigeToewijzingen(autosGroepen []AutoGroep) {
	// Meld groepen die niet volledig zijn toegewezen
}

func toewijzenAutos(autosGroepen []AutoGroep, blokken []Blok) {
	toewijzingenPerGroep = make(map[string][]Toewijzing) // Initialiseren

	// Loop over elke autogroep om te proberen ze toe te wijzen
	for _, groep := range autosGroepen {
		aanvulPogingen := maakAanvulPogingen(groep, blokken)
		nieuweToewijzingen := maakNieuweToewijzingen(groep, blokken, aanvulPogingen)

		// Update de toewijzingenPerGroep met de resultaten van aanvulPogingen en nieuweToewijzingen
		updateToewijzingenPerGroep(groep, aanvulPogingen, nieuweToewijzingen)
	}

	// Controleer na alle toewijzingen of er groepen zijn die niet volledig zijn toegewezen en meld dit
	controleerOnvolledigeToewijzingen(autosGroepen)
}

func toonEindresultaatPerAutogroep() {
	fmt.Println("Eindresultaat van de toewijzingen per autogroep:")
	for bestemming, toewijzingen := range toewijzingenPerGroep {
		fmt.Printf("Autogroep %s:\n", bestemming)
		for _, toewijzing := range toewijzingen {
			fmt.Printf("  Toegewezen aan lijn %s: %d auto('s)\n", toewijzing.LijnID, toewijzing.Aantal)
		}
	}
}

func main() {
	autos, err := leesAutosBestand("Autos.txt")
	if err != nil {
		fmt.Println("Fout bij het lezen van auto's:", err)
		return
	}

	blokken, err := leesBlokkenBestand("Blokken.txt")
	if err != nil {
		fmt.Println("Fout bij het lezen van blokken:", err)
		return
	}

	// Toewijzen van auto's aan de blokken
	toewijzenAutos(groepeerEnSorteerAutos(autos), blokken)

	// Toon het eindresultaat
	toonEindresultaatPerAutogroep()
}
