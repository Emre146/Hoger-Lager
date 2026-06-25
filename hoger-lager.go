package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// Welkomstbericht voor de speler
	fmt.Println("Welcome to the number guessing game!")
	fmt.Println("im thinking of a number between 1 and 100")

	// Variabele voor aantal kansen
	var chances int

	// Speler kiest moeilijkheid
	fmt.Println("Please select a difficulty:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	var choice int
	fmt.Println("Please select your difficulty level:")

	// Input lezen van speler
	_, err := fmt.Scan(&choice)
	if err != nil {
		// Fout bij input → programma stopt
		fmt.Println("Invalid input!")
		os.Exit(1)
	}

	// Bepalen aantal kansen op basis van keuze
	switch choice {
	case 1:
		chances = 10
	case 2:
		chances = 5
	case 3:
		chances = 3
	default:
		// Ongeldige keuze → programma stopt
		fmt.Println("Invalid choice!")
		os.Exit(1)
	}

	// Bevestiging van gekozen moeilijkheid
	fmt.Printf("Great! You have selected the difficulty level with %d chances.\n", chances)
	fmt.Println("Let's start the game!")

	// Random getal tussen 1 en 100
	targetNumber := rand.Intn(100) + 1

	// Teller voor aantal pogingen
	attempts := 1

	// Game loop: blijft lopen zolang speler kansen heeft
	for attempts <= chances {
		fmt.Printf("Enter your %d guess: ", attempts)

		var guess int

		// Input van speler lezen
		_, err := fmt.Scan(&guess)
		if err != nil {
			// Ongeldige input → opnieuw proberen zonder poging te tellen
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}

		// Check of gok te laag is
		if guess < targetNumber {
			fmt.Println("Inncorrect! The number is greater than", guess)

		// Check of gok te hoog is
		} else if guess > targetNumber {
			fmt.Println("Inncorrect! The number is less then ", guess)

		// Correct geraden
		} else {
			fmt.Println("Congratulations! You guessed the correct number", guess)
			fmt.Printf("Attemps")
			return // spel stopt
		}

		// Poging verhogen
		attempts++
	}

	// Als alle kansen op zijn
	fmt.Println("Sorry! You have used all your chances. The correct number was", targetNumber)
}
