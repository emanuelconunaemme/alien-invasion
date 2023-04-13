package play

import (
	"fmt"
	. "saga/aliens/pkg/data"
	. "saga/aliens/pkg/operations"
)

const (
	MAX_ROUNDS = 10000
)

func Play(inputFile string, aliensNumber int, verbose bool) error {
	cities, err := LoadCities(inputFile)
	if err != nil {
		return fmt.Errorf("Unable to load cities from %s: %s", inputFile, err.Error())
	}
	aliens := GenerateAliens(aliensNumber)
	state := InitState(cities, aliens)

	rounds := 0
	fmt.Printf(state.ToString())
	// destroy cities from initial state
	for _, city := range state.Cities {
		if city.HasFight() {
			DestroyCity(&state, city.Name)
			printCityDestroyed(city.Name, city.Aliens[0].Name, city.Aliens[1].Name)
		}
	}
	for {
		// move each alien
		for _, alien := range state.Aliens {
			newCity := MoveAlien(&state, alien.Name)
			if newCity == nil {
				// the alien is isolated, we just move to the next alien
				printInfo(fmt.Sprintf("%s is stuck", alien.Name))
				continue
			} else {
				printInfo(fmt.Sprintf("%s moved to %s", alien.Name, newCity.Name))
				verbosePrint(state.ToString(), verbose)
			}
			if newCity.HasFight() {
				DestroyCity(&state, newCity.Name)
				printCityDestroyed(newCity.Name, newCity.Aliens[0].Name, newCity.Aliens[1].Name)
				break
			}
		}

		if isTerminal, reason := state.IsTerminal(); isTerminal {
			printGameOver(reason)
			return nil
		}

		rounds++
		if rounds == MAX_ROUNDS {
			printGameOver(fmt.Sprintf("Max number of rounds (%d)", MAX_ROUNDS))
			return nil
		}
	}
}

func verbosePrint(text string, verbose bool) {
	if verbose {
		fmt.Printf(text)
	}
}

func printGameOver(reason string) {
	fmt.Printf("[GAME OVER] %s\n", reason)
}

func printCityDestroyed(city string, alien1 string, alien2 string) {
	printInfo(fmt.Sprintf("%s has been destroyed by %s and %s!", city, alien1, alien2))
}

func printInfo(info string) {
	fmt.Printf("[INFO] %s\n", info)
}
