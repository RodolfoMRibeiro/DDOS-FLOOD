package main

import (
	"bufio"
	"ddos-flood/floodsimulator"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const intro = `
.------------------..------------------..------------------.
| .--------------. || .--------------. || .--------------. |
| |  _______     | || | ____    ____ | || |  _______     | |
| | |_   __ \    | || ||_   \  /   _|| || | |_   __ \    | |
| |   | |__) |   | || |  |   \/   |  | || |   | |__) |   | |
| |   |  __ /    | || |  | |\  /| |  | || |   |  __ /    | |
| |  _| |  \ \_  | || | _| |_\/_| |_ | || |  _| |  \ \_  | |
| | |____| |___| | || ||_____||_____|| || | |____| |___| | |
| |              | || |              | || |              | |
| '--------------' || '--------------' || '--------------' |
'------------------''------------------''------------------'
   Hello, My name is Rodolfo, and welcome to my simulator!
`

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Print(intro)
}

func main() {
	url := getInput("Input URL")
	threadCount := getInput("Input Thread Number")
	duration := getInput("Input Attack Duration (seconds)")

	if threads, err := strconv.Atoi(strings.TrimSpace(threadCount)); err == nil {
		if attackDuration, err := strconv.Atoi(strings.TrimSpace(duration)); err == nil {
			startSimulation(url, uint16(threads), uint32(attackDuration))
		} else {
			log.Fatalf("Invalid attack duration: %v\n", err)
		}
	} else {
		log.Fatalf("Invalid thread number: %v\n", err)
	}
}

func getInput(prompt string) string {
	fmt.Print(prompt + ": ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v\n", err)
	}
	return strings.TrimSpace(input)
}

func startSimulation(url string, threads uint16, duration uint32) {
	flooder := floodsimulator.NewFlooder(url, threads, duration)

	flooder.Start()
}
