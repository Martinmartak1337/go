package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var factions = []string{"S", "K", "R"} // forget this
	fmt.Println("HELLO WHICH FACTION DO YOU CHOOSE: \n")
	for i := len(factions) - 1; i >= 0; i-- {
		fmt.Println(factions[i])
	}
	var inputfaction string
	_, err := fmt.Scanln(&inputfaction)
	if err != nil {
		return
	}
	var StartingSpaceShip string
	var WarpFactor int
	const MaxWarpFactor = 7
	var CrewCount = 250
	var ShieldsUp bool
	var IsInEnemyTeritorry bool
	var inputSpaceName string
	var WorldMap = []string{"S", "S", "S", "S", "S", "S", "K", "K", "K", "K", "R", "R", "R", "R", "R", "R"}
	var TypeWorldMap = WorldMap
	var CurrentPosition int
	switch inputfaction {
	case factions[0]:
		for {
			StartingSpaceShip = "USS-"
			fmt.Println("Please Enter the name of your space-ship: \n")
			_, err2 := fmt.Scanln(&inputSpaceName)
			if err2 != nil {
				continue
			} else {
				break
			}
		}
		StartingSpaceShip = StartingSpaceShip + inputSpaceName
		CurrentPosition = 0
		break
	case factions[1]:
		StartingSpaceShip = ""
		for {
			fmt.Println("Please enter the name of your bird of prey: \n")
			_, err5 := fmt.Scanln(&inputSpaceName)
			if err5 != nil {
				continue
			} else {
				break
			}
		}
		StartingSpaceShip = StartingSpaceShip + inputSpaceName
		CurrentPosition = 8
		break
	case factions[2]:
		StartingSpaceShip = "D'Deridex-"
		CurrentPosition = 14
		for{
		   fmt.Println("Please enter the name of your D'Deridex battle-cruiser: \n")
		   _, err00 := fmt.Scanln(&inputSpaceName)
		   if err00 != nil {
		      continue
		   } else {
		      break
		   }
		}
	}
	for {
		NewOutPutMethod := make([]string, 4)
		WorldMap[CurrentPosition] = WorldMap[CurrentPosition] + "P"
		for j := 0; j < len(WorldMap)-1; j++ {
			if j%3 == 0 {
				NewOutPutMethod = append(NewOutPutMethod, WorldMap[j])
				fmt.Println(NewOutPutMethod)
				NewOutPutMethod = nil
			} else {
				NewOutPutMethod = append(NewOutPutMethod, WorldMap[j])
			}
		}
		NewOutPutMethod = nil
		for {
			fmt.Println("Please Enter Your Warp factor, this will determine where you'll jump on the player map per-game-tick! \n")
			_, err3 := fmt.Scanln(&WarpFactor)
			if err3 != nil {
				continue
			}
			if WarpFactor > MaxWarpFactor || WarpFactor == 0 {
				fmt.Println("Maxt Warp factor is 7 and Min. is 1, Please enter a valid Warp Factor!")
			} else {
				break
			}
		}
		fmt.Println("Shields up: ")
		fmt.Println(ShieldsUp)
		fmt.Println("Current crew count: ")
		fmt.Println(CrewCount)
		if CurrentPosition+WarpFactor > len(WorldMap)-1 {
			if CurrentPosition-1 > 0{
			   WorldMap[CurrentPosition-1] = TypeWorldMap[CurrentPosition-1]
			}
			CurrentPosition = len(WorldMap) - 1
			WorldMap[CurrentPosition] = WorldMap[CurrentPosition] + "P"
			WarpFactor = 0
		} else {
			if CurrentPosition-1 > 0{
			   WorldMap[CurrentPosition-1] = WorldMap[CurrentPosition-1]
			}
			CurrentPosition = CurrentPosition + WarpFactor
			WorldMap[CurrentPosition] = WorldMap[CurrentPosition] + "P"
			WarpFactor = 0
		}
		if !strings.Contains(WorldMap[CurrentPosition], inputfaction) {
			IsInEnemyTeritorry = true
		}
		var GeneralInput int
		rand.Seed(time.Now().UnixNano())
		var RandomChance = rand.Intn(1000)
		if IsInEnemyTeritorry == true {
			if RandomChance < 400 {
				for {
					fmt.Println("A Spaceship approaches you and start's firing what do you do?")
					fmt.Println("1. Fire-Back, 2. Shields-up, 3. Nothing, 4.Call them: \n")
					_, err4 := fmt.Scanln(&GeneralInput)
					if err4 != nil {
						continue
					}
					if GeneralInput < 1 || GeneralInput > 4 {
						continue
					} else {
						break
					}
				}
				switch GeneralInput {
				case 1:
					continue
				case 2:
					ShieldsUp = true
				case 3:
					fmt.Println("All decks blow up, the Bridge Blows up The crew is dead!")
					return
				case 4:

				}
			} else {
				continue
			}
		}
	}	
}
