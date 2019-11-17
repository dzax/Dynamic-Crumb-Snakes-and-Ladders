package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
const N = 101

type SnLzzz struct {
    snakeHead, snakeTail, bottomLadder, topLadder, position int
}

var SnL [N] SnLzzz

func main(){
    var game, playerName string
    var position int
    theSnL := 0
    position = 0
    playerName = startGame()

    if theSnL == 0{
        startSnL()
        theSnL = theSnL + 1
    }

    fmt.Scanln(&game)
    for game !="QUIT"{
        playerPosition(playerName, &position, &game)
    }
}

func startGame() string{
    var playerName string
    fmt.Println("Welcome to the Dynamic Crumb Snakes and Ladders!")
    fmt.Print("Please write your name: ")
    fmt.Scanln(&playerName)
    fmt.Println()
    return playerName
}

func startSnL() {
    fmt.Println("Snakes :")
	snakes()
	fmt.Println("\nLadders :")
	ladders()
}

func dice() int{
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(6) + 1
}

func playerPosition(playerName string, position *int, game *string){
    var temp, rolled int
    rolled = dice()

    if *position < 100 {
        fmt.Println(playerName, "Rolled", rolled)
        *position = *position + rolled
        if *position > 100 {
            temp = *position - 100
            *position = 100 - temp
        }
        SnL[*position].position = *position
		*position = checkSnL(*position, playerName)
        fmt.Println(playerName, "is now in square number", *position)
        if *position == 100 {
            fmt.Println("\nCongratulations", playerName, ", You have won this game!")
            *game = "QUIT"
        } else {
            fmt.Scanln(game)
        }
    } else {
        *game = "QUIT"
    }
}

func snakes(){
    var num1, num2 int
    i := 0

    rand.Seed(time.Now().UnixNano())

    for i < 10{
        num1 = rand.Intn(64) + 25
        num2 = rand.Intn(64) + 25
        if SnL[num1].snakeHead == 0 && SnL[num1].snakeTail == 0 && SnL[num2].snakeHead == 0 && SnL[num2].snakeTail == 0 && SnL[num1].topLadder == 0 && SnL[num1].bottomLadder == 0 && SnL[num2].topLadder == 0 && SnL[num2].bottomLadder == 0 {
            if num1 > num2 {
                for num1 - num2 < 11 || num2 > num1 {
                    num2 = rand.Intn(79) + 10
                }
                SnL[num1].snakeHead=num1
                SnL[num1].snakeTail=num2
                fmt.Println("The head of the snake is at :" , SnL[num1].snakeHead, "// The tail of the snake is at :", SnL[num1].snakeTail)
                i++
            } else if num2 > num1 {
                for num2 - num1 < 11 || num1 > num2 {
                    num1 = rand.Intn(79) + 10
                }
                SnL[num2].snakeHead=num2
                SnL[num2].snakeTail=num1
                fmt.Println("The head of the snake is at :" , SnL[num2].snakeHead, "// The tail of the snake is at :", SnL[num2].snakeTail)
                i++
            } else {
                num1 = rand.Intn(64) + 25
                num2 = rand.Intn(64) + 25
            }
        }
    }
}

func ladders(){
    var num1, num2 int
    i := 0

    rand.Seed(time.Now().UnixNano())

    for i < 7{
        num1 = rand.Intn(79) + 10
        num2 = rand.Intn(79) + 10
        if SnL[num1].snakeHead == 0 && SnL[num1].snakeTail == 0 && SnL[num2].snakeHead == 0 && SnL[num2].snakeTail == 0 && SnL[num1].topLadder == 0 && SnL[num1].bottomLadder == 0 && SnL[num2].topLadder == 0 && SnL[num2].bottomLadder == 0 {
            if num1 > num2 {
                for num1-num2 < 11 || num2 > num1{
                    num1 = rand.Intn(79) + 10
                }
                SnL[num2].topLadder=num1
                SnL[num2].bottomLadder=num2
                fmt.Println("The bottom of ladders is at :", SnL[num2].bottomLadder, "// The top of ladder is at :" , SnL[num2].topLadder)
                i++
            } else if num2 > num1 {
                for num2-num1 < 11 || num1 > num2{
                    num1 = rand.Intn(79) + 10
                }
                SnL[num1].topLadder=num2
                SnL[num1].bottomLadder=num1
                fmt.Println("The bottom of ladders is at :", SnL[num1].bottomLadder, "// The top of ladder is at :" , SnL[num1].topLadder)
                i++
            } else {
                num1 = rand.Intn(79) + 10
                num2 = rand.Intn(79) + 10
            }
        }
    }
} 

func checkSnL(position int, playerName string) int {
    if SnL[position].position == SnL[position].bottomLadder {
        fmt.Print(playerName, " - You found a ladder at ", position, " You can climb to ", SnL[position].topLadder)
        fmt.Println()
        position = SnL[position].topLadder
    } else if SnL[position].position == SnL[position].snakeHead {
        fmt.Print(playerName, " - You found a snake at ", position, " You go back to ", SnL[position].snakeTail)
        fmt.Println()
        position = SnL[position].snakeTail
    }
    return position
}

func crumb () int{
    if dice()!=100{
        
    }
}