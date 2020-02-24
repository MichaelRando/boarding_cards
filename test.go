package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TransportMeans int
type Location string

const (
	TRAIN TransportMeans = 1 + iota
	AIRPORT_BUS
	FLIGHT
)

const (
	MADRID         Location = "Madrid"
	BARCELONA               = "Barcelona"
	GERONA_AIRPORT          = "Gerona Airport"
	STOCKHOLM               = "Stockholm"
	NEW_YORK_JFK            = "New York JFK"
)

type boardingCard struct {
	means               TransportMeans
	specificMeans       string
	seatAssignment      string
	origin              Location
	destination         Location
	baggageInstructions string
}

var arrayOfCards = []boardingCard{
	{
		means:               TRAIN,
		specificMeans:       "78A",
		seatAssignment:      "Sit in seat 45B",
		origin:              MADRID,
		destination:         BARCELONA,
		baggageInstructions: "",
	},
	{
		means:               AIRPORT_BUS,
		specificMeans:       "",
		seatAssignment:      "No Seat Assignment",
		origin:              BARCELONA,
		destination:         GERONA_AIRPORT,
		baggageInstructions: "",
	},
	{
		means:               FLIGHT,
		specificMeans:       "SK455",
		seatAssignment:      "Gate 45B, seat 3A",
		origin:              GERONA_AIRPORT,
		destination:         STOCKHOLM,
		baggageInstructions: "Baggage drop at ticket counter 344",
	},
	{
		means:               FLIGHT,
		specificMeans:       "SK22",
		seatAssignment:      "Gate 22, seat 7B",
		origin:              STOCKHOLM,
		destination:         NEW_YORK_JFK,
		baggageInstructions: "Baggage will be automatically transferred from your last leg",
	},
}

func printOutBoardingCardArray(theArray []boardingCard) {
	for _, card := range theArray {
		switch card.means {
		case TRAIN:
			fmt.Printf("Take train %s from %s to %s\n", card.specificMeans, card.origin, card.destination)
			break
		case AIRPORT_BUS:
			fmt.Printf("Take the airport bus from %s to %s. %s\n", card.origin, card.destination, card.seatAssignment)
			break
		case FLIGHT:
			fmt.Printf("From %s, take flight %s to %s. %s. %s\n", card.origin, card.specificMeans, card.destination, card.seatAssignment, card.baggageInstructions)
			break
		}
	}
}

func sortBoardingCards(theArray []boardingCard) []boardingCard {
 // need to match destination from second card to origin of first card
	originMapToBoardingCard := make(map[string]boardingCard)
	originSet := make(map[Location]bool)
	for _, card := range theArray {
		originMapToBoardingCard[string(card.origin)] = card
		if _, exist := originSet[(card.origin)]; !exist {
			originSet[(card.origin)] = true
		}
		originSet[(card.destination)] = false
	}

	var currentLocation string
	for k := range originSet {         // Loop
		if originSet[k] {
			currentLocation = string(k)
		}
	}
	var val boardingCard
	sortedCards := make([]boardingCard, len(theArray))
	for i := 0; i < len(theArray); i += 1 {
		val, _ = originMapToBoardingCard[string(currentLocation)]
		sortedCards[i] = val
		currentLocation = string(val.destination)
	}
	return sortedCards
}


func main() {
	// First pass thru this, we're hardcoding the specific cards into a array, shuffling them,
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arrayOfCards), func(i, j int) { arrayOfCards[i], arrayOfCards[j] = arrayOfCards[j], arrayOfCards[i] })
	// print out the shuffled order
	printOutBoardingCardArray(arrayOfCards)
	fmt.Printf("\n\n")
	// sorting the board cards out
	sortedCards := sortBoardingCards(arrayOfCards)
	// and print out the shuffled version
	printOutBoardingCardArray(sortedCards)
}
