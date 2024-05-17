package main

import (
	"fmt"
	"goctftools/config/colors" // Assuming colors package defines color constants
	"os/exec"
	"time"
)

var (
    ip string
	action int
)

func main() {
    fmt.Print(color.Yellow)
    fmt.Println("+--------------------------------------------------+")
    fmt.Println("| GoCTFTools                                       |")
    fmt.Println("| By Jageun101                                     |")
    fmt.Println("| GITHUB : https://github.com/Jageun101/ctf-tools  |")
    fmt.Println("| VERSION : 1.0.0-alpha                            |")
    fmt.Println("+--------------------------------------------------+")
    fmt.Print(color.Reset, "\n")
	time.Sleep(2 * time.Second)

	exec.Command("clear")

	menu()
}

func menu(){

		// ? Ask ip
		if ip == ""{
			fmt.Print(color.Cyan + "|> Target IP Address: " + color.Magenta)
			fmt.Scan(&ip)
			fmt.Println(color.Reset)
		} else{
			fmt.Println(color.Cyan+"|> Target IP address: "+color.Magenta+ip+color.Reset)
		}

		// ? Menu
		fmt.Println(color.Cyan+ "+=========> GoCTFTools <==============================+", color.Reset)
		fmt.Println(color.Cyan+ "|> [1] Look at the ports of "+ip)
		fmt.Println(color.Cyan+ "+=====================================================+\n", color.Reset)
	
		// ? Ask menu choice
		fmt.Print(color.Cyan + "|> Command Number: " + color.Magenta)
		fmt.Scan(&action)
		fmt.Println(color.Reset)
		

		switch action{
			case 1 :
				exec.Command("btop")
			default : 
		}
}