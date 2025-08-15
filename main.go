package main

// Grab this stuff cuz we'll need it
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Create a struct of all the weapons
type weaponDetails struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Slot              string `json:"slot"`
	Attunement        string `json:"attunment"`         // match typo in JSON
	Attunement_Rating string `json:"attunement_rating"` // JSON has it as a string
	Rarity            string `json:"rarity"`
}

// Make a map so the weapons can be categorized into type
type weaponsByType map[string][]weaponDetails

// And another struct so that that previous map has a struct to go into
type Database struct {
	Weapons weaponsByType `json:"weapons"` // Exported field so json can set it
}

// Run things
func main() {
	// Gotta read the file before I can use it
	// Wouldn't want to forget to log any errors cuz it's there so I should use it
	data, err := os.ReadFile("itemDatabase.json")
	if err != nil {
		log.Fatal(err)
	}

	// Create a Database var to use
	// And another error check because why not
	var db Database
	if err := json.Unmarshal(data, &db); err != nil {
		log.Fatal(err)
	}

	// Let's give a little hint in case something messes up
	if len(db.Weapons) == 0 {
		log.Println("No weapons loaded (check field names/casing and JSON schema).")
	}

	// You make a loop-de-loop then pull and then your shoes are lookin cool
	// Loop through
	for weaponType, list := range db.Weapons {
		fmt.Println("Weapon type:", weaponType)
		for _, w := range list {
			fmt.Printf(" - %s\n", w.Name)
		}
	}
}
