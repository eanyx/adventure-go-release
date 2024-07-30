package main

import "fmt"

func main() {

	fmt.Println("Here is an adventure game in Go")

	var pos = "road"
	var action = ""
	var direction = ""
	var life = 10
	var key = 0
	var matches = 0

	var input int = 0
	
	for ok := true; ok; ok = (input != 2) {

	fmt.Println("What to do? Type help for all commands")

	fmt.Scan(&action)

	switch action {

	case "help":
		fmt.Println("List of commands: help|inv|open|take|watch|dig|walk|go|drink|pos")

	case "walk", "go":
		fmt.Println("In which direction N|S|E|W?")
		fmt.Scan(&direction)
		switch direction {
		case "E":
			fmt.Println("You go East near a great forest")
			pos = "forest"
		case "W":
			fmt.Println("You go West near a beach")
			pos = "beach"
		case "S":
			fmt.Println("You go South near a water fall")
			pos = "waterfall"
		case "N":
			fmt.Println("You go North near a house")
			pos = "house"
		default:
			fmt.Println("You cannot go in that direction")
		}

	case "drink":
		if pos == "waterfall" {
			fmt.Println("You are drinking water")
			if life < 10 {
				life = life + 1
				fmt.Println("life = ", life)
			} else {
				fmt.Println("Your life is at his maximum")
				fmt.Println("life = ", life)
			}
		} else {
			fmt.Println("There is nothing to drink here")
		}

	case "light":
		if pos == "in_house" {
			if matches >= 1 {
				fmt.Println("You light the fireplace")
				matches = matches - 1
			} else {
				fmt.Println("You have no matches")
			}
		}

	case "pos":
		fmt.Println("pos = ", pos)

	case "open":
		if pos == "house" {
			if key >= 1 {
				fmt.Println("You open the door")
				key = key - 1
				pos = "in_house"
			} else {
				fmt.Println("You don't have the key !")
			}
		} else {
			fmt.Println("You cannot open anything here")
		}

	case "dig":
		if pos == "beach" {
			if key == 0 {
				fmt.Println("You find a key")
				key = key + 1
			}
		} else {
			fmt.Println("You cannot dig here")
		}

	case "take":
		if pos == "in_house" {
			fmt.Println("You take the matches")
			matches = matches + 1
		}

	case "watch":
		if pos == "house" {
			fmt.Println("You stand near a house")
		}
		if pos == "in_house" {
			fmt.Println("You are in a great house. There is a fireplace in the middle of the house")
		}
		if pos == "forest" {
			fmt.Println("You are in a great luxuriant forest. There are lot of plants and flowers")
		}
		if pos == "waterfall" {
			fmt.Println("You are near great waterfall. The water is cold")
		}
		if pos == "beach" {
			fmt.Println("You are standing on a beach. There is lot of sand here")
		}

	case "inv":
		fmt.Println("Inventory : ")
		fmt.Println("key = ", key)
		fmt.Println("life = ", life)
		fmt.Println("Matches = ", matches)

	default:
		fmt.Println("Unrecognized, type again")
	}
	}
}
