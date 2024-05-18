package cli

import "fmt"

var(
	protocoleAttack string
)

func Hydra(ip string){
	protocoleAttack = "ftp"
	fmt.Println("===[workingOn.]===["+protocoleAttack+"]===")
}